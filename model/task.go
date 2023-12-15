package model

import (
	"strconv"
	"work/cache"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	/*

			ID        uint `gorm:"primary_key"`
		    CreatedAt time.Time
		    UpdatedAt time.Time
		    DeletedAt *time.Time `sql:"index"`

	*/
	User      User   `gorm:"ForeignKey:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"default:'0'"` //0未完成
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}

func (t *Task) View() int {
	// 浏览次数
	_ = cache.Rd.SetNX(strconv.Itoa(int(t.ID)), 1, 0)
	countStr, _ := cache.Rd.Get(strconv.Itoa(int(t.ID))).Result()
	count, _ := strconv.Atoi(countStr)
	return count
}

func (t *Task) AddView() {
	// 增加浏览次数
	cache.Rd.Incr(strconv.Itoa(int(t.ID)))

}

/*
Incr()
如果键key不存在，会创建一个新的键，并将其初始值设置为0。
如果键key存在且其值可以被解析为整数，那么将其值增加1，并返回增加后的值。
如果键key存在但其值无法被解析为整数，那么该函数将返回一个错误。
*/
