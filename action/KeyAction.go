package action

import (
	"github.com/work-helper/command-search-alfred/model"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("[-_\\s]")

func SearchKeys(searchKey string, projects []model.Project) {
	searchKey = re.ReplaceAllString(searchKey, "")
	// 查找
	var matchCommands []model.Project
	var noMatchCommands []model.Project
	for _, v := range projects {

		//以key开头
		key := re.ReplaceAllString(v.Key, "")
		if strings.HasPrefix(key, searchKey) {
			matchCommands = append(matchCommands, v)
		} else {
			noMatchCommands = append(noMatchCommands, v)
		}

		if strings.EqualFold(key, searchKey) {
			// 找到了则直接处理
			parseCommandValue(&v)
			return
		}
	}
	//没找到完全匹配的则对部分匹配的处理
	if len(matchCommands) == 0 {
		matchCommands = projects
		ParseCommands(matchCommands, true)
	} else {
		matchCommands = append(matchCommands, noMatchCommands...)
		ParseCommands(matchCommands, false)
	}
}
