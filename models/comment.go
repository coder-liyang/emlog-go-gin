package models

type Comment struct {
	Cid           int64     `json:"cid"`
	Gid           int64     `json:"gid"`
	Pid           int64     `json:"pid"`
	Date          int64     `json:"date"`
	Poster        string    `json:"poster"`
	Comment       string    `json:"comment"`
	Mail          string    `json:"mail"`
	Url           string    `json:"url"`
	Ip            string    `json:"ip"`
	Hide          string    `json:"hide"`
	ChildComments []Comment `json:"child_comments"`
}

func GetCommentList(gid int64, pid int64, page int64) []Comment {
	var limit int64 = 10
	var commentList []Comment
	Orm.
		Where("gid = ? and pid = ? and hide = ?", gid, pid, "n").
		Limit(limit).
		Offset((page - 1) * limit).
		Find(&commentList)

	return commentList
}
