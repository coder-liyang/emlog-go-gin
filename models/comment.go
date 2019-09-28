package models

type Comment struct {
	Cid     int64
	Gid     int64
	Pid     int64
	Date    int64
	Poster  string
	Comment string
	Mail    string
	Url     string
	Ip      string
	Hide    string
	ChildComments []Comment
}

func GetCommentList(gid int64, pid int64, page int64) []Comment {
	var limit int64 = 10
	var commentList  []Comment
	Orm.
		Where("gid = ? and pid = ? and hide = ?", gid, pid, "n").
		Limit(limit).
		Offset((page-1) * limit).
		Find(&commentList)

	return commentList
}