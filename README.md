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

MIT License. See the [LICENSE](https://github.com/adtac/relevant-xkcd/blob/master/LICENSE)
file for more details.
