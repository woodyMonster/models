package models

import (
	"time"
	"github.com/woodyMonster/DBLoader/dbloader"
)

var _db = dbloader.Init("WebService")

// User init ORM
type User struct {
	ID        int       `json:"id" form:"id" gorm:"id;AUTO_INCREMENT"`
	Name      string    `json:"name" form:"name" gorm:"name"`
	Account   string    `json:"account" form:"account" gorm:"account"`
	Password  string    `json:"password" form:"password" gorm:"password"`
	UpdatedAt time.Time `gorm:"default:current_time on update current_time"`
}

// Create Create
func (u *User) Create() (id int, err error) {
	err = _db.Create(u).Error
	return u.ID, err
}

// Get 取得 user 資料
func (u *User) Get(uid string) (User, bool) {
	var user User
	err := _db.Where("id = ?", uid).First(&user).RecordNotFound()
	return user, err
}

// Update 更新資料
func (u *User) Update(uid string) error {
	err := _db.Model(User{}).Where("id = ?", uid).Updates(&u).Error
	return err
}

// Delete 刪除資料
func (u *User) Delete(uid string) error {
	err := _db.Model(User{}).Where("id = ?", uid).Delete(&u).Error
	return err
}
