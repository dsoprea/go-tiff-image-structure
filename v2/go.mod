module github.com/dsoprea/go-tiff-image-structure/v2

go 1.12

// Development only
// replace github.com/dsoprea/go-exif/v3 => ../../go-exif/v3
// replace github.com/dsoprea/go-utility/v2 => ../../go-utility/v2

require (
	github.com/dsoprea/go-exif/v3 v3.0.0-20200717071058-9393e7afd446
	github.com/dsoprea/go-logging v0.0.0-20200517223158-a10564966e9d
	github.com/dsoprea/go-utility/v2 v2.0.0-20200717064901-2fccff4aa15e
	golang.org/x/image v0.0.0-20200618115811-c13761719519
)
