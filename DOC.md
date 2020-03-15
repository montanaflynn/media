

# media
`import "."`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package media is for determining information
about media without decoding the entire file.

Example Usage:


	giphy := "<a href="https://media1.giphy.com/media/l0ErxFClZX9L3bgBi/giphy.gif">https://media1.giphy.com/media/l0ErxFClZX9L3bgBi/giphy.gif</a>"
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

MIT License Copyright (c) 2020 Montana Flynn (<a href="https://montanaflynn.com">https://montanaflynn.com</a>)




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [type BMP](#BMP)
  * [func (b BMP) Size() Dimensions](#BMP.Size)
  * [func (b BMP) Type() ContentType](#BMP.Type)
* [type BMPFormat](#BMPFormat)
* [type ContentType](#ContentType)
  * [func DetectContentType(r io.Reader) (ContentType, []byte, error)](#DetectContentType)
* [type Dimensions](#Dimensions)
* [type GIF](#GIF)
  * [func (b GIF) Size() Dimensions](#GIF.Size)
  * [func (b GIF) Type() ContentType](#GIF.Type)
* [type Media](#Media)
  * [func Parse(r io.Reader) (Media, error)](#Parse)
* [type PNG](#PNG)
  * [func (b PNG) Size() Dimensions](#PNG.Size)
  * [func (b PNG) Type() ContentType](#PNG.Type)


#### <a name="pkg-files">Package files</a>
[bmp.go](/bmp.go) [detect.go](/detect.go) [doc.go](/doc.go) [gif.go](/gif.go) [parse.go](/parse.go) [png.go](/png.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    // PNGType media type
    PNGType ContentType = "image/png"

    // GIFType media type
    GIFType = "image/gif"

    // BMPType media type
    BMPType = "image/bmp"

    // JPEGType media type
    JPEGType = "image/jpeg"
)
```

## <a name="pkg-variables">Variables</a>
``` go
var (
    // ErrUnknownContentType is when media is an unknown type
    ErrUnknownContentType = errors.New("media: unknown media type")

    // ErrUnsupportedSize is when media type doesn't implement size
    ErrUnsupportedSize = errors.New("media: unsupported size")

    // ErrPNGMissingIHDR is when a png is missing the HDR header
    ErrPNGMissingIHDR = errors.New("media: invalid png missing IHDR")

    // ErrBMPInvalidHeaderLength is when a bmp has invalid header length
    ErrBMPInvalidHeaderLength = errors.New("media: invalid bmp header length")
)
```



## <a name="BMP">type</a> [BMP](/bmp.go?s=251:342#L19)
``` go
type BMP struct {
    ContentType ContentType
    Dimensions  Dimensions
    Format      BMPFormat
}

```
BMP holds the BMP










### <a name="BMP.Size">func</a> (BMP) [Size](/bmp.go?s=393:423#L26)
``` go
func (b BMP) Size() Dimensions
```
Size returns the BMP size in width and height




### <a name="BMP.Type">func</a> (BMP) [Type](/bmp.go?s=487:518#L31)
``` go
func (b BMP) Type() ContentType
```
Type returns the BMP content type




## <a name="BMPFormat">type</a> [BMPFormat](/bmp.go?s=84:105#L8)
``` go
type BMPFormat string
```
BMPFormat is the type of the media


``` go
const (
    // OS2Format bmp format
    OS2Format BMPFormat = "OS/2"

    // WindowsFormat bmp format
    WindowsFormat = "Windows"
)
```









## <a name="ContentType">type</a> [ContentType](/parse.go?s=619:642#L23)
``` go
type ContentType string
```
ContentType is the type of the media







### <a name="DetectContentType">func</a> [DetectContentType](/detect.go?s=728:792#L35)
``` go
func DetectContentType(r io.Reader) (ContentType, []byte, error)
```
DetectContentType returns the ContentType from the first 32 bytes





## <a name="Dimensions">type</a> [Dimensions](/parse.go?s=892:942#L40)
``` go
type Dimensions struct {
    Width  int
    Height int
}

```
Dimensions holds the Dimensions










## <a name="GIF">type</a> [GIF](/gif.go?s=67:135#L8)
``` go
type GIF struct {
    ContentType ContentType
    Dimensions  Dimensions
}

```
GIF holds the GIF










### <a name="GIF.Size">func</a> (GIF) [Size](/gif.go?s=186:216#L14)
``` go
func (b GIF) Size() Dimensions
```
Size returns the GIF size in width and height




### <a name="GIF.Type">func</a> (GIF) [Type](/gif.go?s=278:309#L19)
``` go
func (b GIF) Type() ContentType
```
Type returns the GIF media type




## <a name="Media">type</a> [Media](/parse.go?s=1004:1067#L46)
``` go
type Media interface {
    Size() Dimensions
    Type() ContentType
}
```
Media is an interface for getting information from media







### <a name="Parse">func</a> [Parse](/parse.go?s=1138:1176#L52)
``` go
func Parse(r io.Reader) (Media, error)
```
Parse returns the media information including type and dimensions





## <a name="PNG">type</a> [PNG](/png.go?s=67:135#L8)
``` go
type PNG struct {
    ContentType ContentType
    Dimensions  Dimensions
}

```
PNG holds the PNG










### <a name="PNG.Size">func</a> (PNG) [Size](/png.go?s=186:216#L14)
``` go
func (b PNG) Size() Dimensions
```
Size returns the PNG size in width and height




### <a name="PNG.Type">func</a> (PNG) [Type](/png.go?s=280:311#L19)
``` go
func (b PNG) Type() ContentType
```
Type returns the PNG content type








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
