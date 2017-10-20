package main

import (
    "fmt"
    "io/ioutil"
    "bytes"
    "encoding/gob"
)

func stringSliceCount(a string, list []string) int {
    count := 0
    for _, b := range list {
        if b == a {
            count += 1
        }
    }

    return count
}

func loadComics() {
    b, err := ioutil.ReadFile("comics.bin")
    if err != nil {
        fmt.Printf("%v\n", err)
        return
    }

    buf := bytes.NewBuffer(b)
    dec := gob.NewDecoder(buf)
    comics = make([]XKCDComic, 0)
    dec.Decode(&comics)

    fmt.Printf("Loaded %d comics\n", len(comics))
}
