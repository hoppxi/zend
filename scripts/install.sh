#!/usr/bin/env bash
set -e

# Determine privilege level
if [ "$(id -u)" -eq 0 ]; then
    SYSTEM_INSTALL=true
    INSTALL_DIR="/opt/zend"
    BIN_DIR="/usr/local/bin"
    CONFIG_DIR="/etc/zend"
else
    SYSTEM_INSTALL=false
    INSTALL_DIR="$HOME/.local/share/zend"
    BIN_DIR="$HOME/.local/bin"
    CONFIG_DIR="$XDG_CONFIG_HOME/zend"
    mkdir -p "$BIN_DIR"
fi

mkdir -p "$INSTALL_DIR"
mkdir -p "$CONFIG_DIR"

echo "Building Go binary..."
go build -o "$INSTALL_DIR/zend" ./cmd

echo "Building frontend..."
cd web
npm install
npm run build
cd ..
cp -r web/dist "$INSTALL_DIR/dist"

echo "Linking binary..."
ln -sf "$INSTALL_DIR/zend" "$BIN_DIR/zend"

echo "Writing default config if missing..."
CONFIG_FILE="$CONFIG_DIR/config.yaml"
if [ ! -f "$CONFIG_FILE" ]; then
    cp "$INSTALL_DIR/config.yaml" "$CONFIG_FILE" 2>/dev/null || echo "# default config" > "$CONFIG_FILE"
fi

echo "Setting ZEND_DIST environment variable..."
ZEND_LINE="export ZEND_DIST=\"$INSTALL_DIR/dist\""
if [ "$SYSTEM_INSTALL" = true ]; then
    echo "$ZEND_LINE" > /etc/profile.d/zend.sh
else
    SHELL_RC="$HOME/.profile"
    grep -Fxq "$ZEND_LINE" "$SHELL_RC" || echo "$ZEND_LINE" >> "$SHELL_RC"
fi

echo "Installation complete."
echo "Binary: $BIN_DIR/zend"
echo "Config: $CONFIG_FILE"
echo "Dist: $INSTALL_DIR/dist"
echo "Restart your terminal if using per-user install."
