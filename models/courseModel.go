package models

type Course struct{
	Id int
	Uid int
	Packid int
	Courseid int
	Jindu int
	Vtime int
	Ustate int
	State int
	Viewtype int
	Isread int
	Stime string
	Utime string
	Ctime string
}

func (m *Course) TableName() string  {
	return TableName("qk_course_user")
}