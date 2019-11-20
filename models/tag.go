package models

type Tag struct {
	Tid int64 `json:"tid"`
	Tagname string `json:"tagname"`
	Gid string `json:"-"`
}

func GetTagList() []Tag  {
	var tagList []Tag
	Orm.Find(&tagList)
	return tagList
}