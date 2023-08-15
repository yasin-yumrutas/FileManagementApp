package main

import (
	"fmt"
	"io"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// FileOperations interface, SOLID prensiplerine uygun olarak dosya işlemleri için soyutlama sağlar.
type FileOperations interface {
	CreateFile(filename string, content string) error
	ReadFile(filename string) (string, error)
	WriteFile(filename string, content string) error
	DeleteFile(filename string) error
	MoveFile(source string, destination string) error
	CopyFile(source string, destination string) error
}

// FileManager, FileOperations arayüzünü uygular ve dosya işlemlerini gerçekleştirir.
type FileManager struct{}

func (fm *FileManager) CreateFile(filename string, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, content)
	return err
}

func (fm *FileManager) ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (fm *FileManager) WriteFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, content)
	return err
}

func (fm *FileManager) DeleteFile(filename string) error {
	return os.Remove(filename)
}

func (fm *FileManager) MoveFile(source string, destination string) error {
	return os.Rename(source, destination)
}

func (fm *FileManager) CopyFile(source string, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Dosya Yönetimi Uygulaması")

	fileManager := &FileManager{}

	dosyaAdiGirdi := widget.NewEntry()
	icerikGirdi := widget.NewMultiLineEntry()
	silinecekDosyaGirdi := widget.NewEntry()

	// Dosya oluşturma formu
	createForm := container.NewVBox(
		widget.NewLabel("Dosya Oluşturma"),
		widget.NewForm(
			&widget.FormItem{Text: "Dosya Adı:", Widget: dosyaAdiGirdi},
			&widget.FormItem{Text: "İçerik:", Widget: icerikGirdi},
			widget.NewButton("Oluştur", func() {
				// Formdaki değerleri al ve dosya oluştur
				dosyaAdi := dosyaAdiGirdi.Text
				icerik := icerikGirdi.Text
				err := fileManager.CreateFile(dosyaAdi, icerik)
				if err != nil {
					fmt.Println("Dosya oluşturma hatası:", err)
				}
			}),
		),
	)

	// Dosya okuma alanı
	readFileContent := widget.NewMultiLineEntry()

	// Dosya silme formu
	deleteForm := container.NewVBox(
		widget.NewLabel("Dosya Silme"),
		widget.NewForm(
			&widget.FormItem{Text: "Silinecek Dosya Adı:", Widget: silinecekDosyaGirdi},
			widget.NewButton("Sil", func() {
				// Formdaki değeri al ve dosya sil
				dosyaAdi := silinecekDosyaGirdi.Text
				err := fileManager.DeleteFile(dosyaAdi)
				if err != nil {
					fmt.Println("Dosya silme hatası:", err)
				}
			}),
		),
	)

	// Ana pencere düzeni
	myWindow.SetContent(container.NewVBox(
		createForm,
		widget.NewDivider(),
		widget.NewLabel("Dosya İçeriği"),
		readFileContent,
		widget.NewDivider(),
		deleteForm,
	))

	myWindow.ShowAndRun()
}
