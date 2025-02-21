package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a directory path as an argument")
	}
	directoryPath := os.Args[1]
	imageUriPrefix := "file://"

	randomImagePath, err := getRandomImagePathInDirectory(directoryPath)
	if err != nil {
		log.Fatalf("Failed to get random image from directory (%s): %v\n", err, directoryPath)
	}
	log.Printf("Random image path: %s\n", randomImagePath)
	imagePath := imageUriPrefix + randomImagePath
	changeWallpaper(imagePath)
}

func getRandomImagePathInDirectory(directoryPath string) (string, error) {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return "", err
	}

	// filter out directories
	var fileList []os.DirEntry
	for _, f := range files {
		if !f.IsDir() {
			fileList = append(fileList, f)
		}
	}
	if len(fileList) == 0 {
		return "", fmt.Errorf("no files found")
	}

	//pick a random file
	return directoryPath + "/" + fileList[rand.Intn(len(fileList))].Name(), nil
}

// We have to change wallpaper depending on colour scheme :(
func changeWallpaper(imagePath string) {
	colourScheme := getColourScheme()

	// Set wallpaper for either light or dark, depending on the current colour scheme
	cmd := getCmdExecFromColourScheme(colourScheme, imagePath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to set wallpaper: %v\nOutput: %s", err, string(output))
	}

	log.Println("Wallpaper changed successfully for current colour scheme")
}

func getColourScheme() string {
	colourSchemeCmd := exec.Command("gsettings", "get", "org.gnome.desktop.interface", "color-scheme")
	colourSchemeOutput, err := colourSchemeCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to get colour scheme: %v\nOutput: %s", err, string(colourSchemeOutput))
	}

	colourScheme := strings.TrimSpace(string(colourSchemeOutput))
	log.Printf("Colour scheme detected: %s\n", colourScheme)
	return colourScheme
}

func getCmdExecFromColourScheme(colourScheme string, imagePath string) *exec.Cmd {
	//todo: can there be any, other than prefer-dark?
	var pictureUriFlag string

	// grrrr golang give me ternary >:(
	if colourScheme == "'prefer-dark'" {
		pictureUriFlag = "picture-uri-dark"
	} else {
		pictureUriFlag = "picture-uri"
	}
	return exec.Command("gsettings", "set", "org.gnome.desktop.background", pictureUriFlag, imagePath)
}
