#!/bin/sh
set -e

REPO="JoaoPedr0Maciel/dev"
BINARY="dev"
INSTALL_DIR="/usr/local/bin"

# ── detect OS ────────────────────────────────────────────────────────────────
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS" in
  linux)  ;;
  darwin) ;;
  *)
    echo "Error: unsupported OS '$OS'."
    echo "Install via go install: go install github.com/${REPO}/cmd/dev@latest"
    exit 1
    ;;
esac

# ── detect architecture ──────────────────────────────────────────────────────
ARCH=$(uname -m)
case "$ARCH" in
  x86_64)        ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *)
    echo "Error: unsupported architecture '$ARCH'."
    exit 1
    ;;
esac

# ── resolve latest version ───────────────────────────────────────────────────
VERSION=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" \
  | grep '"tag_name"' \
  | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$VERSION" ]; then
  echo "Error: could not determine latest version."
  exit 1
fi

# ── resolve install directory ────────────────────────────────────────────────
if [ ! -w "$INSTALL_DIR" ]; then
  INSTALL_DIR="$HOME/.local/bin"
  mkdir -p "$INSTALL_DIR"
fi

# ── download ─────────────────────────────────────────────────────────────────
FILENAME="${BINARY}_${OS}_${ARCH}"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${FILENAME}"

echo "Installing dev ${VERSION} (${OS}/${ARCH})..."
curl -fsSL "$URL" -o "${INSTALL_DIR}/${BINARY}"
chmod +x "${INSTALL_DIR}/${BINARY}"

echo ""
echo "✓ Installed to ${INSTALL_DIR}/${BINARY}"

# ── PATH hint ────────────────────────────────────────────────────────────────
if ! command -v "$BINARY" >/dev/null 2>&1; then
  echo ""
  echo "Note: '${INSTALL_DIR}' is not in your PATH."
  echo "Add this line to your shell profile (~/.bashrc or ~/.zshrc):"
  echo ""
  echo "  export PATH=\"\$PATH:${INSTALL_DIR}\""
fi
