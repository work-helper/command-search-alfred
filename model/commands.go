package model

import (
	_ "encoding/json"
	"encoding/xml"
)

type Project struct {
	Key string

	Values []Value

	Remark string
}

type Value struct {
	Cmd string

	Remark string
}

// 结果

type Items struct {
	XMLName xml.Name `xml:"items"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	Autocomplete string `xml:"autocomplete,attr"`
	Uid          string `xml:"uid,attr"`
	Arg          string `xml:"arg,attr"`
	Title        string `xml:"title"`
	Subtitle     string `xml:"subtitle"`
	Icon         string `xml:"icon"`
}
