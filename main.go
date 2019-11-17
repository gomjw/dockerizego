package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gomjw/color"
)

func main() {
	// Creating subdirectory
	color.Auto("Creating subdirectory...")
	_ = os.Mkdir("dockerizego", os.ModeDir)

	// Get local path
	dir, _ := os.Getwd()
	dir = filepath.Base(dir)

	// Build golang binary
	color.Auto("Building binary file...")
	cmd := exec.Command("go", "build", "-a", "-installsuffix", "cgo", "-o", "./dockerizego/main")
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0", "GOOS=linux")

	if err := cmd.Run(); err != nil {
		color.Auto("Error: could not build golang application!")
		os.Exit(1)
	}

	// Add Dockerfile
	color.Auto("Writing Dockerfile...")
	data := []byte("FROM alpine\nADD main /\nRUN chmod +x main\nENTRYPOINT [\"./main\"]")
	err := ioutil.WriteFile("./dockerizego/Dockerfile", data, 0644)
	if err != nil {
		color.Auto("Error: could not write Dockerfile!")
		os.Exit(1)
	}

	color.Auto("Success: Successfully dockerized application!")
}
