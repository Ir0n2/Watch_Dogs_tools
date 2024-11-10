package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
)

//var md5hash = "77f62e3524cd583d698d51fa24fdff4f"
//var sha256hash = "95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced"

// HashWithSHA1 hashes the input string using SHA-1 and returns the hash in hexadecimal format
func HashWithSHA1(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

// HashWithSHA512 hashes the input string using SHA-512 and returns the hash in hexadecimal format
func HashWithSHA512(input string) string {
	hasher := sha512.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Figlet(text string, font string) (string) {
	// Run `figlet` with the specified text and font
	cmd := exec.Command("figlet", "-f", font, text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
		//return "", fmt.Errorf("figlet command failed: %w", err)
	}
	return string(output)
}
func main() {
	
	var sel string 
	var wordlist string
	var md5hash string
	var sha1hash string
	var sha256hash string
	var sha512hash string
	
	fmt.Println(Figlet("HashCracker", "slant"))

	fmt.Println("1: MD5\n2: SHA256\n3: SHA1\n4: SHA512")
	fmt.Scanln(&sel)
	
	if sel == "1" {
		fmt.Println("md5 hash:")
		fmt.Scanln(&md5hash)
	} else if sel == "2"{
		fmt.Println("sha256 hash:")
		fmt.Scanln(&sha256hash)
	} else if sel == "3" {
		fmt.Println("sha1 hash:")
		fmt.Scanln(&sha1hash)
	} else if sel == "4" {
		fmt.Println("sha512 hash:")
		fmt.Scanln(&sha512hash)
	} else {
		return
	}

	fmt.Println("Please paste wordlist url")
	fmt.Scanln(&wordlist)


	f, err := os.Open(wordlist)
	if err != nil {
		log.Fatalln(err)
	}

	//defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		password := scanner.Text()

		switch sel {
			case "1":
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == md5hash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
			}
			case "2":
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sha256hash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
			}
			case "3":
                        hash := HashWithSHA1(password)
                        if hash == sha1hash {
                                fmt.Printf("[+] Password found (SHA-1): %s\n", password)

			}
			case "4":
                        hash := HashWithSHA512(password)
                        if hash == sha512hash {
                                fmt.Printf("[+] Password found (SHA-1): %s\n", password)
			}
                       

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("nothing found!")
		log.Fatalln(err)
	}
	//fmt.Println("Couldn't crack it boss...")
}
