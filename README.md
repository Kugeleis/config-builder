# Go YAML Configurator CLI

This is a simple, interactive command-line interface (CLI) tool written in Go that guides a user through the creation of a YAML configuration file based on a template. It simplifies the setup process for end-users by prompting them for each required value.

## Getting Started

### Installation

You can download the latest pre-compiled binary for your operating system from this project's **GitHub Releases** page.

### Usage

Once downloaded, you can run the tool from your terminal.

```bash
# Run with the default embedded template
./config-builder

# Specify a custom output file
./config-builder -o my-config.yaml

# Use a custom template file
./config-builder -t my-template.yaml -o my-config.yaml

# Display the help message for all options
./config-builder --help
```

**Command-line Options:**

*   `-t, --template <path>`: Path to a custom YAML template file. If not provided, the tool uses a default internal template.
*   `-o, --output <path>`: Path for the generated YAML file. Defaults to `user-config.yaml`.
*   `-h, --help`: Shows the help message.

## Template File Explained

The tool uses a YAML template to generate the interactive prompts. You can provide your own template using the `-t` flag. The structure of the YAML determines the type of question asked.

### Example `template.yaml`

Here is an example of a valid template, which is the same as the one embedded in the tool:

```yaml
# This is the template for a new user configuration file.
# The tool will prompt for each setting.

general:
  # Do you want to enable debug logging? (true/false)
  debug_logging: false
  # What is your API key? (string)
  api_key: ""
  # How many concurrent threads should be used? (integer)
  threads: 4

services:
  # Which services should be enabled? (multiselect)
  # Possible options are: "web_server", "database", "message_queue", "caching_layer"
  enabled_services:
    - web_server
    - database

  # This is a nested group for database configuration.
  database:
    # What is the database host?
    host: localhost
    # What is the database port?
    port: 5432
```

### How Template Values Work

*   **Strings (`api_key: ""`)**: Generates a text input prompt, with the value as the default.
*   **Integers (`threads: 4`)**: Generates a text input prompt, with the value as the default.
*   **Booleans (`debug_logging: false`)**: Generates a `[y/N]` confirmation prompt.
*   **Nested Maps (`database: ...`)**: Groups related questions under a common heading (e.g., `services.database.host`).
*   **Arrays (`enabled_services: [...]`)**: Generates a multi-select prompt where the user can choose multiple options. The values in the template array are used as the default selections.
    *   **Note:** In the current version of this tool, the *available options* for the multi-select prompt (`web_server`, `database`, etc.) are hard-coded in the application itself, not read from the template.

## For Developers

### Build

To build the binaries from source, you can use GoReleaser.

First, install GoReleaser by following the instructions here: https://goreleaser.com/install/

Then, run the following command to build the binaries:
```bash
goreleaser build --snapshot --clean
```
The binaries will be located in the `dist` directory.

### Release

To create a new release, create a new tag and push it to the repository. The GitHub Actions workflow will then automatically build and release the new version.
```bash
git tag -a v0.2.0 -m "Release v0.2.0"
git push origin v0.2.0
```
