package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID        string `xml:"ID,attr"`
	NumCode   string `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   string `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}

type ValuteJSON struct {
	NumCode  int64   `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type InvalidNumCodeError struct {
	NumCode string
	Valute  Valute
}

func (e InvalidNumCodeError) Error() string {
	return fmt.Sprintf("invalid NumCode '%s' for element: %+v", e.NumCode, e.Valute)
}

type FailedFileOpenError struct {
	FilePath string
}

func (e FailedFileOpenError) Error() string {
	return "no such file or directory: " + e.FilePath
}

type FailedFileCloseError struct {
	FilePath string
}

type FailedCreateFileError struct {
	FilePath string
}

func (e FailedCreateFileError) Error() string {
	return "failed to create file: " + e.FilePath
}

type FailedCreateDirsError struct {
	DirPath string
}

func (e FailedCreateDirsError) Error() string {
	return "failed to create directories: " + e.DirPath
}

func (e FailedFileCloseError) Error() string {
	return "failed to close file: " + e.FilePath
}

type XMLDecodeError struct {
	FilePath string
	Err      error
}

func (e XMLDecodeError) Error() string {
	return fmt.Sprintf("failed to decode XML file '%s': %v", e.FilePath, e.Err)
}

type FailedEncodeError struct {
	FilePath string
}

func (e FailedEncodeError) Error() string {
	return "failed to encode JSON file: " + e.FilePath
}

type ConversionError struct {
	Value string
	Err   error
}

func (e ConversionError) Error() string {
	return "failed to convert valute '" + e.Value + "': " + e.Err.Error()
}

func ParseValuteXML(path string) (ValCurs, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return ValCurs{}, FailedFileOpenError{FilePath: path}
	}

	defer func() {
		if err := xmlFile.Close(); err != nil {
			panic(FailedFileCloseError{FilePath: path})
		}
	}()

	var valCurs ValCurs

	decoder := xml.NewDecoder(xmlFile)

	decoder.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		if encoding == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	if err := decoder.Decode(&valCurs); err != nil {
		return ValCurs{}, XMLDecodeError{FilePath: path, Err: err}
	}

	return valCurs, nil
}

func ConvertValutesToJSON(valutes []Valute) ([]ValuteJSON, error) {
	valutesJSON := make([]ValuteJSON, 0, len(valutes))

	for _, valute := range valutes {
		value, err := parseValue(valute.Value)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Parsing %+v\n", valute)

		var numCode int64
		if valute.NumCode == "" {
			numCode = 0
		} else {
			var err error

			numCode, err = strconv.ParseInt(valute.NumCode, 10, 64)
			if err != nil {
				return nil, InvalidNumCodeError{NumCode: valute.NumCode, Valute: valute}
			}
		}

		valutesJSON = append(valutesJSON, ValuteJSON{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	return valutesJSON, nil
}

func SaveToJSON(valutesJSON []ValuteJSON, outputPath string) error {
	var err error

	const dirPerm = 0o755

	err = os.MkdirAll(filepath.Dir(outputPath), dirPerm)
	if err != nil {
		return FailedCreateDirsError{DirPath: filepath.Dir(outputPath)}
	}

	jsonFile, err := os.Create(outputPath)
	if err != nil {
		return FailedCreateFileError{FilePath: outputPath}
	}

	defer func() {
		if err := jsonFile.Close(); err != nil {
			panic(FailedFileCloseError{FilePath: outputPath})
		}
	}()

	encoder := json.NewEncoder(jsonFile)

	encoder.SetIndent("", "  ")

	if err := encoder.Encode(valutesJSON); err != nil {
		return FailedEncodeError{FilePath: outputPath}
	}

	return nil
}

func parseValue(toConvert string) (float64, error) {
	toConvert = strings.Replace(toConvert, ",", ".", 1)

	// InvalidNumCodeError
	val, err := strconv.ParseFloat(toConvert, 64)
	if err != nil {
		return 0, ConversionError{Value: toConvert, Err: err}
	}

	return val, nil
}
