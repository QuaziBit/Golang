package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tcolgate/mp3"
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

		//fmt.Printf("Byte slice of %s\n", fileName)
		//fmt.Print(fileB)
		fmt.Print("\n\n")

		// TEST
		// Print first 3 bytes of the file
		/*
			fmt.Printf("1: %v\n", fileB[0])
			fmt.Printf("2: %v\n", fileB[1])
			fmt.Printf("3: %v\n", fileB[2])

			header := string(fileB[0])
			header += string(fileB[1])
			header += string(fileB[2])
		*/

		/*
			var index int
			var title []byte
			for index = 0; index < 30; index++ {
				//title += string(fileB[index])
				title = append(title, fileB[index])
				tmp := fileB[index]

				fmt.Printf("Index: %d", index)
				fmt.Printf("\thex: %#x --- bit: %b --- str: %s --- val: %v\n", tmp, tmp, string(tmp), tmp)

			}
		*/

		//fmt.Printf("Header: %v\n", header)
		//fmt.Printf("Title: %v\n", title)

		// Convert byte slice into String
		//fmt.Print(hex.Dump(fileB))

		//Loop via byte slice

		// Just loop via all file bytes
		// =============================================================================================== //
		for i, byte := range fileB {
			index := i
			tmpByte := byte
			tmpBits := fmt.Sprintf("%b", byte)

			// Building full byte
			for len(tmpBits) < 8 {
				tmp := "0"
				tmpBits = tmp + tmpBits
			}

			//fmt.Printf("[%d] --- byte %v --- bits: %b\n", (i + 1), byte, byte)
			fmt.Printf("[%d] --- hex: %#x --- bits: %s\n", index, tmpByte, tmpBits)
			// =============================================================================================== //

			fmt.Print("\n\n")
		}
	}
}

func readWriteFileBits(path string, fileName string) {

	fileB, err := ioutil.ReadFile(path + fileName)

	if err != nil {
		panic(err)
	} else {
		fmt.Print("File was opened successfully\n")

		fileSize := len(fileB)
		fmt.Printf("File size in bytes: %d\n\n", fileSize)

		//fmt.Printf("Byte slice of %s\n", fileName)
		//fmt.Print(fileB)
		fmt.Print("\n\n")

		// Write file bits in to the text file
		// =============================================================================================== //

		// Truncates if file already exists, be careful!
		//outFileName := "fileBits.txt"
		outFileName := "fileDecimal.txt"
		file, err := os.Create(outFileName)
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		defer file.Close() // Make sure to close the file when you're done

		for _, byte := range fileB {
			//index := i
			//tmpByte := byte
			tmpBits := fmt.Sprintf("%b", byte)

			// Building full byte
			for len(tmpBits) < 8 {
				tmp := "0"
				tmpBits = tmp + tmpBits
			}

			// Write bits
			/*
				tmp := fmt.Sprintf("%s\n", tmpBits)
				_, err := file.WriteString(tmp)
			*/

			// Write decimal representation of bits

			//key := 1
			//tmpDecimal := fmt.Sprintf("%d%d\n", key, byte)
			//_, err := file.WriteString(tmpDecimal)
			tmpDecimal := fmt.Sprintf("%d", byte)
			_, err := file.WriteString(tmpDecimal)

			if err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}

			//fmt.Printf("[%d] --- hex: %#x --- bits: %s\n", index, tmpByte, tmpBits)
		}
		// =============================================================================================== //

		// Loop via all file bytes
		// =============================================================================================== //
		/*
					bitString := ""
					for i, byte := range fileB {
						index := i
						tmpByte := byte
						tmpBits := fmt.Sprintf("%b", byte)

						// Building full byte
						for len(tmpBits) < 8 {
							tmp := "0"
							tmpBits = tmp + tmpBits
						}

						bitString += tmpBits
						if len(bitString) == 80 {
							printBitSquare(bitString, len(bitString))
							break
						}

						fmt.Printf("[%d] --- hex: %#x --- bits: %s\n", index, tmpByte, tmpBits)

				// =============================================================================================== //
			}
		*/

		fmt.Print("\n\n")
	}
}

/*
func printBitSquare(bitString string, size int) {
	fmt.Printf("\n%s\n", bitString)

	j := 8
	for i := 0; i < size; i++ {
		if i == j {
			fmt.Printf("\n")
			j = j * 2
		} else {
			fmt.Printf("%T", bitString[i])
		}
	}
}
*/

/*
func CreateFile() {
	file, err := os.Create("test.txt") // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close() // Make sure to close the file when you're done

	len, err := file.WriteString("The Go Programming Language, also commonly referred to as Golang, is a general-purpose programming language, developed by a team at Google.")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile Name: %s", file.Name())
}
*/

func readMP3(path string, fileName string) {
	file, err := os.Open(path + fileName)

	if err != nil {
		panic(err)
	} else {
		fmt.Print("MP3 file was opened successfully\n")

		tmp, _ := file.Stat()
		fileSize := tmp.Size()

		fmt.Printf("File size in bytes: %d\n", fileSize)

		decoder := mp3.NewDecoder(file)
		var frames mp3.Frame
		var skipped int

		err := decoder.Decode(&frames, &skipped)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Frames: %v\n", &frames)

	}
	defer file.Close()
}

func main() {
	fmt.Print("\n")
	path := "../audio/"
	//fileName := "audio_sample.mp3"
	fileName := "audio_sample_2.mp3"

	readWriteFileBits(path, fileName)
}
