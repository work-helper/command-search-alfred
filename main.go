package main

import (
	"flag"
	"os"
	"encoding/json"
	"io/ioutil"
)

const defaultDataPath = "./data.json"

func main() {
	// search key
	searchKey := flag.String("s", "", "search key")
	dataPath := flag.String("p", defaultDataPath, "data source")
	flag.Parse()

	// 读取json文件
	if !PathExist(*dataPath) {
		_, e := os.Create(*dataPath)
		if e != nil {
			panic(e)
		}
	}
	bytes, e := ioutil.ReadFile(*dataPath)
	checkError(e)
	//转换命令
	var allCommands []Command
	json.Unmarshal(bytes, &allCommands)
	if len(allCommands) == 0 {
		parseCommands(allCommands)
		return
	}

	// 搜索key不为空时则搜索
	if len(*searchKey) != 0 {
		searchKeys(*searchKey, allCommands)
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 检查文件夹是否存在
 */
func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
