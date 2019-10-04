package models

type Navi struct {
	Id        int64  `json:"id"`
	Naviname  string `json:"naviname"`
	Url       string `json:"url"`
	Newtab    string `json:"newtab"`
	Hide      string `json:"hide"`
	Taxis     int64  `json:"taxis"`
	Pid       int64  `json:"pid"`
	Isdefault string `json:"isdefault"`
	Type      int64  `json:"type"` //1,2,3:系统;4:分类;5:页面;0:自定义
	Type_id   int64  `json:"type_id"`
	ChildNavi []Navi `json:"child_navi"`
}

//导航列表
func GetNaviList(pid int64) []Navi {
	var naviList []Navi
	qs := Orm.Where("pid = ? and hide = 'n'", pid).Order("taxis asc")
	qs.Find(&naviList)
	//查询有子导航
	for index, item := range naviList {
		childNaviList := GetNaviList(item.Id)
		if len(childNaviList) > 0 {
			naviList[index].ChildNavi = childNaviList
		} else if item.Type == 4 {
			//如果item.type=4,则此导航下是分类,需要查询分类信息
			sort := GetSortList(item.Type_id)
			for _, sortItem := range sort {
				naviList[index].ChildNavi = append(naviList[index].ChildNavi, Navi{
					Id:        0,
					Naviname:  sortItem.Sortname,
					Url:       "",
					Newtab:    "n",
					Hide:      "n",
					Taxis:     0,
					Pid:       item.Id,
					Isdefault: "",
					Type:      4,
					Type_id:   0,
					ChildNavi: nil,
				})
			}
		}
	}
	return naviList
}
