package models

type Sort struct {
	Sid         int64 `json:"sid",gorm:"primary_key;auto_increment"`
	Sortname    string
	Alias       string
	Taxis       int64
	Pid         int64
	Description string
	Template    string
	ChildSort	[]*Sort `orm:"-"`
}
