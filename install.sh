#!/bin/bash
# Run op de linux server via root user
read -p "Press any key to start deployment"

# Variables
TARGET_DIR="/home/minecraft/mc-roestbak-dashboard"
BINARY_NAME="server-executable"

echo "Preparing target directory..."
mkdir -p "$TARGET_DIR"

echo "Copying binary..."
cp "./../dashboard-executable" "$TARGET_DIR/$BINARY_NAME"

echo "Copying templates..."
cp -r "./src/html" "$TARGET_DIR/html"

echo "Setting permissions..."
chown -R minecraft:minecraft "$TARGET_DIR"
chmod +x "$TARGET_DIR/$BINARY_NAME"

echo "Deployment finished. Target folder:"
ls -l "$TARGET_DIR"

read -p "Press any key to exit"
