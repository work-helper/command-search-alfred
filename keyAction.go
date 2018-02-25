package main

import (
	"strings"
)

func searchKeys(searchKey string,allCommands []Command) {
	// 查找
	var matchCommands []Command
	for _, v := range allCommands {

		//以key开头
		if strings.HasPrefix(v.Key, searchKey) {
			matchCommands = append(matchCommands, v)
		}

		if strings.EqualFold(v.Key, searchKey) {
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
