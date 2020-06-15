package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printFileInfo(path string, fileName string) {
	file, err := os.Open(path + fileName)
	if err != nil {
		panic(err)
	} else {
		fmt.Print("File was opened successfully\n")

		tmp, _ := file.Stat()
		fileSize := tmp.Size()

		fmt.Printf("File size in bytes: %d\n", fileSize)
	}
}

func readFile1(path string, fileName string) {

	file, err := os.Open(path + fileName)
	if err != nil {
		panic(err)
	} else {
		fmt.Print("File was opened successfully\n")

		tmp, _ := file.Stat()
		fileSize := tmp.Size()

		fmt.Printf("File size in bytes: %d\n", fileSize)

		// DEMO
		// ===================================================== //
		fmt.Printf("File: %v\n", file)
		fmt.Printf("tmp: %v\n", tmp)
		// ===================================================== //
	}

	fmt.Print("\n")

	defer file.Close()

}

func readFile2(path string, fileName string) {

	fileB, err := ioutil.ReadFile(path + fileName)

	if err != nil {
		panic(err)
	} else {
		fmt.Print("File was opened successfully\n")

		fileSize := len(fileB)
		fmt.Printf("File size in bytes: %d\n\n", fileSize)

		fmt.Printf("Byte slice of %s\n", fileName)
		fmt.Print(fileB)
		fmt.Print("\n\n")

		// Convert byte slice into String
		//fmt.Print(hex.Dump(fileB))

		//Loop via byte slice
		for i, byte := range fileB {
			//fmt.Printf("[%d] --- byte: %#x\n", (i + 1), byte)
			//fmt.Printf("[%d] --- byte: %v\n", (i + 1), byte)
			fmt.Printf("[%d] --- byte: %b\n", (i + 1), byte)
		}
		fmt.Print("\n\n")
	}
}

func main() {
	fmt.Print("\n")
	path := "../source/"
	fileName := "text.txt"
	//fileName := "image.jpg"

	printFileInfo(path, fileName)

	//readFile1(path, fileName)
	readFile2(path, fileName)

}
