package action

import (
	"encoding/xml"
	"fmt"
	"github.com/work-helper/command-search-alfred/model"
	"strconv"
)

var openItem = model.Item{
	Autocomplete: "open",
	Uid:          "9999999",
	Arg:          "open",
	Title:        "打开命令配置",
	Subtitle:     "该命令会打开当前文件",
	Icon:         "./icon.png",
}

func ParseCommands(commands []model.Project, isAppendOpen bool) {
	var items = make([]model.Item, 0, len(commands))
	for _, v := range commands {
		var item = model.Item{
			Autocomplete: v.Key,
			Arg:          "",
			Title:        v.Key,
			Subtitle:     v.Remark,
			Icon:         "./icon.png",
		}
		items = append(items, item)
	}

	if isAppendOpen {
		items = append(items, openItem)
	}

	var result = model.Items{Items: items}
	marshal, _ := xml.Marshal(result)
	fmt.Println(string(marshal))
}

func parseCommandValue(command *model.Project) {
	var items = make([]model.Item, 0, len(command.Values))
	for index, v := range command.Values {
		var item = model.Item{
			Autocomplete: command.Key,
			Uid:          strconv.Itoa(index),
			Arg:          v.Cmd,
			Title:        v.Remark,
			Subtitle:     v.Cmd,
			Icon:         "./icon.png",
		}
		items = append(items, item)
	}
	var result = model.Items{Items: items}

	marshal, _ := xml.Marshal(result)
	fmt.Println(string(marshal))
}
