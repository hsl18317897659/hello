package models

type Upgrade struct {
	Id int
	Uid int
	Action int
	Version int
	Pasttercount int
	Novice int
	State int
	Ctime int
	Utime int
}

func (m*Upgrade) TableName() string  {
	return TableName("uc_towm_upgrade_user")
}
