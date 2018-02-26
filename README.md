# Alfred一款命令搜索workflow

### 介绍
日常开发中要记住的一些长命令太多,打起来很费事,因此使用workflow帮助管理命令.

该workflow把命令分为key -> values形式,如下所示,key属于大分类,匹配到key后会显示其下全部value.
```json
  {
    "key":"git-common",
    "values":[
      {
        "cmd":"git branch -r | sed 's/origin\///g' | grep '/' | xargs git push origin --delete",
        "remark":"git批量删除远程分支"
      }
    ],
    "remark":"git通用命令"
  }
```


主要功能:
1. cmd(触发关键词)->搜索key->选择value->复制到粘贴板(可以自动粘贴)
2. cmd(触发关键词)->open->选择打开命令配置->调用你喜欢的编辑器打开命令配置的json(主要是alfred添加数据不太好用)


### 数据保存
该插件对应json命令数据都是外置的(便于云端保存,丢到同步盘中即可),因此自己指定一个路径后,以参数形式传入即可.

也由于json外置,如果你觉得workflow添加太麻烦的话,直接改这个json,添加对应的数据即可.

![](http://oobu4m7ko.bkt.clouddn.com/1519546315.png)



### 演示

<p>
<img src="http://oobu4m7ko.bkt.clouddn.com/201811111111.gif" alt="" data-canonical-src="http://oobu4m7ko.bkt.clouddn.com/201811111111.gif" style="max-width:100%;">
</p>

### 其他问题

**1.不想要自动粘贴**

在alfred插件中选择粘贴这个环节

![](http://oobu4m7ko.bkt.clouddn.com/1519546487.png)

然后勾掉自动粘贴即可.

![](http://oobu4m7ko.bkt.clouddn.com/1519546513.png)

### 更新记录

#### 2018.02.26

关键词匹配去除`-_空白`特殊字符
