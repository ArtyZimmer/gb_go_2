package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	createFile()
}

func createFile() {
	var newNameOfFile string
	for i := 0; i < 1000000; i++ {
		a := strconv.Itoa(i)
		newNameOfFile = "emptyFile_" + a
		file, _ := os.Create(newNameOfFile)
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				file.Close()
			}
		}()
	}
}
