package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	str "strings"
)

// ParseURL get page url and returns file url
func ParseURL(url string) string {
	var fileURL string

	if str.Contains(url, "instagram") {
		fmt.Println("> Instagram")

		if str.Contains(url, "/p/") {
			// Find the review items
			fmt.Println(">> Photo")
		} else if str.Contains(url, "stories") {
			fmt.Println(">> Stories")

			// Photo
			fmt.Println(">>> Photo")

			// Video
			fmt.Println(">>> Video")
		}
	} else if str.Contains(url, "z0r.de") {
		fmt.Println("> z0r")

		// Add ?flash if it's not already there
		if !str.Contains(url, "?flash") {
			url = url + "?flash"
		}

		// Get html from the page
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		html, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading HTML of the page: %v", err)
			return ""
		}
		htmlStr := string(html)

		// Extracting src from the embed tag
		embedStart := str.Index(htmlStr, "<embed")
		if embedStart == -1 {
			return ""
		}

		embedPart := htmlStr[embedStart:]

		srcStart := str.Index(embedPart, `src="`)
		if srcStart == -1 {
			return ""
		}

		srcValue := embedPart[srcStart+5:]

		srcEnd := str.Index(srcValue, `"`)
		if srcEnd == -1 {
			return ""
		}

		fileURL = srcValue[:srcEnd]

		// Convert relative path to absolute
		if str.HasPrefix(fileURL, "../L/") {
			fileURL = "https://z0r.de/L/" + fileURL[5:]
		}
	} else {
		fmt.Println("Nothing to download...")
	}

	return fileURL
}

// DownloadFile will download a url to a local file.
// It's efficient because it will write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var url, fileURL string

	if len(os.Args) > 1 {
		url = os.Args[1]
		fileURL = ParseURL(url)

		if fileURL != "" {
			// Extract the filename from the URL
			fileName := fileURL
			lastSlash := str.LastIndex(fileName, "/")
			if lastSlash != -1 {
				fileName = fileName[lastSlash+1:]
			}

			// Sanitize the filename from invalid characters
			for _, char := range []string{"<", ">", ":", "\"", "/", "\\", "|", "?", "*"} {
				fileName = str.ReplaceAll(fileName, char, "_")
			}

			err := DownloadFile(fileName, fileURL)

			if err != nil {
				panic(err)
			}
		}
	} else {
		log.Fatal("Required context `url` was not specified.")
		return
	}
}
