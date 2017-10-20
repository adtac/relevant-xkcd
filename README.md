## relevant-xkcd

<p align="center">You know how the old saying goes - "There's an XKCD for that!"</p>

<p align="center"><a href="https://relevant-xkcd.github.io/">
Check out a hosted version of the project here.
</a></p>

<p align="center"><img src="https://user-images.githubusercontent.com/7521600/31832568-1a3e39ee-b5e5-11e7-901f-b57f69e82c68.png" alt="XKCD Search Engine"></p>

### What is this?

This is a search engine + crawler backend for XKCD comics. There's a simple API
to search for comics. This will search through titles, transcript, and discussion.
The comics list is automatically refreshed in the background every 6 hours (even
though Randall updates less frequently - every Monday, Wednesday, and Friday).

### How this works

The lovely folks at [ExplainXKCD](http://www.explainxkcd.com/wiki/index.php/Main_Page)
have a detailed explanation, transcript and everything in between for every XKCD
comic ever. Searching that is probably a good idea to get a relevant XKCD.

### Installation

The server is dockerized. To get an instance running on port 8080 of your host
machine, simply do:

```bash
$ docker pull adtac/relevant-xkcd:latest
$ docker run -d -p 8080:8080 adtac/relevant-xkcd:latest
```

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
