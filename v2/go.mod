module github.com/dsoprea/go-tiff-image-structure/v2

go 1.12

// Development only
// replace github.com/dsoprea/go-exif/v3 => ../../go-exif/v3
// replace github.com/dsoprea/go-utility/v2 => ../../go-utility/v2

require (
	github.com/dsoprea/go-exif/v3 v3.0.0-20221003141534-0e3dba6a88f3
	github.com/dsoprea/go-logging v0.0.0-20200710184922-b02d349568dd
	github.com/dsoprea/go-utility/v2 v2.0.0-20221003141017-b854b1d8773e
	golang.org/x/image v0.0.0-20200618115811-c13761719519
)
