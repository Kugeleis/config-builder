package main

import (
	"bytes"
	"os"
	"testing"
)

// TestLoadTemplateBytes_Fallback tests that the function correctly falls back to the
// internal template when no path is provided.
func TestLoadTemplateBytes_Fallback(t *testing.T) {
	// Call the function with an empty path.
	content, err := loadTemplateBytes("")
	if err != nil {
		t.Fatalf("loadTemplateBytes() with empty path returned an unexpected error: %v", err)
	}

	// Check if the content matches the internal template.
	expectedContent := []byte(templateYAML)
	if !bytes.Equal(content, expectedContent) {
		t.Errorf("Expected content to match internal template, but it did not")
	}
}

// TestLoadTemplateBytes_ExternalFile tests that the function correctly reads
// from a provided external file.
func TestLoadTemplateBytes_ExternalFile(t *testing.T) {
	// Define test file content and path.
	testContent := []byte("external_template: true")
	testPath := "test_template.yaml"

	// Create a temporary file for the test.
	err := os.WriteFile(testPath, testContent, 0644)
	if err != nil {
		t.Fatalf("Failed to create temporary test file: %v", err)
	}
	// Ensure the file is cleaned up after the test.
	t.Cleanup(func() {
		os.Remove(testPath)
	})

	// Call the function with the path to the temporary file.
	content, err := loadTemplateBytes(testPath)
	if err != nil {
		t.Fatalf("loadTemplateBytes() with valid path returned an unexpected error: %v", err)
	}

	// Check if the content matches the file content.
	if !bytes.Equal(content, testContent) {
		t.Errorf("Expected content to match the external file, but it did not. Got: %s, Want: %s", content, testContent)
	}
}

// TestLoadTemplateBytes_FileError tests that the function correctly returns an
// error when the provided file path does not exist.
func TestLoadTemplateBytes_FileError(t *testing.T) {
	nonExistentPath := "this_file_should_not_exist.yaml"

	// Call the function with a non-existent path.
	content, err := loadTemplateBytes(nonExistentPath)
	if err == nil {
		t.Fatalf("loadTemplateBytes() with non-existent path did not return an error, but one was expected.")
	}

	// Content should be nil when an error occurs.
	if content != nil {
		t.Errorf("Expected content to be nil on error, but it was not. Got: %s", content)
	}
}
