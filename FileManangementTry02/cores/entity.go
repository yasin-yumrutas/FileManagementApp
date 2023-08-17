package cores

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yasin-yumrutas/actions"
)

var err error
var scanner = bufio.NewScanner(os.Stdin)

func CreateInfoFile() {
	fmt.Print("Lütfen dosya adını girin: ")
	scanner.Scan()
	fileName := scanner.Text()
	fmt.Print("Dosya içeriği girmek ister misiniz?y/n : ")
	scanner.Scan()
	c_Answer := scanner.Text()
	fmt.Println()
	content := []byte("")
	if c_Answer == "y" {
		fmt.Print("Sizi dinliyorum :")
		scanner.Scan()
		content = scanner.Bytes()
	}
	CreateCoreFile(fileName, content)
}
func CreateCoreFile(fileName string, content []byte) {
	err := actions.CreateFileAction(fileName, content)
	if err != nil {
		fmt.Println("Dosya oluşturulurken hata oluştu:", err)
		return
	}
	fmt.Println("Dosya oluşturuldu:", fileName)
	fmt.Println("Bu dosyanın yedeğinin bastırılmasını ister misin?y/n : ")
	scanner.Scan()
	cAnswer := scanner.Text()
	if cAnswer == "y" {
		CopyCoreFile()
	}
}

func PathSearchAlgo(path string) string {
	index := -1
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '\\' {
			index = i
			break
		}
	}
	fileName := path[index+1:]

	return fileName
}

func CopyCoreFile() {
	fmt.Print("Dosyalamak istedğiniz dosyayı yolula beraber belirtir misiniz? :")
	scanner.Scan()
	sourceFolder := scanner.Text()
	fmt.Printf("\nŞimdide aynı şekilde göndermek istediğiniz noktadaki dosyayı belirtirmisiniz? :")
	scanner.Scan()
	destinationFolder := scanner.Text()
	_, err = os.Stat(destinationFolder)
	if os.IsNotExist(err) {
		fileName := PathSearchAlgo(destinationFolder)
		CreateCoreFile(fileName, nil)
	}

	err := actions.CopyFileAction(sourceFolder, destinationFolder)
	if err != nil {
		fmt.Printf("Error copying files: %v\n", err)
	} else {
		fmt.Println("Files copied successfully.")
	}
}

func DeleteCoreFile() {
	fmt.Print("Hangi dosyayı silmek istersiniz? : ")
	scanner.Scan()
	fileName := scanner.Text()
	err = actions.DeleteFile(fileName)
	if err != nil {
		fmt.Println("Dosya silinirken hata oluştu:", err)
		return
	}
	fmt.Println("Dosya silindi:", fileName)
}
