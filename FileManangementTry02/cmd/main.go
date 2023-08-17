package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/yasin-yumrutas/actions"
	"github.com/yasin-yumrutas/cores"
)

var err error
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for {
		fmt.Println("*1 _ Dosya Oluşturma ")
		fmt.Println("*2 _ Dosya Kopyalama ")
		fmt.Println("*3 _ Dosya Silme ")
		fmt.Println("*4 _ Programdan çıkmak istiyorum ")
		fmt.Print("Ne tür bir dosya işlemi yapmak istiyorsunuz?(number) :")
		scanner.Scan()
		var key = scanner.Text()
		keyReal, err := strconv.Atoi(key)
		if err != nil {
			fmt.Errorf(err.Error())
		}

		switch keyReal {
		case 1:
			//Dosya Oluşturma
			cores.CreateInfoFile()
		case 2:
			//Dosya kopyalama
			cores.CopyCoreFile()
		case 3:
			// Dosya silme
			cores.DeleteCoreFile()
		case 4:
			//Uygulamayı sonlandırma
			actions.ShutDown()
		default:
			fmt.Println("Hatalı giriş !!!Tekrar deneyiniz!!!")
		}
	}
}
