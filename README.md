# **Go YAML Configurator CLI**

This is a simple, interactive command-line interface (CLI) tool written in Go that guides a user through the creation of a YAML configuration file based on an internal template. It is designed to be a self-contained executable that simplifies the setup process for end-users by prompting them for each required value.

## **Description**

The application reads a predefined YAML template, which includes default values and data types for different configuration fields. It then presents a series of interactive prompts to the user, allowing them to:

* Confirm true/false settings.  
* Enter string and integer values.  
* Select multiple options from a predefined list.

Once all prompts are answered, the tool generates a new user-config.yaml file with the user's choices, ready for use.

## **How to Use**

1. **Run the Tool:** Navigate to the directory where you saved main.go and execute the program with the following command:
   go run main.go

The application will then begin the interactive question-and-answer session. Follow the prompts to configure your file.

## **Build**

To build the binaries for all major operating systems, you can use GoReleaser.

First, install GoReleaser by following the instructions here: https://goreleaser.com/install/

Then, run the following command to build the binaries:

```
goreleaser build --snapshot --rm-dist
```

The binaries will be located in the `dist` directory.

## **Release**

To create a new release, you need to create a new tag and push it to the repository. The GitHub Actions workflow will then automatically build and release the new version.

```
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
```

## **Output**

After you complete the prompts, a file named user-config.yaml will be created in the same directory. The content will reflect your answers, for example:

general:  
  debug\_logging: true  
  api\_key: some-secret-key-123  
  threads: 8  
services:  
  database:  
    host: example.com  
    port: 3306  
  enabled\_services:  
  \- web\_server  
  \- database  
  \- message\_queue  
