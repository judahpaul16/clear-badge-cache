package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

const (
	defaultFILE = ""
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		help()
		return
	}

	repoURL := getRepoURLFromArgs()
	badgeFile := getBadgeFileFromArgs()

	validateURL := regexp.MustCompile(`^https?:\/\/(www\.)?github\.com\/[\w-]+\/[\w-]+(?:\.git)?$`)
	if !validateURL.MatchString(repoURL) {
		log.Fatalf("Invalid repository URL: %s, Must be a GitHub repository URL.\n", repoURL)
	}

	fmt.Println("")
	fmt.Printf("Purging cached badge images for %s%s...\n", repoURL, badgeFile)

	urls, err := getBadgeURLs(repoURL, badgeFile)
	if err != nil {
		log.Fatalf("Failed to retrieve badge URLs from %s%s\n%v\n", repoURL, badgeFile, err)
	}
	if len(urls) == 0 {
		log.Fatalf("No badge URLs found to purge.\n")
	}
	purgeBadges(urls)
}

func getBadgeURLs(repoURL, badgeFile string) ([]string, error) {
	var badgeURLs []string
	fullURL := repoURL + badgeFile
	fmt.Println("Requesting badges from:", fullURL)

	client := &http.Client{}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("User-Agent", "KLUDGE (Linux;)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	scanner := bufio.NewScanner(resp.Body)
	urlRegex := regexp.MustCompile(`https://camo\.githubusercontent\.com/[^\s"]+`)
	for scanner.Scan() {
		line := scanner.Text()
		urls := urlRegex.FindAllString(line, -1)
		if urls != nil {
			badgeURLs = append(badgeURLs, urls...)
		}
	}
	return badgeURLs, nil
}

func getRepoURLFromArgs() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}

	// If no argument provided, prompt for URL
	fmt.Print("Enter repository URL: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func getBadgeFileFromArgs() string {
	if len(os.Args) > 2 {
		return os.Args[2]
	}
	return defaultFILE
}

func help() {
	fmt.Println("usage: clear-badge-cache [ -h | --help | [ repoURL [ badgeFile ] ] ]")
	fmt.Printf("\n       repoURL   required unless provided on command line\n       badgeFile defaults to '%s'\n\n", defaultFILE)
	os.Exit(0)
}

func purgeBadges(urls []string) {
	userAgentString := "KLUDGE (Linux;) (please fix github tickets 224 116 111 218 414 etc)"
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	for _, url := range urls {
		req, err := http.NewRequest("PURGE", url, nil)
		if err != nil {
			fmt.Printf("Failed to create request for %s: %v\n", url, err)
			continue
		}
		req.Header.Set("User-Agent", userAgentString)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error purging %s: %v\n", url, err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close() // Close immediately instead of deferring
		if err != nil {
			fmt.Printf("Failed to read response for %s: %v\n", url, err)
			continue
		}
		fmt.Printf("Purged Badge: %s - Status: %s\n", url, body)
	}
}
