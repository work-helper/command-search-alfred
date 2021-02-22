package main

import (
	"flag"
	"github.com/work-helper/command-search-alfred/action"
	"github.com/work-helper/command-search-alfred/model"
	"github.com/work-helper/command-search-alfred/util"
	"os"
	"strings"
)

const defaultDataPath = "./dataTpl.yaml"

var errorIcon = model.Path{Path: "error.png"}
var openIcon = model.Path{Path: "open.png"}

func main() {
	// search key
	searchKey := flag.String("s", "", "search key")
	dataPath := flag.String("p", defaultDataPath, "data source")
	flag.Parse()

	// 检查文件是否存在
	_, err := os.Stat(*dataPath)
	if nil != err && !os.IsExist(err) {
		item := model.Item{
			Type:     "default",
			Title:    "配置文件不存在",
			Subtitle: "请检查数据路径配置是否正确",
			Icon:     errorIcon,
		}
		action.PrintResult([]model.Item{item})
		return
	}

	// 判断是否为打开文件
	if 0 == strings.Compare("open", *searchKey) {
		item := model.Item{
			Type:     "default",
			Arg:      "open",
			Title:    "打开命令配置",
			Subtitle: "该命令会打开当前文件",
			Icon:     openIcon,
		}
		action.PrintResult([]model.Item{item})
		return
	}

	// 进行数据搜索
	projects := util.ReadProjects(*dataPath)
	action.SearchKeysAndPrint(*searchKey, projects)
}
