package business

import "time"

// Stu 定义结构体(xorm 支持双向映射)：没有表会进行创建
type Stu struct {
	Id      int64     `xorm:"pk autoincr" json:"id"`
	StuNum  string    `xorm:"unique" json:"stu_num"`
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}
