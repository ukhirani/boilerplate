<img width="1100" height="409" alt="image" src="https://github.com/user-attachments/assets/5cf7a26d-aaca-489e-84ec-4b54643c92de" />


## Installation

Make sure you have Go installed, then run:

```bash
go install github.com/ukhirani/boilerplate/bp@latest
```

## Quick Start

### List Available Templates

View all templates in your collection:

```bash
bp list
# or
bp ls
```

### Generate from a Template

Create a file or directory from an existing template:

```bash
bp generate <template-name>
# or use aliases
bp gen <template-name>
bp create <template-name>
```

This will copy the template to your current directory, preserving its structure.

### Add a New Template

Save a file or directory as a reusable template:

```bash
bp add <file-or-directory> --name <template-name>
# or use short flag
bp add <file-or-directory> -n <template-name>
```

The template name must contain only letters, numbers, and underscores. The file or directory will be saved to your templates directory for future use.

### Check Version

Display the current version:

```bash
bp version
# or
bp --version
bp -v
```

### Get Help

Show available commands and options:

```bash
bp --help
# or
bp -h
```

## Usage Examples

```bash
# List all available templates
bp list

# Generate a React component template
bp generate react-component

# Add a file as a new template
bp add ./my-script.sh --name shell-script

# Add a directory as a template
bp add ./project-structure --name starter-template
```
