package models

type Prop struct {
	Id int
	Tid int
	Scene string
	Name string
	Icon string
	Onpic string
	Quality int
	Prosperity int
	Woods int
	Minerals int
	Bronzes int
	Recoveryprice int
	State int
	Issend int
	Utime int
	Ctime int
}

func (m*Prop) TableName() string {
	return TableName("uc_town_prop")
}
