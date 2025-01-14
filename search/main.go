package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Replace this with the username you want to check
	var username string
	fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣠⠤⠖⠒⠒⠒⠶⠶⣤⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⣀⡴⠊⠁⠀⠀⣀⣀⠀⠀⠀⠀⠀⠈⠙⠶⣄⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⢀⡼⠋⠀⠀⢠⠞⠉⠀⠀⠀⠀⠀⣀⠀⠀⠀⠀⠈⢳⣄⠀⠀⠀⠀⠀\n⠀⠀⠀⢠⠞⠀⠀⠀⠀⠈⠓⠦⠤⠒⠊⠉⠉⠁⠀⠀⠀⠀⠀⠀⠹⣆⠀⠀⠀⠀\n⠀⠀⢠⠏⠀⠀⠀⠀⠀⠀⣠⡀⠀⠀⠀⢀⣀⠀⠀⠀⠀⠀⠀⠀⠀⢻⡄⠀⠀⠀\n⠀⢀⡟⠀⠀⠀⠀⠀⠀⠘⠿⢃⡜⠉⠁⠈⠉⠀⠀⠀⠀⠀⠀⠀⠀⠈⣧⠀⠀⠀\n⠀⣸⠁⠀⠀⠀⠀⠀⠀⠇⠀⠸⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⢦⣄⠀\n⠀⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣧\n⣰⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼\n⣿⢹⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⠴⠞⠁\n⠈⢿⣇⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠇⠀⠀⠀\n⠀⠀⠹⣦⠀⠀⠈⠳⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⠴⠀⠀⠀⢠⠟⠀⠀⠀⠀\n⠀⠀⠀⠈⠳⣄⠀⠀⠀⠉⠉⠛⠛⠒⠒⠛⠛⠋⠉⠁⠀⠀⣀⡴⠋⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠈⠛⠶⣤⣄⣀⡀⠀⠀⠀⠀⠀⢀⣀⣤⠴⠞⠋⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠙⡿⠿⠿⠿⠿⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣴⡇⠀⣠⡴⠚⢛⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡟⠀⡈⠻⠋⠁⢧⡴⠊⢳⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⠟⠋⠀⠀⠀⠀⠀⠀⠀⠀⢧⠀⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠀⣰⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣇⠀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⣰⢻⠁⠀⠀⠀⠀⠀⠀⠀⣰⠀⠀⠀⢹⡀⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⢠⣿⣧⠀⢀⣤⣆⠀⠀⢀⣾⣿⠀⠀⠀⠈⣇⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⢸⢱⣿⣶⣿⣿⣿⣷⣠⣾⣿⣿⡟⠋⠙⠛⢿⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⣆⠟⣿⣿⡟⠙⢿⣿⣿⣿⠟⢻⡇⠀⠀⠀⠸⣧⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⢸⣿⡀⠸⣿⠀⠀⠈⢿⣿⠃⠀⠀⣧⢠⡀⢤⣸⣿⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠸⢿⡇⠀⠁⠀⠀⠀⠀⠀⠀⠀⠀⠙⠚⠓⠛⣉⣹⡇⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⢸⣷⣤⣤⣤⣀⣀⣀⣠⣤⣤⣤⣶⣶⣾⣿⣿⣿⡇⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠿⠃⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⠀⠉⠉⠉⢙⣏⠉⠉⠉⢹⣿⣯⣄⣀⣀⣸⡇⠀⠀⠀⠀⠀⠀⠀\n⠀⠀⠀⠀⠀⠀⣀⣠⣤⣴⣾⣿⣭⣭⢿⣿⢿⣷⠦⣬⠭⠭⠿⠒⠤⣀⠀⠀⠀⠀\n⠀⠀⠀⣴⠞⠛⠉⠉⠙⢻⡿⠛⠉⠀⠀⠈⠓⠟⠓⠋⠀⠀⢀⣀⣀⣬⠇⠀⠀⠀\n⠀⠀⠘⠿⠶⠶⠶⠶⠒⠚⠻⠶⢶⠶⠶⠶⠶⠶⠛⠛⠛⠉⠉⠉⠀⠀⠀⠀⠀⠀\n")
	fmt.Println("paste name here:")
	fmt.Scanln(&username)
	

	// Social media sites to check
	sites := map[string]string{
		"Twitter":  "https://twitter.com/%s",
		"Instagram": "https://www.instagram.com/%s",
		"GitHub":   "https://github.com/%s",
		"Youtube": "https://www.youtube.com/%s",
	}

	for site, url := range sites {
		checkUsername(site, fmt.Sprintf(url, username))
	}
}

func checkUsername(site, url string) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Checking %s at %s...\n", site, r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == http.StatusOK {
			fmt.Printf("\033[32m{FOUND}\033[0m %s has an account at %s\n", site, r.Request.URL)
		} else {
			fmt.Printf("\033[31m{NOT FOUND}\033[0m %s does not have an account at %s\n", site, r.Request.URL)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("\033[31mERROR\033[0m Failed to check %s: %v\n", r.Request.URL, err)
	})

	// Start scraping
	err := c.Visit(url)
	if err != nil {
		log.Printf("\033[31mERROR\033[0m Visit failed: %v\n", err)
	}
}

