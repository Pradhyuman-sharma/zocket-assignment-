package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func download(url string, filename string, ch chan<- string) {
    // Perform the download
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprintf("Error downloading %s: %s", url, err)
        return
    }
    defer resp.Body.Close()

    // Create the output file
    file, err := os.Create(filename)
    if err != nil {
        ch <- fmt.Sprintf("Error creating file %s: %s", filename, err)
        return
    }
    defer file.Close()

    // Copy the response body to the output file
    _, err = io.Copy(file, resp.Body)
    if err != nil {
        ch <- fmt.Sprintf("Error writing to file %s: %s", filename, err)
        return
    }

    ch <- fmt.Sprintf("Downloaded %s", url)
}

func main() {
    // URLs and filenames to download
    downloads := map[string]string{
        "https://www.example.com/image1.jpg": "image1.jpg",
        "https://www.example.com/image2.jpg": "image2.jpg",
        "https://www.example.com/image3.jpg": "image3.jpg",
    }

    
    ch := make(chan string)

    for url, filename := range downloads {
        go download(url, filename, ch)
    }

  
    for range downloads {
        fmt.Println(<-ch)
    }
}
