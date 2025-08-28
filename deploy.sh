#!/bin/bash
set -e

# Variables
MC_DIR="/home/minecraft/minecraft-server"
IDLE_DIR="$MC_DIR/idle_worlds"
BACKUP_DIR="$MC_DIR/backups"
SERVICE_FILE="/etc/systemd/system/mc-dashboard.service"
USER="minecraft"
BINARY="$MC_DIR/mc-dash"

# 1️⃣ Ensure directories
mkdir -p "$IDLE_DIR" "$BACKUP_DIR"

# 2️⃣ Build the Go binary
cd "$MC_DIR"
go build -ldflags "-s -w" -o mc-dash main.go worlds.go server.go logs.go util.go handlers.go

# 3️⃣ Setup systemd service
sudo tee "$SERVICE_FILE" > /dev/null <<EOF
[Unit]
Description=PaperMC Go Dashboard
After=network.target

[Service]
User=$USER
WorkingDirectory=$MC_DIR
ExecStart=$BINARY
Restart=always
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=full
ProtectHome=true

[Install]
WantedBy=multi-user.target
EOF

# 4️⃣ Enable and start service
sudo systemctl daemon-reload
sudo systemctl enable --now mc-dashboard

echo "Deployment done!"
echo "Dashboard should be running at http://<server-ip>:8080"
echo "Idle worlds folder: $IDLE_DIR"
echo "Backups folder: $BACKUP_DIR"

# 5️⃣ Optional: sudoers setup for server control (run only once)
SUDO_FILE="/etc/sudoers.d/mc-dashboard"
if [ ! -f "$SUDO_FILE" ]; then
  echo "$USER ALL=(root) NOPASSWD: /usr/bin/systemctl start papermc, /usr/bin/systemctl stop papermc, /usr/bin/systemctl restart papermc, /usr/bin/journalctl -u papermc *" | sudo tee "$SUDO_FILE"
  sudo visudo -c
  echo "Sudoers entry added for dashboard to control PaperMC."
fi
