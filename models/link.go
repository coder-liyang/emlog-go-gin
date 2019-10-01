package models

type Link struct {
	Id          int64  `json:"id"`
	Sitename    string `json:"sitename"`
	Siteurl     string `json:"siteurl"`
	Description string `json:"description"`
	Hide        string `json:"hide"`
	Taxis       string `json:"taxis"`
}

func GetLinkList() []Link {
	//SELECT siteurl,sitename,description FROM " . DB_PREFIX . "link WHERE hide='n' ORDER BY taxis ASC
	var links []Link
	Orm.Where("hide = ?", "n").Order("taxis asc").Find(&links)
	return links
}