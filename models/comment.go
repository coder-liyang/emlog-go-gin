package models

//import "fmt"

type Comment struct {
	Cid           int64     `json:"cid"`
	Gid           int64     `json:"gid"`
	Pid           int64     `json:"pid"`
	Date          int64     `json:"date"`
	Poster        string    `json:"poster"`
	Comment       string    `json:"comment"`
	Mail          string    `json:"-"`
	Url           string    `json:"url"`
	Ip            string    `json:"-"`
	Hide          string    `json:"-"`
	ChildComments []Comment `json:"child_comments"`
}

//评论列表,pageSize传入0,此时会查询所有评论
func GetCommentList(gid int64, pid int64, page int64, pageSize int64) []Comment {
	var commentList []Comment
	qs := Orm.Where("gid = ? and pid = ? and hide = ?", gid, pid, "n")
	qs = qs.Order("date desc")
	if pageSize > 0 {
		qs = qs.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	qs.Find(&commentList)
	//查询有无子评论
	for index, item := range commentList {
		childCommentList := GetCommentList(gid, item.Cid, 1, 0)
		if len(childCommentList) > 0 {
			commentList[index].ChildComments = childCommentList
		}
	}
	return commentList
}

//取一条评论
func GetComment(cid int64) (Comment, error) {
	var comment Comment
	err := Orm.Where("cid = ?", cid).First(&comment).Error
	return comment, err
}

func CreateComment(comment *Comment) (Comment, error){
	err := Orm.Create(&comment).Error
	return *comment, err
}