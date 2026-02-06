<img width="1100" height="409" alt="image" src="https://github.com/user-attachments/assets/5cf7a26d-aaca-489e-84ec-4b54643c92de" />

Boilerplate is a CLI tool that lets developers save and reuse file or project templates, including automated setup commands, to quickly generate code structures and eliminate repetitive work.

## Installation

### Using Homebrew :

```bash
brew tap ukhirani/bp
brew install bp
```

### Linux

```bash
curl -fsSL https://raw.githubusercontent.com/ukhirani/boilerplate/main/install.sh | sh
```

### If you have go already installed.

```bash
go install github.com/ukhirani/boilerplate/bp@latest
```

## Usage

```bash
# Get current version
bp --version

# List available templates
bp list

# Generate from template
bp gen <template>
bp gen <template> <custom-name>  # Custom filename (files only)
bp gen <template> --dir <relative-target-dir>    # Specify relative target directory

# Add new template
bp add <file-or-directory> --name <template-name>

# Preview template
bp preview <template>
bp preview <template> --config  # Show template configuration including pre and post commands.

# Configure a template
bp config <template> # This will open up the config of the template your the default editor (default : vscode)

# Copy template to clipboard (file type only)
bp clip <template>
```

Templates are stored in `$HOME/boilerplate/templates/`. Use `bp --help` for more options.
