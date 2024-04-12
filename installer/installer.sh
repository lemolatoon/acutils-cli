#!/bin/bash

# 定義
INSTALL_DIR="$HOME/.acutils-cli"
BIN_DIR="$INSTALL_DIR/bin"
BINARY_URL="https://lemolatoon.github.io/acutils-cli/acutils-cli"

# $HOME/.acutils-cli がなければ作成
if [ ! -d "$INSTALL_DIR" ]; then
    echo "Creating installation directory at $INSTALL_DIR"
    mkdir -p "$INSTALL_DIR"
fi

# $HOME/.acutils-cli/bin/acutils-cli があれば削除
if [ -f "$BIN_DIR/acutils-cli" ]; then
    echo "Removing old binary..."
    rm "$BIN_DIR/acutils-cli"
fi

# bin ディレクトリがなければ作成
if [ ! -d "$BIN_DIR" ]; then
    echo "Creating bin directory at $BIN_DIR"
    mkdir -p "$BIN_DIR"
fi

# バイナリをダウンロードして配置
echo "Downloading acutils-cli from $BINARY_URL"
curl -sSL "$BINARY_URL" -o "$BIN_DIR/acutils-cli"

# バイナリに実行権限を付与
chmod +x "$BIN_DIR/acutils-cli"

# PATH に $HOME/.acutils-cli/bin を加えるように促す
echo ""
echo "#########################################################################################################"
echo "##  Installation completed successfully."
echo "##  Please add the following line to your shell configuration file (e.g., .bashrc, .zshrc):"
echo "##  export PATH=\"$HOME/.acutils-cli/bin:\$PATH\""
echo "##  Or, run this command to update your session's environment:"
echo "##  export PATH=\"$HOME/.acutils-cli/bin:\$PATH\""
echo "#########################################################################################################"
