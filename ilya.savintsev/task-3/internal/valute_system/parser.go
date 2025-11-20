package valsys

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"os"

	"golang.org/x/net/html/charset"
)

var (
	errOpenXML = errors.New("no such file or directory")
	errReadXML = errors.New("invalid file")
	errDecdXML = errors.New("invalid encoding")
)

func ParseXML(filepath string) (*ValCurs, error) {
	valuteCurs, err := os.Open(filepath)
	if err != nil {
		return nil, errOpenXML
	}

	defer func() {
		panicIfErr(valuteCurs.Close())
	}()

	content, err := io.ReadAll(valuteCurs)
	if err != nil {
		return nil, errReadXML
	}

	contentWithDots := bytes.ReplaceAll(content, []byte(","), []byte("."))

	decoder := xml.NewDecoder(bytes.NewReader(contentWithDots))
	decoder.CharsetReader = charset.NewReaderLabel

	var curs ValCurs
	if err := decoder.Decode(&curs); err != nil {
		return nil, errDecdXML
	}

	return &curs, nil
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
