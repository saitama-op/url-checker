package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// Open the file containing URLs
	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	client := http.Client{
		Timeout: 5 * time.Second, // timeout for requests
	}

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" { // ignore empty lines
			continue
		}

		resp, err := client.Get(url)
		if err != nil {
			fmt.Printf("[ERR] %s ❌ DOWN (%v)\n", url, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Printf("[%d] %s ✅ UP\n", resp.StatusCode, url)
		} else {
			fmt.Printf("[%d] %s ❌ DOWN\n", resp.StatusCode, url)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
