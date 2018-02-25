package main

import (
	"strconv"
	"fmt"
	"encoding/xml"
)

var openItem = Item{
	Autocomplete: "open",
	Uid:          "99999",
	Arg:          "open",
	Title:        "打开命令配置",
	Subtitle:     "该命令会打开当前文件",
	Icon: "./icon.png",
}

func parseCommands(commands []Command) {
	var items = make([]Item,0,len(commands))
	for index, v := range commands {
		var item = Item{
			Autocomplete: v.Key,
			Uid:          strconv.Itoa(index),
			Arg:          "",
			Title:        v.Key,
			Subtitle:     v.Remark,
			Icon: "./icon.png",
		}
		items = append(items, item)
	}


	var result = Items{Items: append(items,openItem)}
	marshal, _ := xml.Marshal(result)
	fmt.Println(string(marshal))
}

func parseCommandValue(command *Command) {
	var items = make([]Item,0,len(command.Values))
	for index, v := range command.Values {
		var item = Item{
			Autocomplete: command.Key,
			Uid:          strconv.Itoa(index),
			Arg:          v.Cmd,
			Title:        v.Remark,
			Subtitle:     v.Cmd,
			Icon: "./icon.png",
		}
		items = append(items, item)
	}
	var result = Items{Items: items}

	marshal, _ := xml.Marshal(result)
	fmt.Println(string(marshal))
}

