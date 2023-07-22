package main

import (
	"gorm/database"
	"time"
)

// 结构体的蛇形复数作为表名（UserInfo --> user_infos）
// 默认列名是字段名的蛇形小写（CreatedAt --> created_at）
type UserInfo struct {
	ID        uint // 默认为主键
	Name      string
	Age       int
	CreatedAt time.Time // 创建记录时，如果该字段值为零值，则将该字段的值设为当前时间
	UpdatedAt time.Time // 更新记录时，将该字段的值设为当前时间，创建记录时，如果该字段值为零值，则将该字段的值设为当前时间
}

// 将 UserInfo 的表名设置为 `user`
func (UserInfo) TableName() string {
	return "user"
}

// type User struct {
// 	ID       uint   `gorm:"column:user_id"`  // 将列名设为 `user_id`
// 	UserName string `gorm:"column:name"`     // 将列名设为 `name`
// 	Age      int64  `gorm:"column:user_age"` // 将列名设为 `user_age`
// }

// type User struct {
// 	ID       uint   `gorm:"autoIncrement"`                    // 自增主键
// 	Name     string `gorm:"size:10"`                          // 大小为 10 个字符
// 	Age      int    `gorm:"size:3; check:age > 0"`            // 大小为 3 位数，并且值要 > 0
// 	Email    string `gorm:"type:varchar(25); unique;"`        // 类型为 varchar(25)，唯一
// 	Password string `gorm:"type:varchar(20); default:123456"` // 类型为 varchar(20)，默认值为 123456
// }

// type User struct {
// 	gorm.Model
// 	Name string
// 	Age  int
// }

func main() {
	database.Init()

	database.DB.AutoMigrate(&UserInfo{})
}
