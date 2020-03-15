# Media

[![][travis-svg]][travis-url] [![][coveralls-svg]][coveralls-url] [![][goreport-svg]][goreport-url] [![][godoc-svg]][godoc-url] [![][pkggodev-svg]][pkggodev-url] [![][license-svg]][license-url]

A Golang package for determining information about media without decoding the entire file.

Currently it supports finding the dimensions of `png`, `gif` and `bmp` media by only reading 32 bytes.

**Note:** The API is in active development and may change.

- [Install](#install)
- [Example Usage](#Example-Usage)
- [MIT License](#Mit-License)

## Install

```
go get github.com/montanaflynn/media
```

## Example Usage

From [cmd/media-info-http](cmd/media-info-http/main.go):

```go
giphy := "https://media1.giphy.com/media/l0ErxFClZX9L3bgBi/giphy.gif"
r, err := http.Get(giphy)
if err != nil {
	log.Fatal(err)
}
defer r.Body.Close()
m, err := media.Parse(r.Body)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("media dimensions: %v\n", m.Size())
// media dimensions: {480 270}

fmt.Printf("content type: %q\n", m.Type())
// content type: "image/gif"

j, err := json.Marshal(m)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("%s\n", j)
// {"ContentType":"image/gif","Dimensions":{"Width":480,"Height":270}}
```

## MIT License

Copyright (c) 2020 Montana Flynn (https://montanaflynn.com)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORpublicS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

[travis-url]: https://travis-ci.org/montanaflynn/media
[travis-svg]: https://img.shields.io/travis/montanaflynn/media.svg
[coveralls-url]: https://coveralls.io/r/montanaflynn/media?branch=master
[coveralls-svg]: https://img.shields.io/coveralls/montanaflynn/media.svg
[goreport-url]: https://goreportcard.com/report/github.com/montanaflynn/media
[goreport-svg]: https://goreportcard.com/badge/github.com/montanaflynn/media
[godoc-url]: https://godoc.org/github.com/montanaflynn/media
[godoc-svg]: https://godoc.org/github.com/montanaflynn/media?status.svg
[pkggodev-url]: https://pkg.go.dev/github.com/montanaflynn/media
[pkggodev-svg]: https://gistcdn.githack.com/montanaflynn/b02f1d78d8c0de8435895d7e7cd0d473/raw/17f2a5a69f1323ecd42c00e0683655da96d9ecc8/badge.svg
[license-url]: https://github.com/montanaflynn/media/blob/master/LICENSE
[license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
