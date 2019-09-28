package models

type Sort struct {
	Sid         int64 `json:"sid",gorm:"primary_key;auto_increment"`
	Sortname    string `json:"sortname"`
	Alias       string `json:"alias"`
	Taxis       int64 `json:"taxis"`
	Pid         int64 `json:"pid"`
	Description string `json:"description"`
	Template    string `json:"template"`
	ChildSort	[]Sort `json:"child_sort"`
}


//分类列表
func GetSortList(pid int64) []Sort {
	var sortList []Sort
	qs := Orm.Where("pid = ?", pid)
	qs.Find(&sortList)
	//查询有子分类
	for index, item := range sortList {
		childSortList := GetSortList(item.Sid)
		if len(childSortList) > 0 {
			sortList[index].ChildSort = childSortList
		}
	}
	return sortList
}
