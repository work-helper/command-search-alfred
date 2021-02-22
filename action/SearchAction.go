package action

import (
	"encoding/json"
	"fmt"
	"github.com/work-helper/command-search-alfred/model"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("[-_\\s]")
var icon = model.Path{Path: "icon.png"}

func SearchKeysAndPrint(searchKey string, projects []model.Project) {
	// 为空则默认返回
	if "" == searchKey {
		parseCommands(projects)
		return
	}

	// 查找
	searchKey = re.ReplaceAllString(searchKey, "")
	var matchCommands []model.Project
	var noMatchCommands []model.Project
	for _, v := range projects {
		//以key开头
		key := re.ReplaceAllString(v.Key, "")
		tag := re.ReplaceAllString(v.Tags, "")
		if strings.Contains(key, searchKey) || strings.Contains(tag, searchKey) {
			matchCommands = append(matchCommands, v)
		} else {
			noMatchCommands = append(noMatchCommands, v)
		}

		// 完全匹配则直接处理
		if strings.EqualFold(key, searchKey) {
			parseCommandValue(&v)
			return
		}
	}

	//部分匹配的处理,匹配上的优先排序
	if len(matchCommands) == 0 {
		matchCommands = projects
		parseCommands(matchCommands)
	} else {
		matchCommands = append(matchCommands, noMatchCommands...)
		parseCommands(matchCommands)
	}
}

// 解析外围指令
func parseCommands(commands []model.Project) {
	var items = make([]model.Item, 0, len(commands))
	for _, v := range commands {
		var temp = ""
		if "" != v.Tags {
			temp = "\t\t--" + v.Tags
		}

		var item = model.Item{
			Autocomplete: v.Key,
			Arg:          "",
			Title:        v.Key,
			Subtitle:     v.Remark + temp,
			Icon:         icon,
		}
		items = append(items, item)
	}

	PrintResult(items)
}

// 解析匹配后内部指令
func parseCommandValue(command *model.Project) {
	var items = make([]model.Item, 0, len(command.Values))
	for _, v := range command.Values {
		var item = model.Item{
			Autocomplete: command.Key,
			Arg:          v.Cmd,
			Title:        v.Remark,
			Subtitle:     v.Cmd,
			Icon:         icon,
		}
		items = append(items, item)
	}
	PrintResult(items)
}

func PrintResult(items []model.Item) {
	result := model.Items{Items: items}
	bytes, _ := json.Marshal(result)
	fmt.Println(string(bytes))
}
