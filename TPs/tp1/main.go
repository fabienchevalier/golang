package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	// Image paths or URLs
	image1 = flag.String("image1", "", "Path to the first image or its URL")
	image2 = flag.String("image2", "", "Path to the second image or its URL")
)

func main() {
	flag.Parse()

	// Initialize the logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log

	// Handle empty flag
	if *image1 == "" || *image2 == "" {
		log.Fatal("Please provide both file paths or URLs using the flags -image1 and -image2")
	}

	// Generate img hash
	hash1, err := hashInput(*image1)
	if err != nil {
		log.Fatalf("Error hashing image 1: %v", err)
	}

	hash2, err := hashInput(*image2)
	if err != nil {
		log.Fatalf("Error hashing image 2: %v", err)
	}

	if hash1 == hash2 {
		log.Println("The images are the same.")
	} else {
		log.Println("The images are different.")
	}
}

func hashInput(input string) (string, error) {
	// Use hashURL if input starts with "http://" or "https://", otherwise use hashFile
	if len(input) > 7 && (input[:7] == "http://" || input[:8] == "https://") {
		log.Println("Downloading image...")
		return hashURL(input)
	}
	return hashFile(input)
}

func hashFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash), nil
}

func hashURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download the image: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash), nil
}
