#!/usr/bin/env bash
set -e

# ---------- Styling ----------
BOLD="\033[1m"
DIM="\033[2m"
GREEN="\033[32m"
CYAN="\033[36m"
YELLOW="\033[33m"
RED="\033[31m"
RESET="\033[0m"

info() { printf "${CYAN}${BOLD}==>${RESET} %s\n" "$1"; }
success() { printf "${GREEN}${BOLD}OK${RESET}  %s\n" "$1"; }
error() { printf "${RED}${BOLD}ERR${RESET} %s\n" "$1"; }

# ---------- Banner ----------
clear

printf "${BOLD}${CYAN}"
cat <<"EOF"
██████╗ ██████╗ 
██╔══██╗██╔══██╗
██████╔╝██████╔╝
██╔══██╗██╔═══╝ 
██████╔╝██║     
╚═════╝ ╚═╝     
EOF
printf "${RESET}${DIM}Boilerplate CLI Installer${RESET}\n\n"

# ---------- Spinner ----------
spinner() {
  local pid=$1
  local delay=0.1
  local spin='|/-\'
  while kill -0 "$pid" 2>/dev/null; do
    for i in {0..3}; do
      printf "\r${CYAN}${BOLD}[%c]${RESET} " "${spin:$i:1}"
      sleep $delay
    done
  done
  printf "\r    \r"
}

# ---------- Detect OS & Arch ----------
info "Detecting system..."

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
x86_64) ARCH=amd64 ;;
aarch64 | arm64) ARCH=arm64 ;;
*)
  error "Unsupported architecture: $ARCH"
  exit 1
  ;;
esac

success "Detected $OS/$ARCH"

# ---------- Download ----------
URL="https://github.com/ukhirani/boilerplate/releases/latest/download/bp_${OS}_${ARCH}.tar.gz"

info "Downloading bp..."

(
  curl -fsSL "$URL" | tar -xz
) &
spinner $!

success "Download complete"

# ---------- Install ----------
info "Installing to /usr/local/bin..."

(
  sudo mv bp /usr/local/bin/bp 2>/dev/null
  hash -r 2>/dev/null || true
) &
spinner $!

success "bp installed successfully"
printf "\nRun: ${CYAN}bp --help${RESET}\n\n"
