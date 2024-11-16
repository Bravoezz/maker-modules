# Module Creator Script in Go

This is a Go-based script that generates a Node.js module using TypeScript. It creates the following components for the module:

- **Controller**
- **Service**
- **Repository**
- **Contracts**

## Features

- Modular and structured generation.
- Accepts customizable paths for module creation.

## Usage

### Requirements

- Go 1.19+ installed.
- Node.js installed.

### Run the Script

The script takes two arguments:
- \`-name\` (required): Specifies the name of the module to be created.
- \`-path\` (optional): Specifies the path where the module will be created.

Example usage:

\`\`\`bash
go run main.go -name UserModule -path ./modules
\`\`\`

This will create a module named \`UserModule\` in the \`./modules\` directory.

If no \`-path\` is specified, the module will be created in the current directory.

### Flags

| Flag   | Description                           | Required |
|--------|---------------------------------------|----------|
| \`-name\`| Name of the module to be created.     | Yes      |
| \`-path\`| Directory where the module is created.| No       |

## Structure of Generated Module

The generated module will have the following structure:

\`\`\`
ModuleName/
├── controller.ts
├── service.ts
├── repository.ts
└── contracts.ts
\`\`\`

## Error Handling

- If the \`-name\` flag is not provided, the script will exit with an error message.
- If the provided path does not exist, it will attempt to create it.

## Example Output

Upon successful execution, the script outputs:

\`\`\`
*** Exit program ***
\`\`\`

## Notes

- Customize the module templates by editing the internal logic of the \`NewModuleCreator\` function in the script.
- Contributions are welcome!
