<img width="1100" height="409" alt="image" src="https://github.com/user-attachments/assets/5cf7a26d-aaca-489e-84ec-4b54643c92de" />

Boilerplate is a CLI tool that helps developers eliminate repetitive code by managing reusable file and directory templates. Save common code structures, project scaffolds, and configuration files as templates, then generate them instantly when needed.
Templates can include pre and post-generation commands to automate setup tasks, making it ideal for quickly bootstrapping projects, components, or any repetitive codes.

## Installation

```bash
go install github.com/ukhirani/boilerplate/bp@latest
```

## Usage

```bash
# List available templates
bp list


# Generate from template
bp generate <template-name>
bp generate <template-name> --name <custom-name>  # Custom filename (files only)
bp generate <template-name> --dir <relative-target-dir>    # Specify relative target directory

# Add new template
bp add <file-or-directory> --name <template-name>

# Preview template
bp preview <template-name>
bp preview <template-name> --config  # Show template configuration including pre and post commands.

# Copy template to clipboard (file type only)
bp clip <template-name>
```

Templates are stored in `$HOME/boilerplate/templates/`. Use `bp --help` for more options.
