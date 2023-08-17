package actions

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateFileAction(fileName string, content []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

func CopyFileAction(sourceDir, destDir string) error {
	fileList := []string{}

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	for _, srcPath := range fileList {
		relativePath, err := filepath.Rel(sourceDir, srcPath)
		if err != nil {
			return err
		}
		destPath := filepath.Join(destDir, relativePath)

		err = os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
		if err != nil {
			return err
		}

		err = copyFile(srcPath, destPath)
		if err != nil {
			fmt.Printf("Error copying %s: %v\n", srcPath, err)
		}
	}

	return nil
}

func DeleteFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

func ShutDown() {
	fmt.Print("Program kapanıyor!!!")
	return
}
