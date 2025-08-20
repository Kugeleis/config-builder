// main.go
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v2"
)

// This is our template YAML content, embedded directly in the Go code
// to make the example fully self-contained. In a real-world application,
// you would load this from a file.
const templateYAML = `
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

`

// A map to hold the configuration data. `interface{}` allows for any YAML type.
type Config map[string]interface{}

// loadTemplateBytes decides which template to use (external file or internal fallback),
// reads it, and returns its content as a byte slice.
func loadTemplateBytes(templatePath string) ([]byte, error) {
	if templatePath != "" {
		// If a template path is provided, read the file.
		return ioutil.ReadFile(templatePath)
	}
	// Otherwise, use the embedded template.
	return []byte(templateYAML), nil
}

// main function to run the CLI tool.
func main() {
	// Define and parse command-line flags.
	var templatePath string
	flag.StringVar(&templatePath, "template", "", "Path to a custom YAML template file.")
	flag.StringVar(&templatePath, "t", "", "Path to a custom YAML template file (shorthand).")
	flag.Parse()

	fmt.Println("Welcome to the Interactive YAML Configurator!")
	fmt.Println("--------------------------------------------")
	fmt.Println("I will guide you through creating a user-config.yaml file.")

	// Load the template content using the new testable function.
	yamlTemplateContent, err := loadTemplateBytes(templatePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading template: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal the template YAML into a Go map.
	var templateConfig Config
	err = yaml.Unmarshal(yamlTemplateContent, &templateConfig)
	if err != nil {
		fmt.Printf("Error parsing template YAML: %v\n", err)
		os.Exit(1)
	}

	// Create a new empty map to store the user's responses.
	userConfig := make(Config)

	// Call the recursive function to process the config map.
	processConfig(templateConfig, userConfig, "")

	// Marshal the user's collected configuration back into YAML.
	outputYAML, err := yaml.Marshal(userConfig)
	if err != nil {
		fmt.Printf("Error generating output YAML: %v\n", err)
		os.Exit(1)
	}

	// Write the final YAML to a file.
	err = ioutil.WriteFile("user-config.yaml", outputYAML, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n--------------------------------------------")
	fmt.Println("Configuration complete! The file 'user-config.yaml' has been created.")
}

// processConfig is a recursive function to walk the template map.
// It takes the template map, the map to save results to, and the current path (for context).
func processConfig(templateMap, userMap Config, path string) {
	// Loop over each key-value pair in the template.
	for key, value := range templateMap {
		// Construct the full path for the current key.
		currentPath := key
		if path != "" {
			currentPath = path + "." + key
		}

		// Use a type switch to determine what kind of data we are processing.
		switch v := value.(type) {
		case map[interface{}]interface{}:
			// If the value is a nested map, it's a sub-group.
			// Recursively call the function to process the nested map.
			nestedTemplate := make(Config)
			for subKey, subValue := range v {
				// Convert map keys to string.
				nestedTemplate[fmt.Sprintf("%v", subKey)] = subValue
			}
			nestedUserMap := make(Config)
			processConfig(nestedTemplate, nestedUserMap, currentPath)
			userMap[key] = nestedUserMap
		case bool:
			// If the value is a boolean, prompt with a Confirm question.
			// The prompt text is derived from the key and its comment.
			prompt := &survey.Confirm{
				Message: fmt.Sprintf("Set '%s' to:", currentPath),
				Default: v,
			}
			var result bool
			survey.AskOne(prompt, &result)
			userMap[key] = result
		case string:
			// If the value is a string, prompt with a simple Input question.
			prompt := &survey.Input{
				Message: fmt.Sprintf("Enter value for '%s':", currentPath),
				Default: v,
			}
			var result string
			survey.AskOne(prompt, &result)
			userMap[key] = result
		case int:
			// If the value is an integer, prompt with an Input question.
			// The response is then converted back to an integer.
			prompt := &survey.Input{
				Message: fmt.Sprintf("Enter value for '%s':", currentPath),
				Default: fmt.Sprintf("%d", v),
			}
			var result string
			survey.AskOne(prompt, &result)
			// For a simple example, we'll store the string. In a production app,
			// you'd add validation and conversion to an int.
			userMap[key] = result
		case []interface{}:
			// If the value is a list, it's a multiselect question.
			// We need to provide a list of options for the user to select from.
			// For this example, we will hardcode the options. In a real app,
			// these might be derived from a specific field in the template or an external source.
			options := []string{
				"web_server", "database", "message_queue", "caching_layer",
			}
			// Use the survey.MultiSelect prompt.
			prompt := &survey.MultiSelect{
				Message: fmt.Sprintf("Select enabled services for '%s':", currentPath),
				Options: options,
			}
			var results []string
			survey.AskOne(prompt, &results)
			userMap[key] = results
		default:
			// Handle any other types with a generic input prompt.
			prompt := &survey.Input{
				Message: fmt.Sprintf("Enter value for '%s' (type unknown):", currentPath),
				Default: fmt.Sprintf("%v", v),
			}
			var result string
			survey.AskOne(prompt, &result)
			userMap[key] = result
		}
	}
}

// In a real application, you would need to:
// 1. Add robust error handling for user input (e.g., converting string to int).
// 2. Dynamically determine the multiselect options from the template file itself,
//    perhaps using a custom comment format or a specific key like `_options`.
// 3. Add more sophisticated handling for different data types like arrays and maps.
// 4. Use a dedicated file for the template rather than a constant string.
