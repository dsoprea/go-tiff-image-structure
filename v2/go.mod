module github.com/dsoprea/go-tiff-image-structure/v2

go 1.12

// Development only
// replace github.com/dsoprea/go-exif/v3 => ../../go-exif/v3
// replace github.com/dsoprea/go-utility/v2 => ../../go-utility/v2

require (
	github.com/dsoprea/go-exif/v2 v2.0.0-20200321225314-640175a69fe4 // indirect
	github.com/dsoprea/go-exif/v3 v3.0.0-20221003143021-0a0262e4b8b6
	github.com/dsoprea/go-logging v0.0.0-20200710184922-b02d349568dd
	github.com/dsoprea/go-utility/v2 v2.0.0-20221003142440-7a1927d49d9d
	github.com/jessevdk/go-flags v1.5.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 // indirect
	golang.org/x/image v0.0.0-20220902085622-e7cb96979f69
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e // indirect
)
