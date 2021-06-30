//Напишите программу, которая принимает в качестве аргумента путь директории.
//Программа проверяет список файлов и папок и записывает имена в файл LS.txt.
//Содержимое файла должно быть примерно такое:
//<имя файла или директории> (<размер в байтах> bytes) [<FILE | DIRECTORY>]

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Wrong args. Usage dirls.exe <target directory>")
	}

	destination := args[1]
	filename := "LS.txt"

	result, err := os.ReadDir(destination)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filename, []byte{}, 0777)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(filename, os.O_APPEND, 0777)
	if err != nil {
		log.Fatal(err)
	}

	for _, fs := range result {
		info, err := fs.Info()
		var dir string
		if err != nil {
			log.Fatal(err)
			continue
		}
		if info.IsDir() {
			dir = "DIRECTORY"
		} else {
			dir = "FILE"
		}
		_, err = file.WriteString(fmt.Sprintf("%s (%d bytes) [%s]\n",
			info.Name(), info.Size(), dir))
		if err != nil {
			log.Fatal(err)
		}
	}
}
