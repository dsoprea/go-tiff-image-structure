package main

import (
	"fmt"
	"os"

	"encoding/json"
	"io/ioutil"

	"github.com/jessevdk/go-flags"

	"github.com/dsoprea/go-exif/v3"
	"github.com/dsoprea/go-logging/v2"
	"github.com/dsoprea/go-tiff-image-structure/v2"
)

var (
	arguments = &struct {
		Filepath string `short:"f" long:"filepath" required:"true" description:"File-path of image ('-' for STDIN)"`

		// Currently this is the only way of printing, but users that require
		// JSON should pass it for forward compatibility
		Json bool `short:"j" long:"json" description:"Print as JSON"`

		DoPrintVerbose       bool `short:"v" long:"verbose" description:"Print logging"`
		DoUniversalTagSearch bool `short:"u" long:"universal-tags" description:"If tags not found in known mapped IFDs, fallback to trying all IFDs."`
	}{}
)

func main() {
	defer func() {
		if errRaw := recover(); errRaw != nil {
			err := errRaw.(error)
			log.PrintError(err)

			os.Exit(-2)
		}
	}()

	_, err := flags.Parse(arguments)
	if err != nil {
		os.Exit(-1)
	}

	if arguments.DoPrintVerbose == true {
		cla := log.NewConsoleLogAdapter()
		log.AddAdapter("console", cla)

		scp := log.NewStaticConfigurationProvider()
		scp.SetLevel(log.LevelDebug)

		log.LoadConfiguration(scp)
	}

	// Read

	var data []byte
	if arguments.Filepath == "-" {
		var err error
		data, err = ioutil.ReadAll(os.Stdin)
		log.PanicIf(err)
	} else {
		var err error
		data, err = ioutil.ReadFile(arguments.Filepath)
		log.PanicIf(err)
	}

	// Parse

	tmp := tiffstructure.NewTiffMediaParser()
	intfc, parseErr := tmp.ParseBytes(data)

	if intfc == nil {
		if parseErr == nil {
			// We should never get a `nil` `intfc` value back *and* a `nil`
			// `parseErr`.
			log.Panicf("Could not parse.")
		} else {
			log.Panic(parseErr)
		}
	}

	// Export

	_, rawExif, err := intfc.Exif()
	if err != nil {
		if log.Is(err, exif.ErrNoExif) == true {
			fmt.Println("No EXIF data found.")
			fmt.Println("")

			os.Exit(-3)
		} else {
			log.Panic(err)
		}
	}

	entries, _, err := exif.GetFlatExifDataUniversalSearch(rawExif, nil, arguments.DoUniversalTagSearch)
	log.PanicIf(err)

	// Print

	raw, err := json.MarshalIndent(entries, "  ", "  ")
	log.PanicIf(err)

	fmt.Println(string(raw))
}
