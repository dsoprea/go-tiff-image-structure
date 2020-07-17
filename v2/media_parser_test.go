package tiffstructure

import (
	"fmt"
	"os"
	"testing"

	"io/ioutil"

	"github.com/dsoprea/go-logging"
)

func TestNewTiffMediaParser(t *testing.T) {
	tmp := NewTiffMediaParser()
	if tmp == nil {
		t.Fatalf("NewTiffMediaParser() returned nil.")
	}
}

func TestTiffMediaParser_Parse(t *testing.T) {
	filepath := getTestExifImageFilepath()

	tmp := NewTiffMediaParser()

	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	stat, err := f.Stat()
	log.PanicIf(err)

	size := stat.Size()

	_, err = tmp.Parse(f, int(size))
	log.PanicIf(err)
}

func TestTiffMediaParser_ParseFile(t *testing.T) {
	filepath := getTestExifImageFilepath()

	tmp := NewTiffMediaParser()

	_, err := tmp.ParseFile(filepath)
	log.PanicIf(err)
}

func TestTiffMediaParser_ParseBytes(t *testing.T) {
	filepath := getTestExifImageFilepath()

	tmp := NewTiffMediaParser()

	imageBytes, err := ioutil.ReadFile(filepath)
	log.PanicIf(err)

	_, err = tmp.ParseBytes(imageBytes)
	log.PanicIf(err)
}

func TestTiffMediaParser_LooksLikeFormat(t *testing.T) {
	filepath := getTestExifImageFilepath()

	data, err := ioutil.ReadFile(filepath)
	log.PanicIf(err)

	pmp := NewTiffMediaParser()

	if pmp.LooksLikeFormat(data) != true {
		t.Fatalf("not detected as png")
	}
}

func ExampleTiffMediaParser_LooksLikeFormat() {
	filepath := getTestExifImageFilepath()

	data, err := ioutil.ReadFile(filepath)
	log.PanicIf(err)

	pmp := NewTiffMediaParser()

	isPng := pmp.LooksLikeFormat(data)
	fmt.Printf("%v\n", isPng)

	// Output:
	// true
}
