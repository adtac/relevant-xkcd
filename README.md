# relevant-xkcd

**[See here for a live run.](https://relevant-xkcd.github.io/)**

You know how the old saying goes - *"There's always a relevant XKCD comic regardless of the situation"*.

Inspired by this: [Relevant XKCD](http://thomaspark.co/2017/01/relevant-xkcd/).

### How this works

The lovely folks at [ExplainXKCD](http://www.explainxkcd.com/wiki/index.php/Main_Page)
have a detailed explanation, transcript and everything in between for every XKCD
comic ever. Searching that is probably a good idea to get a relevant XKCD.

### Dependencies

While the server doesn't have any dependencies, the crawler depends on
[goquery](https://github.com/PuerkitoBio/goquery). `go get` that first before
compiling and running the crawler.

### How do I run this?

First get the comics metadata:

```bash
$ go build crawler.go
$ ./crawler
```

That should download the entire list of comics (not the images, just the metadata) and store them in a binary file named `comics.bin`.

Then start the server:

```bash
$ go build main.go
$ ./main
```

This will open a web server on port 8080. You may have to change the `BACKEND` variable in the `index.html`'s javascript to point to your domain and port.

### License

```
Copyright (C) 2017  Adhityaa Chandrasekar

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
```

See the [LICENSE](LICENSE) file for more details.
