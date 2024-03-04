package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
)

var (
	// Image paths
	image1Path = flag.String("image1", "", "Path to the first image")
	image2Path = flag.String("image2", "", "Path to the second image")
)

func main() {
	flag.Parse()

	// Handle empty flag
	if *image1Path == "" || *image2Path == "" {
		fmt.Println("Please provide both file paths using the flags -file1 and -file2")
		return
	}

	// Generate img hash
	hash1, err := hashFile(*image1Path)
	if err != nil {
		fmt.Println(err)
		return
	}

	hash2, err := hashFile(*image2Path)
	if err != nil {
		fmt.Println(err)
		return
	}

	if hash1 == hash2 {
		fmt.Println("The images are the same.")
	} else {
		fmt.Println("The images are different.")
	}
}

func hashFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash), nil
}
