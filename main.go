package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

type Result struct {
	URL    string `json:"url"`
	Status string `json:"status"`
	Code   int    `json:"code"`
	Err    string `json:"error,omitempty"`
}

const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[31m"
	ColorGreen = "\033[32m"
)

func worker(id int, urls <-chan string, results chan<- Result, client *http.Client) {
	for url := range urls {
		resp, err := client.Get(url)
		if err != nil {
			results <- Result{URL: url, Status: "DOWN", Err: err.Error()}
			continue
		}
		resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			results <- Result{URL: url, Status: "UP", Code: resp.StatusCode}
		} else {
			results <- Result{URL: url, Status: "DOWN", Code: resp.StatusCode}
		}
	}
}

func main() {
	// Command-line flags
	numWorkers := flag.Int("workers", 20, "Number of concurrent workers")
	urlFile := flag.String("file", "urls.txt", "File containing URLs to check")
	csvFile := flag.String("csv", "results.csv", "Output CSV file")
	jsonFile := flag.String("json", "", "Optional JSON output file")
	flag.Parse()

	// Open URL file
	file, err := os.Open(*urlFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	urlList := []string{}
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url != "" {
			urlList = append(urlList, url)
		}
	}

	if len(urlList) == 0 {
		fmt.Println("No URLs found in", *urlFile)
		return
	}

	urls := make(chan string, *numWorkers)
	results := make(chan Result, len(urlList))
	client := &http.Client{Timeout: 5 * time.Second}

	// Start worker pool
	for i := 0; i < *numWorkers; i++ {
		go worker(i, urls, results, client)
	}

	// Send URLs to workers
	for _, url := range urlList {
		urls <- url
	}
	close(urls)

	// Collect all results
	allResults := []Result{}
	for i := 0; i < len(urlList); i++ {
		res := <-results
		allResults = append(allResults, res)
	}

	// Sort: Failed (DOWN) first
	sort.Slice(allResults, func(i, j int) bool {
		if allResults[i].Status == "DOWN" && allResults[j].Status == "UP" {
			return true
		}
		if allResults[i].Status == "UP" && allResults[j].Status == "DOWN" {
			return false
		}
		return allResults[i].URL < allResults[j].URL
	})

	// Print table with colors
	fmt.Printf("%-50s %-8s %-6s %-s\n", "URL", "Status", "Code", "Error")
	fmt.Println(strings.Repeat("-", 90))
	for _, res := range allResults {
		statusColor := ColorGreen
		if res.Status == "DOWN" {
			statusColor = ColorRed
		}
		fmt.Printf("%-50s %s%-8s%s %-6d %-s\n", res.URL, statusColor, res.Status, ColorReset, res.Code, res.Err)
	}

	// Save CSV
	csvF, err := os.Create(*csvFile)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvF.Close()

	writer := csv.NewWriter(csvF)
	defer writer.Flush()

	writer.Write([]string{"URL", "Status", "Code", "Error"})
	for _, res := range allResults {
		codeStr := fmt.Sprintf("%d", res.Code)
		writer.Write([]string{res.URL, res.Status, codeStr, res.Err})
	}
	fmt.Println("Results saved to", *csvFile)

	// Save JSON if specified
	if *jsonFile != "" {
		jsonF, err := os.Create(*jsonFile)
		if err != nil {
			fmt.Println("Error creating JSON file:", err)
			return
		}
		defer jsonF.Close()

		encoder := json.NewEncoder(jsonF)
		encoder.SetIndent("", "  ")
		err = encoder.Encode(allResults)
		if err != nil {
			fmt.Println("Error writing JSON:", err)
			return
		}
		fmt.Println("Results saved to", *jsonFile)
	}
}
