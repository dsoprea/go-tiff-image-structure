package tiffstructure

import (
    "path"

    "go/build"

    "github.com/dsoprea/go-logging"
)

var (
    assetsPath = ""
)

// getModuleRootPath returns our source-path when running from source during
// tests.
func getModuleRootPath() string {
    p, err := build.Default.Import(
        "github.com/dsoprea/go-tiff-image-structure",
        build.Default.GOPATH,
        build.FindOnly)

    log.PanicIf(err)

    packagePath := p.Dir
    return packagePath
}

// getTestAssetsPath returns the test-asset path.
func getTestAssetsPath() string {
    moduleRootPath := getModuleRootPath()
    assetsPath := path.Join(moduleRootPath, "assets")

    return assetsPath
}

// getTestExifImageFilepath returns the file-path of the default EXIF image
// asset.
func getTestExifImageFilepath() string {
    return path.Join(assetsPath, "file_example_TIFF_1MB-scaled.tiff")
}

func init() {
    assetsPath = getTestAssetsPath()
}
