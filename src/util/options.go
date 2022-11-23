package util

import (
	"encoding/json"
	"io/ioutil"
)

const (
	OPT_FILE_MODE = 0650
	OPT_INDENT    = "    "
)

type Options struct {
	PixelSize int32 `json:"pixel_size"`
}

var GlobalOptions Options

func LoadOptions(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &GlobalOptions)
}

func SaveOptions(filename string) error {
	data, err := json.MarshalIndent(GlobalOptions, "", OPT_INDENT)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, OPT_FILE_MODE)
}
