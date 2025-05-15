#!/bin/bash
WORkING_DIR="$(mktemp -d)"
cd $WORkING_DIR
# Set the installation directory to the user's bin directory
INSTALL_DIR="$HOME/.blessnet/bin"
GREEN="\033[32m"
BRIGHT_GREEN="\033[92m"
RED="\033[31m"
NC="\033[0m"
# Create the bin directory if it doesn't exist
mkdir -p $INSTALL_DIR

# Determine the operating system and architecture
OS=$(uname -s)
ARCH=$(uname -m)

# Map architectures to download names
case $ARCH in
    "x86_64")
        ARCH_NAME="amd64"
        ;;
    "aarch64"|"arm64")
        ARCH_NAME="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Check if jq is installed
if ! command -v jq &> /dev/null; then
    echo "jq could not be found. Please install jq to proceed."
    exit 1
fi
# Check if curl is installed
if ! command -v curl &> /dev/null; then
    echo "curl could not be found. Please install curl to proceed."
    exit 1
fi

bls_version=`curl -s https://api.github.com/repos/blocklessnetwork/bls-c2w/releases/latest|jq -r .tag_name`

# Determine the download URL based on the operating system
case $OS in
    "Linux")
        if [[ "$ARCH" == "x86_64" ]]; then
            URL="https://github.com/blocklessnetwork/bls-c2w/releases/download/${bls_version}/bls-c2w-linux.amd64.tar.gz"
        elif [[ "$ARCH" == "aarch64" ]]; then
            URL="https://github.com/blocklessnetwork/bls-c2w/releases/download/${bls_version}/bls-c2w-linux.arm64.tar.gz"
        fi
        ;;
    "Darwin")
        if [[ "$ARCH" == "x86_64" ]]; then
            URL="https://github.com/blocklessnetwork/bls-c2w/releases/download/${bls_version}/bls-c2w-darwin.amd64.tar.gz"
        elif [[ "$ARCH" == "aarch64" || "$ARCH" == "arm64" ]]; then
            URL="https://github.com/blocklessnetwork/bls-c2w/releases/download/${bls_version}/bls-c2w-darwin.arm64.tar.gz"
        fi
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

# Download the binary
echo "Downloading and Extracting Blockless C2W from $URL..."
curl -L $URL | tar -xz 
if [ $? -ne 0 ]; then
    echo "Error downloading Blockless C2W. Please check your internet connection."
    exit 1
fi
# Check if the download was successful
if [ ! -f bls-c2w || -f bls-c2wnet ]; then
    echo "Error: Blockless C2W binary not found in the downloaded archive."
    exit 1
fi

# Move the binary to the user's bin directory
echo "Installing Blockless C2W to $INSTALL_DIR..."
mv bls-c2w bls-c2wnet $INSTALL_DIR

# Make sure the binary is executable
chmod +x $INSTALL_DIR/bls-c2w $INSTALL_DIR/bls-c2wnet

# Clean up
rm  -rf $WORkING_DIR

# Add bin to PATH if not already added
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo -e "${RED}Please add follow line to your shell profile...${NC}"
    echo -e "${BRIGHT_GREEN}export PATH=$INSTALL_DIR:\$PATH${NC}"
fi

# Verify the installation
echo -e "Install complete!"
$INSTALL_DIR/bls-c2w -h