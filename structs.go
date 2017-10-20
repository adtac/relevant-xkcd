package main

type XKCDComic struct {
    Number int           `json:"number"`
    Title string         `json:"title"`
    TitleText string     `json:"titletext"`
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
