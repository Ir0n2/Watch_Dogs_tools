package main

import (
	//"crypto/md5"
	//"crypto/sha1"
	//"crypto/sha256"
	//"crypto/sha512"
	"fmt"
	"regexp"
)

// HashType identifies and returns the type of the hash based on its length and pattern
func HashType(hash string) string {
	switch len(hash) {
	case 32:
		if match, _ := regexp.MatchString("^[a-fA-F0-9]{32}$", hash); match {
			return "MD5"
		}
	case 40:
		if match, _ := regexp.MatchString("^[a-fA-F0-9]{40}$", hash); match {
			return "SHA-1"
		}
	case 64:
		if match, _ := regexp.MatchString("^[a-fA-F0-9]{64}$", hash); match {
			return "SHA-256"
		}
	case 96:
		if match, _ := regexp.MatchString("^[a-fA-F0-9]{96}$", hash); match {
			return "SHA-384"
		}
	case 128:
		if match, _ := regexp.MatchString("^[a-fA-F0-9]{128}$", hash); match {
			return "SHA-512"
		}
	default:
		return "Unknown hash type"
	}
	return "Unknown hash type"
}

// main function to test the hash identification
func main() {
	// Sample hashes for testing
	/*hashes := []string{
		"5d41402abc4b2a76b9719d911017c592",
		"bd7b3348c1edb97305b1cd53256f1ed2",// MD5
		"2fd4e1c67a2d28fced849ee1bb76e7391b93eb12", // SHA-1
		"9e107d9d372bb6826bd81d3542a419d6e0a5e6da3dfe7a2b3c2b927f914bd205", // SHA-256
		"cb00753f45a35e8bb5a03d699ac65007272c32ab0eded1631a8b605a43ff5bed8086072ba1e7cc2358baeca134c825a7", // SHA-512
	}*/
	var hash string
	fmt.Println("Paste Hash Here:")
	fmt.Scanln(&hash)
	fmt.Printf("Hash: %s\nType: %s\n\n", hash, HashType(hash))
}

