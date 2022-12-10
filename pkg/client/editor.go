package client

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
)

func openTempFileWithContent(content []byte) (*os.File, error) {
	f, err := os.CreateTemp("", "tmp")
	if err != nil {
		return nil, err
	}
	_, err = f.Write(content)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// editStructInEditor will open the users default text editor and allow them to edit the content provided
// upon closing the editor, the content will be returned
func editStructInEditor(content interface{}) ([]byte, error) {
	// Marshall the content to json
	b, err := json.MarshalIndent(content, "", "    ")
	if err != nil {
		return nil, err
	}
	// Create a temp file with the content
	f, err := openTempFileWithContent(b)
	if err != nil {
		return nil, err
	}
	tmp := f.Name()
	f.Close()
	// Open the file in the users default editor
	defaultEditor, err := determineDefaultTextEditor()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(defaultEditor, tmp)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		os.Remove(tmp)
		return nil, err
	}
	// Read the file back in
	updatedContent, err := os.ReadFile(tmp)
	os.Remove(tmp)
	if string(updatedContent) == string(b) {
		return nil, errors.New("no changes made")
	}
	return updatedContent, err
}

func determineDefaultTextEditor() (string, error) {
	editor := os.Getenv("EDITOR")
	if editor != "" {
		return editor, nil
	}
	editor = os.Getenv("VISUAL")
	if editor != "" {
		return editor, nil
	}
	return "", errors.New("could not determine default text editor")
}
