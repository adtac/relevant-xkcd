package main

import (
    "fmt"
    "time"
    "net/http"
)

var comics []XKCDComic

func main() {
    loadComics()

    go func() {
        for {
            crawl()
            loadComics()
            time.Sleep(crawlInterval)
        }
    }()
    fmt.Printf("Started background crawl\n");

    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/search", SearchHandler)

    http.ListenAndServe(":8080", nil)
}
