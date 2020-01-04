package main

import (
	"crypto/sha256"
	"log"
	"os"
	"fmt"
)

// Automate the creation of multiple directories
func createDirs(dirNum int, uDir string) {

	// Change into a directory, specified by the user
	os.Chdir(uDir)

	for i := 1; i < dirNum + 1; i++ {
		
		fmt.Print("Enter a directory name: ")
		var uDirName string
		fmt.Scanf("%s", &uDirName)
		os.Mkdir(uDirName, 0666)

	}

}

// Calculate a SHA256 file checksum
func getFileHash(fileLocation string)  {
	
	f, err := os.Open(fileLocation)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	digest := sha256.Sum256([]byte(fileLocation))

	fmt.Printf("%x", digest)

}

