package main

import (
	xj "github.com/basgys/goxml2json"
	"strings"
)

func XML2JSON(xml, prefix, indent string) ([]byte, error) {
	xmlData := strings.NewReader(xml)
	j, err := xj.Convert(xmlData)
	if err != nil {
		return nil, err
	}

	b, err := FormatJSON(j.Bytes(), prefix, indent)
	if err != nil {
		return nil, err
	}

	return b, nil
}