# relevant-xkcd

You know how the old saying goes - *"There's always a relevant XKCD comic regardless of the situation"*.

Inspired by this: [Relevant XKCD](http://thomaspark.co/2017/01/relevant-xkcd/).

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
    Copyright 2017 Adhityaa Chandrasekar

    Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
