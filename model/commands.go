package model

import (
	_ "encoding/json"
)

// 数据结果
type Project struct {
	Key string

	// 参与搜索
	Tags string

	Values []Value

	Remark string
}

type Value struct {
	Cmd string

	Remark string
}

// alfred 模型
type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	// 可选,每个 item 的唯一标识，后续 Alfred 能依托这个 uid，根据用户操作进行排序。如果想保持自己脚本返回的顺序，不用 Alfred 的排序，可以不设置这个字段。
	Uid string `json:"uid,omitempty"`
	//  “default” | “file” | “file:skipcheck”（可选，默认 “default”）
	Type string `json:"type"`
	// 顾名思义，Row 的标题。
	Title string `json:"title"`
	// subtitle（可选）
	Subtitle string `json:"subtitle"`
	// 非常建议 item 包含这个字段，它会作为 item 的输出，传递给下一个事件。如果不包含这个字段，就无法知道用户选的是哪个。
	Arg          string `json:"arg"`
	Autocomplete string `json:"autocomplete,omitempty"`
	// 以相对路径形式引用 workflow 沙盒中的图标
	Icon Path `json:"icon"`
}

type Path struct {
	Path string `json:"path"`
}
