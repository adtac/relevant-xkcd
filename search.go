package main

import (
    "net/http"
    "encoding/json"
    "strings"
    "sort"
    "fmt"
)

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
                if comic.Number == 1190 || comic.Number == 1608 || comic.Number == 1037 || comic.Number == 1335 || comic.Number == 1663 {
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

            maxResults := 15
            if len(scores) < maxResults {
                maxResults = len(scores)
            }

            for i := 0; i < maxResults; i++ {
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
    fmt.Fprintf(w, "%s", string(json))
}
