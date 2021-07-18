package main

import (
	"encoding/xml"
	"fmt"
	"github.com/goccy/go-json"
)

func FormatJSON(j []byte, prefix, indent string) ([]byte, error) {
	type jsonData map[string]interface{}
	var jd jsonData

	if err := json.Unmarshal(j, &jd); err != nil {
		return nil, err
	}
	// jsonデータを接頭辞、インデント付きで見やすく整形
	return json.MarshalIndent(jd, prefix, indent)
}

func FormatXML(x []byte, prefix, indent string) ([]byte, error) {
	//type xmlData map[string]interface{}
	var xd map[string]interface{}

	if err := xml.Unmarshal(x, &xd); err != nil {
		return nil, err
	}

	x, err := xml.MarshalIndent(xd, prefix, indent)
	if err != nil {
		return nil, err
	}

	// TODO: xmlの整形手段を作成する
	return nil, fmt.Errorf("SAMPLE ERROR\n")
}

