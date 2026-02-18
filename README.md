<img  height="409" alt="image" src="https://github.com/user-attachments/assets/31d049e8-c313-46b9-9da7-46c293257f75" />

Boilerplate is a CLI tool that lets developers save and reuse file or project templates, including automated setup commands, to quickly generate code structures and eliminate repetitive work.

Find other templates and share your templates at [bp-hub](https://bp-hub.vercel.app)

Refer the official documentation for boilerplate at  [bp-hub/docs](https://bp-hub.vercel.app/docs)

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

#### Get current version
```bash
bp --version
```

#### See the help guide for bp
```bash
bp --help
bp help
```

#### List available templates
```bash
bp list
```

#### Generate from template
```bash
bp gen <template>
bp gen <template> <custom-name>  # Custom filename (files only)
bp gen <template> --dir <relative-target-dir>    # Specify relative target directory
```

#### Add new template
```bash
bp add <file-or-directory> --name <template-name>
```

#### Preview template
```bash
bp preview <template>
bp preview <template> --config  # Show template configuration including pre and post commands.
```

#### Configure a template
```bash
bp config <template> # This will open up the config of the template your the default editor (default : vscode)
bp config <template> -e OR --editor <editor-name> # This is to override the default editor for opening configs
```

#### Copy template to clipboard (file type only)
```bash
bp clip <template>
```

#### Clone template to your system locally (from bp-hub)
```bash
bp clone <username>/<template> --alias <alias-name> # Here, alias is what you want to call this template in your system
```

#### Run the pre and post command of any template (without generating the template)
```bash
bp work <template> # By default runs pre and post commands both (sequentially)
bp work <template> --pre # Runs only the pre commands
bp work <template> --post # Runs only the post commands
```
