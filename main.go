package main

import (
    "sort"
    "strings"
    "bytes"
    "encoding/gob"
    "net/http"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type XKCDComic struct {
    Number int           `json:"number"`
    Title string         `json:"title"`
    TitleFields []string `json:"-"`
    URL string           `json:"url"`
    Image string         `json:"image"`
    Date string          `json:"date"`
    Text string          `json:"-"`
}

type SearchResults struct {
    Success bool           `json:"success"`
    Message string         `json:"message"`
    Results []XKCDComic    `json:"results"`
}

var indexHTML []byte
var comics []XKCDComic

var TitleWordWeight int = 400
var TitleIntextWeight int = 200
var TranscriptWeight int = 30
var TextWordWeight int = 8
var TextWeight int = 1

func stringSliceCount(a string, list []string) int {
    count := 0
    for _, b := range list {
        if b == a {
            count += 1
        }
    }

    return count
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s", indexHTML)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
    res := SearchResults{Success: false, Results: make([]XKCDComic, 0)}

    if r.Method == "POST" {
        terms := r.PostFormValue("search")
        for _, term := range strings.Fields(terms) {
            term = strings.ToLower(term)

            type ScoreRecord struct {
                Index int
                Score int
            }

            scores := make([]ScoreRecord, len(comics))

            for i, comic := range comics {
                if comic.Number == 1190 || comic.Number == 1608 || comic.Number == 1037 || comic.Number == 1335 {
                    continue
                }

                if len(res.Results) >= 15 {
                    break
                }

                scores[i].Index = i

                scores[i].Score = TitleWordWeight * stringSliceCount(term, comic.TitleFields)
                scores[i].Score += TitleIntextWeight * strings.Count(strings.ToLower(comic.Title), term)

                lx := strings.Index(comic.Text, "==Transcript==")
                rx := len(comic.Text)
                if lx < 0 {
                    lx = 0
                } else {
                    rx = strings.Index(comic.Text[lx+14:], "==")
                    if rx < 0 {
                        rx = len(comic.Text)
                    } else {
                        rx += lx+14
                    }
                }

                scores[i].Score += TranscriptWeight * strings.Count(comic.Text[lx:rx], term)

                scores[i].Score += TextWordWeight * stringSliceCount(term, strings.Fields(strings.ToLower(comic.Text)))
                scores[i].Score += TextWeight * strings.Count(strings.ToLower(comic.Text), term)
            }

            sort.Slice(scores, func(i, j int) bool {
                return scores[i].Score > scores[j].Score
            })

            for i := 0; i < 15; i++ {
                if scores[i].Score > 0 {
                    res.Results = append(res.Results, comics[scores[i].Index])
                }
            }
        }

        res.Success = true
    } else {
        res.Message = "Only POST requests allowed."
    }

    json, _ := json.Marshal(res)
    fmt.Fprintf(w, "%s", string(json))
}

func main() {
    b, _ := ioutil.ReadFile("comics.bin")
    buf := bytes.NewBuffer(b)
    dec := gob.NewDecoder(buf)
    comics = make([]XKCDComic, 0)
    dec.Decode(&comics)

    fmt.Printf("Loaded %d comics\n", len(comics))

    indexHTML, _ = ioutil.ReadFile("index.html")

    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/search", SearchHandler)

    http.ListenAndServe(":80", nil)
}
