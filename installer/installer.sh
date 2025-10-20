#!/bin/bash

set -euo pipefail

# 定義
INSTALL_DIR="$HOME/.acutils-cli"
BIN_DIR="$INSTALL_DIR/bin"
BINARY_BASE_URL="https://lemolatoon.github.io/acutils-cli"
TARGET_BINARY="$BIN_DIR/acutils-cli"

detect_platform() {
    local os arch

    case "$(uname -s)" in
        Linux) os="linux" ;;
        Darwin) os="darwin" ;;
        *)
            echo "Unsupported operating system: $(uname -s)" >&2
            exit 1
            ;;
    esac

    case "$(uname -m)" in
        x86_64|amd64) arch="amd64" ;;
        arm64|aarch64) arch="arm64" ;;
        *)
            echo "Unsupported CPU architecture: $(uname -m)" >&2
            exit 1
            ;;
    esac

    echo "${os}-${arch}"
}

PLATFORM="$(detect_platform)"
BINARY_URL="${BINARY_BASE_URL}/acutils-cli-${PLATFORM}"
TEMP_BINARY="${TARGET_BINARY}.tmp"

# $HOME/.acutils-cli がなければ作成
if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creating installation directory at $INSTALL_DIR"
    mkdir -p "$INSTALL_DIR"
fi

# bin ディレクトリがなければ作成
if [ ! -d "$BIN_DIR" ]; then
    echo "Creating bin directory at $BIN_DIR"
    mkdir -p "$BIN_DIR"
fi

# バイナリをダウンロードして配置
echo "Detected platform: ${PLATFORM}"
echo "Downloading acutils-cli from $BINARY_URL"
if ! curl -fsSL "$BINARY_URL" -o "$TEMP_BINARY"; then
    echo "Failed to download binary for ${PLATFORM}." >&2
    rm -f "$TEMP_BINARY"
    exit 1
fi

mv "$TEMP_BINARY" "$TARGET_BINARY"

# バイナリに実行権限を付与
chmod +x "$TARGET_BINARY"

# PATH に $HOME/.acutils-cli/bin を加えるように促す
echo ""
echo "#########################################################################################################"
echo "##  Installation completed successfully."
echo "##  Please add the following line to your shell configuration file (e.g., .bashrc, .zshrc):"
echo "##  export PATH=\"$HOME/.acutils-cli/bin:\$PATH\""
echo "##  Or, run this command to update your session's environment:"
echo "##  export PATH=\"$HOME/.acutils-cli/bin:\$PATH\""
echo "#########################################################################################################"
