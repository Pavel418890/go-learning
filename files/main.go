package main

import (
	"bufio"
	"fmt"
	"io/ioutil"

	"log"
	"os"
)

func main() {
	newFile, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	fileInfo, err := os.Stat("info.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	_ = fileInfo

	renamedFile, err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	os.Rename("info.txt", "data.txt")

	defer renamedFile.Close()

	// // err = os.Remove("data.txt")
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

	// nf, err := os.Open("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer nf.Close()
	// data, err := ioutil.ReadAll(nf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s \n", data)

	newFile2, err := os.OpenFile("info.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile2.Close()

	writter := bufio.NewWriter(newFile2)
	bs := []byte{97, 98, 99}

	bytesWritten, err := writter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes was written \n", bytesWritten)
	fmt.Printf("%d bytes avalable\n", writter.Available())

	bytesWritten, err = writter.WriteString("\nJust a random string")
	if err != nil {
		log.Fatal(err)
	}
	unflushedBufferSize := writter.Buffered()
	log.Printf("Bytes buffered %d\n", unflushedBufferSize)
	writter.Flush()

	newFile3, err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile3.Close()

	scanner := bufio.NewScanner(newFile3)
	if scanner.Scan() == false {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}
	fmt.Println(scanner.Text())
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	d := []byte("Go is cool!")

	if err := ioutil.WriteFile("info.txt", d, 0644); err != nil {
		log.Fatal(err)
	}
}
