package Dao

import (
	"ginFreamwork/common"
)

type TodoList struct {
	Id        int              `json:"id"`
	Title     string           `json:"title"`
	EndTime   string           `json:"end_time"`
	Status    int              `json:"status"`
	Content   string           `json:"content"`
	Uid       int              `json:"uid"`
	Day       string           `json:"-"`
	CreatedAt common.JSONTime  `json:"created_at" `
	UpdatedAt common.JSONTime  `json:"-" `
	DeletedAt *common.JSONTime `json:"-" `
}

func FindAll(page int, pageSize int, where interface{}) []TodoList {
	todoList := []TodoList{}
	//if where != nil {
	//	common.DBHelper = common.DBHelper.Where(where)
	//}
	offset := (page - 1) * pageSize
	common.DBHelper.Where(where).Offset(offset).Limit(pageSize).Find(&todoList)
	return todoList
}

func GetInfoById(id int) TodoList {
	todo := TodoList{}
	common.DBHelper.First(&todo, id)
	return todo
}

func Add(data TodoList) int {
	common.DBHelper.Create(&data)
	return data.Id
}

func Del(data TodoList) int64 {
	return common.DBHelper.Delete(&data).RowsAffected

}

func SetStatus(id int, status int) int64 {
	return common.DBHelper.Model(&TodoList{}).Where("id=?", id).Update("status", status).RowsAffected
}
