package models

type Navi struct {
	Id int64
	Naviname string
	Url string
	Newtab string
	Hide string
	Taxis int64
	Pid int64
	Isdefault string
	Type int64
	Type_id int64
	ChildNavi []Navi
}


//导航列表
func GetNaviList(pid int64) []Navi {
	var naviList []Navi
	qs := Orm.Where("pid = ? and hide = 'n'", pid).Order("taxis asc")
	qs.Find(&naviList)
	//查询有子导航
	for index, item := range naviList {
		//TODO 如果item.type=4,则此导航下是分类,需要查询分类信息
		childNaviList := GetNaviList(item.Id)
		if len(childNaviList) > 0 {
			naviList[index].ChildNavi = childNaviList
		}
	}
	return naviList
}
