package main

import (
    "time"
)

var TitleWordWeight int = 400
var TitleIntextWeight int = 200
var TranscriptWeight int = 30
var TextWordWeight int = 8
var TextWeight int = 1

var URL = "http://www.explainxkcd.com/wiki/index.php?title=List_of_all_comics_(full)&printable=yes"

var crawlInterval = 6 * time.Hour
