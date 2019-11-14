package models

type Propuid struct {
	Id int
	Uid int
	Propid int
	Tid int
	Pointx float32
	Pointy float32
	State int
	Sort int
	Utime int
	Ctime int
}

func (m*Propuid) TableName() string  {
	return TableName("uc_town_prop_uid")
}
