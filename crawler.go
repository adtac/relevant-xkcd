package main

import (
    "strings"
    "bytes"
    "encoding/gob"
    "os"
    "strconv"

    "github.com/PuerkitoBio/goquery"
)

func crawl() {
    doc, _ := goquery.NewDocument(URL)

    tmp_comics := make([]XKCDComic, 0)

    doc.Find("tr").Each(func(i int, row *goquery.Selection) {
        comic := XKCDComic{}
        explanationURL := ""

        row.Find("td").Each(func(j int, col *goquery.Selection) {
            text := strings.TrimSpace(col.Text())

            switch j {
                case 0:
                    comic.URL = text
                    comic.Number, _ = strconv.Atoi(text[strings.Index(text, "/")+1:])

                case 1:
                    comic.Title = strings.TrimSpace(text[:strings.Index(text, "(create)")-1])
                    comic.TitleFields = strings.Fields(comic.Title)

                    explanationURL, _ = col.Find("a").Attr("href")
                    explanationURL = "http://www.explainxkcd.com" + explanationURL[:15] + "?action=edit&title=" + explanationURL[16:]

                    exp, err := goquery.NewDocument(explanationURL)
                    if err == nil {
                        comic.Text = exp.Find("textarea").Text()
                    }

                case 3:
                    comic.Image = "https://imgs.xkcd.com/comics/" + strings.Replace(text, " ", "_", -1)

                case 4:
                    comic.Date = text
            }
        })

        index := strings.Index(comic.Text, "titletext = ")
        if index > 0 {
            comic.TitleText = comic.Text[index+12:]
            comic.TitleText = comic.TitleText[:strings.Index(comic.TitleText, "}")-1]
        }

        tmp_comics = append(tmp_comics, comic)

        var buf bytes.Buffer
        enc := gob.NewEncoder(&buf)
        enc.Encode(tmp_comics)

        f, _ := os.Create("comics.bin")
        f.Write(buf.Bytes())
        f.Close()
    })

    comics = tmp_comics
}
