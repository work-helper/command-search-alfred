package util

import (
	"command-search-alfred/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadProjects(dataPath string) []model.Project {
	// 读取json文件
	if !PathExist(dataPath) {
		panic(fmt.Sprintf("datapath %s not exist", dataPath))
	}
	bytes, e := ioutil.ReadFile(dataPath)
	if nil != e {
		panic(e)
	}
	//转换命令
	var projects []model.Project
	json.Unmarshal(bytes, &projects)
	return projects
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