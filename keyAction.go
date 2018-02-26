package main

import (
	"strings"
	"regexp"
)

var re = regexp.MustCompile("[-_\\s]")

func searchKeys(searchKey string,allCommands []Command) {
	searchKey = re.ReplaceAllString(searchKey,"")
	// 查找
	var matchCommands []Command
	for _, v := range allCommands {

		//以key开头
		key := re.ReplaceAllString(v.Key, "")
		if strings.HasPrefix(key, searchKey) {
			matchCommands = append(matchCommands, v)
		}

		if strings.EqualFold(key, searchKey) {
			// 找到了则直接处理
			parseCommandValue(&v)
			return
		}
	}
	//没找到完全匹配的则对部分匹配的处理
	if len(matchCommands) == 0 {
		matchCommands = allCommands
	}
	parseCommands(matchCommands)
}
