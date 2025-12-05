package valute

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
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
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

func ConvertValutesToJSONBytes(valutes []Valute) ([]byte, error) {
	result := make([]map[string]any, 0, len(valutes))

	for _, valute := range valutes {
		value, err := ParseValue(valute.Value)
		if err != nil {
			return nil, err
		}

		if valute.NumCode != "" {
			if _, err := strconv.ParseInt(valute.NumCode, 10, 64); err != nil {
				return nil, InvalidNumCodeError{NumCode: valute.NumCode, Valute: valute}
			}
		}

		numCodeInt := int64(0)

		if valute.NumCode != "" {
			numCodeInt, err = strconv.ParseInt(valute.NumCode, 10, 64)
			if err != nil {
				return nil, InvalidNumCodeError{NumCode: valute.NumCode, Valute: valute}
			}
		}

		valuteMap := map[string]any{
			"num_code":  numCodeInt,
			"char_code": valute.CharCode,
			"value":     value,
		}

		result = append(result, valuteMap)
	}

	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, FailedEncodeError{FilePath: ""}
	}

	return jsonBytes, nil
}

func SaveJSONBytes(jsonBytes []byte, outputPath string) error {
	const dirPerm = 0o755

	if err := os.MkdirAll(filepath.Dir(outputPath), dirPerm); err != nil {
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

	if _, err := jsonFile.Write(jsonBytes); err != nil {
		return FailedEncodeError{FilePath: outputPath}
	}

	return nil
}

func ParseValue(toConvert string) (float64, error) {
	toConvert = strings.Replace(toConvert, ",", ".", 1)

	val, err := strconv.ParseFloat(toConvert, 64)
	if err != nil {
		return 0, ConversionError{Value: toConvert, Err: err}
	}

	return val, nil
}
