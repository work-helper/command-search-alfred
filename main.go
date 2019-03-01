package main

import (
	"flag"
	"github.com/work-helper/command-search-alfred/action"
	"github.com/work-helper/command-search-alfred/util"
	"strings"
)

const defaultDataPath = "./data.json"

func main() {
	// search key
	searchKey := flag.String("s", "", "search key")
	dataPath := flag.String("p", defaultDataPath, "data source")
	flag.Parse()

	projects := util.ReadProjects(*dataPath)

	if len(projects) == 0 || 0 == strings.Compare("open", *searchKey) {
		projects = projects[:0]
		action.ParseCommands(projects, true)
		return
	}

	// 搜索key不为空时则搜索
	if len(*searchKey) != 0 {
		action.SearchKeys(*searchKey, projects)
	}
}
