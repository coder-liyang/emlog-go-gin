package models

type Blog struct {
	Gid         int64 `json:"gid"`
	Title       string `json:"title"`
	Date        int64 `json:"date"`
	Content     string `json:"content"`
	Excerpt     string `json:"excerpt"`
	Alias       string `json:"alias"`
	Author      int64 `json:"author"`
	Sortid      int64 `json:"sortid"`
	Type        string `json:"type"`
	Views       int `json:"views"`
	Comnum      int `json:"comnun"`
	Attnum      int `json:"attnum"`
	Top         string `json:"top"`
	Sortop      string `json:"sortop"`
	Hide        string `json:"hide"`
	Checked     string `json:"checked"`
	AllowRemark string `json:"all_remark"`
	Password    string `json:"password"`
	Template    string `json:"template"`
}

//指定表名
//func (m *Blog) TableName() string {
//	return "e_blog"
//}

func GetBlogById(gid int64) Blog {
	var blog Blog
	Orm.Where(&Blog{Gid:gid}).First(&blog)
	return blog
}