package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	blsc2w "github.com/blocklessnetwork/bls-c2w"
	"github.com/containerd/containerd/archive"
	"github.com/urfave/cli"
)

const defaultOutputFile = "bls-out.wasm"

var dockerfile = blsc2w.Dockerfile

func main() {
	app := cli.NewApp()
	app.Name = "bls-c2w"
	app.Usage = "container to wasm converter"
	app.UsageText = fmt.Sprintf("%s [options] image-name [output file]", app.Name)
	app.ArgsUsage = "image-name [output-file]"
	var flags []cli.Flag
	app.Flags = append([]cli.Flag{
		cli.StringFlag{
			Name:  "dockerfile",
			Usage: "Custom location of Dockerfile (default: embedded to this command)",
		},
		cli.StringFlag{
			Name:  "builder",
			Usage: "Bulider command to use",
			Value: "docker",
		},
		cli.StringFlag{
			Name:  "target-arch",
			Usage: "target architecture of the source image to use",
			Value: "amd64",
		},
		cli.StringSliceFlag{
			Name:  "build-arg",
			Usage: "Additional build arguments",
		},
	}, flags...)
	app.Action = action
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func action(clicontext *cli.Context) error {
	arg1 := clicontext.Args().First()
	outputPath := clicontext.Args().Get(1)
	if arg1 == "" {
		return fmt.Errorf("specify image name")
	}
	builderPath, err := exec.LookPath(clicontext.String("builder"))
	if err != nil {
		return err
	}
	destDir, destFile := ".", defaultOutputFile
	tmpdir, err := os.MkdirTemp("", "blsc2w")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpdir)
	srcImgPath := filepath.Join(tmpdir, "img")
	if err := os.Mkdir(srcImgPath, 0755); err != nil {
		return err
	}
	srcImgName := arg1
	if err := preparedImage(builderPath, srcImgName, srcImgPath, clicontext.String("target-arch")); err != nil {
		return fmt.Errorf("failed to prepare image: %w", err)
	}
	if outputPath != "" {
		d, f := filepath.Split(outputPath)
		destDir, err = filepath.Abs(d)
		if err != nil {
			return err
		}
		if f != "" {
			destFile = f
		}
	}
	return build(builderPath, srcImgPath, destDir, destFile, clicontext)
}

func build(builderPath string, srcImgPath string, destDir, destFile string, clicontext *cli.Context) error {
	buildxArgs := []string{
		"buildx", "build", "--progress=plain",
		"--platform=linux/amd64",
	}
	var dockerfilePath string
	if o := clicontext.String("dockerfile"); o != "" {
		dockerfilePath = o
	} else {
		f, err := os.CreateTemp("", "bls-c2w")
		if err != nil {
			return err
		}
		defer os.Remove(f.Name())
		if _, err := f.Write([]byte(dockerfile)); err != nil {
			return err
		}
		dockerfilePath = f.Name()
	}
	buildxArgs = append(buildxArgs, "-f", dockerfilePath)
	buildxArgs = append(buildxArgs, "--output", fmt.Sprintf("type=local,dest=%s", destDir))
	if destFile != "" {
		buildxArgs = append(buildxArgs, "--build-arg", fmt.Sprintf("OUTPUT_NAME=%s", destFile))
	}
	for _, a := range clicontext.StringSlice("build-arg") {
		buildxArgs = append(buildxArgs, "--build-arg", a)
	}
	buildxArgs = append(buildxArgs, srcImgPath)
	log.Printf("buildx args: %+v\n", buildxArgs)

	cmd := exec.Command(builderPath, buildxArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// save the image to temp directory
func preparedImage(builderPath, imgName, tmpdir, targetarch string) error {
	log.Printf("saving %q to %q\n", imgName, tmpdir)
	needsPull := false
	if idata, err := exec.Command(builderPath, "image", "inspect", imgName).Output(); err != nil {
		needsPull = true
	} else if targetarch != "" {
		inspectData := make([]map[string]interface{}, 1)
		if err := json.Unmarshal(idata, &inspectData); err != nil {
			return err
		}
		if a := inspectData[0]["Architecture"]; a != targetarch {
			log.Printf("unexpected archtecture %v (target: %v). Try \"--target-arch\" when specifying an architecture.\n", a, targetarch)
			needsPull = true
		}
	}
	if needsPull {
		args := []string{"pull"}
		if targetarch != "" {
			args = append(args, "--platform=linux/"+targetarch)
		}
		args = append(args, imgName)
		log.Printf("cannot get image %q locally; pulling it from the registry...\n", imgName)
		pullCmd := exec.Command(builderPath, args...)
		pullCmd.Stdout = os.Stdout
		pullCmd.Stderr = os.Stderr
		if err := pullCmd.Run(); err != nil {
			return fmt.Errorf("failed to pull the image. Try \"--target-arch\" when specifying an architecture: %w", err)
		}
	}
	saveCmd := exec.Command(builderPath, "save", imgName)
	outR, err := saveCmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer outR.Close()
	saveCmd.Stderr = os.Stderr
	if err := saveCmd.Start(); err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	if _, err := archive.Apply(ctx, tmpdir, outR, archive.WithNoSameOwner()); err != nil {
		return err
	}
	if err := saveCmd.Wait(); err != nil {
		return err
	}

	now := time.Now().Local()
	return filepath.Walk(tmpdir, func(p string, info fs.FileInfo, err error) error {
		return os.Chtimes(p, now, now)
	})
}
