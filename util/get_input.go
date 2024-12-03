package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	sessionCookie = "53616c7465645f5f6c958025c3a9be346a5bd61447fb3e6d51484b522f6b6b6b00e7f9ec1e850ad994734b7e9169fdb24dd3b9a3d6821baa111cd96726887913" // Replace with your actual session cookie
	year          = "2024"                                                                                                                             // Change to the current year as needed
)

func main() {
	currentDay := time.Now().Day()
	dayFormatted := fmt.Sprintf("%02d", currentDay)

	// Create a directory for the day
	dayDir := filepath.Join(".", "day"+dayFormatted)
	err := os.MkdirAll(dayDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Download the input file
	if err := downloadInput(dayDir, currentDay); err != nil {
		fmt.Println("Error downloading input:", err)
		return
	}

	// Download the task description
	if err := downloadTask(dayDir, currentDay); err != nil {
		fmt.Println("Error downloading task description:", err)
		return
	}

	// Create a Go template file
	if err := createGoTemplate(dayDir, dayFormatted); err != nil {
		fmt.Println("Error creating Go template:", err)
		return
	}

	fmt.Println("Setup for Day", dayFormatted, "complete.")
}

func downloadInput(dayDir string, day int) error {
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%d/input", year, day)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", "session="+sessionCookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(dayDir, "input.txt"), body, 0644)
}

func downloadTask(dayDir string, day int) error {
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%d", year, day)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Save the HTML content directly
	return ioutil.WriteFile(filepath.Join(dayDir, "task.html"), body, 0644)
}

func createGoTemplate(dayDir string, day string) error {
	goTemplate := fmt.Sprintf(`package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    fmt.Println("Advent of Code - Day %s")

    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening input file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // Process each line of input
        fmt.Println(line)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input file:", err)
    }
}
`, day)

	return ioutil.WriteFile(filepath.Join(dayDir, "solution.go"), []byte(goTemplate), 0644)
}
