package main

import (
    "net/http"
    "fmt"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, `
<html>

<style>
* {
    font-family: monospace;
    font-size: 14px;
    width: 76ch;
}

.title, .item {
    font-weight: bold;
    padding-bottom: 6px;
}

.title {
    padding-top: 16px;
}

.content, .item-desc {
    margin-left: 4ch;
}

.example {
    padding-top: 4px;
}
</style>

<div class='title'>relevant-xkcd(7)</div>

<div class='title'>NAME</div>
<div class='content'>
relevant-xkcd - a search engine for XKCD comics
</div>

<div class='title'>DESCRIPTION</div>
<div class='content'>
<div class='item'>POST /search</div>
<div class='item-desc'>search for the given term</div>
</div>

<div class='title'>EXAMPLE</div>
<div class='content'>
To search for the term 'velociraptors':
    <div class='item-desc example'>
    $ curl --data 'search=velociraptors' http://adtac.pw:8000/search
    </div>
</div>

<div class='title'>AUTHOR</div>
<div class='content'>
<a href='https://adtac.github.io/'>Adhityaa (@adtac)</a>
</div>

<div class='title'>REPORTING BUGS</div>
<div class='content'>
Please use the <a href='https://github.com/adtac/relevant-xkcd/issues'>Github issue tracker</a> to file bugs. Note that this is feature-complete and will not be accepting new feature suggestions. Feel free to fork the code (while still adhering to the license agreement).
</div>

<div class='title'>COPYRIGHT</div>
<div class='content'>
Copyright (C) 2017  Adhityaa Chandrasekar<br>
<br>
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.<br>
<br>
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.<br>
<br>
You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
</div>

<div class='title'>SEE ALSO</div>
<div class='content'>
<a href='https://relevant-xkcd.github.io/'>A pretty front-end for this API.</a>
</div>

</div>
</html>
`)
}
