package models

type Attire struct {
	Id      int
	Type    int
	Subtype int
	Name    string
	Onpic   string
	State   int
	Utime   string
	Ctime   int
}

func (m*Attire) TableName()  string{
	return TableName("uc_town_attire")
}
