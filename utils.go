package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var conversionMap = map[string]string{
	".docx": ".pdf",
	".odt":  ".pdf",
	".epub": ".html",
}

func SplitFilename(filename string) (name, ext string) {
	ext = strings.ToLower(filepath.Ext(filename))
	name = filename[:len(filename)-len(ext)]
	return name, ext
}

func JoinFilename(name, ext string) string {
	if len(ext) > 0 && ext[0] != '.' {
		ext = "." + ext
	}
	return name + ext
}

func DownloadFile(filePath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func ConvertFile(inputFile string) (string, string) {
	fileName, fileExtension := SplitFilename(inputFile)

	outputExtension, exists := conversionMap[fileExtension]
	if !exists {
		println("Error. Format not supported:", fileExtension)
		return "", ""
	}

	outputFile := JoinFilename(fileName, outputExtension)
	println("Converting...")
	cmd := exec.Command("pandoc", inputFile, "--pdf-engine=tectonic", "-o", outputFile)

	output, err := cmd.CombinedOutput()
	if err != nil {
		println("Error. Command execution failed:", err)
		return "", ""
	}
	println(output)
	return fileName, outputExtension
}
