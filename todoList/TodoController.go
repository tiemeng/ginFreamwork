package todoList

import (
	"ginFreamwork/Dao"
	"ginFreamwork/common"
	"github.com/gin-gonic/gin"
	strconv2 "strconv"
	"sync"
	"time"
)

//使用加锁方式使用map，解决并发时出现的不安全性
var ReData = struct{
	sync.RWMutex
	m map[string]interface{}
}{m: make(map[string]interface{})}

var Where =struct{
	sync.RWMutex
	w map[string]interface{}
}{w: make(map[string]interface{})}


type FormData struct {
	Title   string `form:"title" binding:"required,min=4,max=200"`
	Date    string `form:"date" binding:"required"`
	Time    string `form:"time" binding:"required"`
	Content string `form:"content" binding:"required,min=10"`
}

func TodoList(c *gin.Context) {
	page, _ := strconv2.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv2.Atoi(c.DefaultQuery("pageSize", "10"))
	status, _ := strconv2.Atoi(c.DefaultQuery("status", "1"))
	tType, _ := strconv2.Atoi(c.DefaultQuery("type", "0"))
	where := make(map[string]interface{})
	where["status"] = status
	where["day"] = time.Now().Format("2006-01-02")
	data := make(map[string]interface{})
	data["今日事项"] = Dao.FindAll(page, pageSize, where)
	if tType == 0{
		nTime := time.Now()
		yesTime := nTime.AddDate(0, 0, -1)
		where["day"] = yesTime.Format("2006-01-02")
		data["昨日事项"] = Dao.FindAll(page, pageSize, where)
	}
	ReData.Lock()
	ReData.m["code"] = 200
	ReData.m["message"] = "请求成功"
	ReData.m["data"] = data
	common.ReJson(c, ReData.m)
	ReData.Unlock()
}

func GetDetail(c *gin.Context) {
	id, _ := strconv2.Atoi(c.Query("id"))
	data := Dao.GetInfoById(id)

	ReData.Lock()
	if data.Id <= 0 {
		ReData.m["code"] = 201
		ReData.m["message"] = "记录不存在"
		ReData.m["data"] = nil
	} else {
		ReData.m["code"] = 200
		ReData.m["message"] = "success"
		ReData.m["data"] = data
	}
	common.ReJson(c, ReData.m)
	ReData.Unlock()
}

func Index(c *gin.Context) {
	page, _ := strconv2.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv2.Atoi(c.DefaultQuery("pageSize", "10"))
	where := make(map[string]interface{})
	ReData.Lock()
	ReData.m["code"] = 200
	ReData.m["message"] = "success"
	ReData.m["data"] = Dao.FindAll(page, pageSize, where)
	common.ReJson(c, ReData.m)
	ReData.Unlock()
}

func Add(c *gin.Context) {
	data := FormData{}
	err := c.ShouldBind(&data)
	addData := Dao.TodoList{
		Title:   data.Title,
		Content: data.Content,
		EndTime: data.Date + " " + data.Time,
		Status:  1,
		Day:     time.Now().Format("2006-01-02"),
	}
	ReData.Lock()
	if err != nil {
		ReData.m["code"] = 201
		ReData.m["message"] = err.Error()
		ReData.m["data"] = nil
	} else {
		id := Dao.Add(addData)
		if id > 0 {
			ReData.m["code"] = 200
			ReData.m["message"] = "添加成功"
			ReData.m["data"] = id
		} else {
			ReData.m["code"] = 201
			ReData.m["message"] = "添加失败"
			ReData.m["data"] = nil
		}

	}
	common.ReJson(c, ReData.m)
	ReData.Unlock()
}

func Del(c *gin.Context) {
	id, _ := strconv2.Atoi(c.Query("id"))
	todo := Dao.GetInfoById(id)
	ReData.Lock()
	if todo.Id <= 0 {
		ReData.m["code"] = 201
		ReData.m["message"] = "记录不存在"
		ReData.m["data"] = nil
	} else {
		affectRow := Dao.Del(todo)
		if affectRow > 0 {
			ReData.m["code"] = 200
			ReData.m["message"] = "删除成功"
			ReData.m["data"] = nil
		} else {
			ReData.m["code"] = 201
			ReData.m["message"] = "删除失败"
			ReData.m["data"] = nil
		}
	}
	common.ReJson(c, ReData.m)
	ReData.Unlock()
}

func Finish(c *gin.Context) {
	id, _ := strconv2.Atoi(c.Query("id"))
	status, _ := strconv2.Atoi(c.Query("status"))
	todo := Dao.GetInfoById(id)
	ReData.Lock()
	if todo.Id <= 0 {
		ReData.m["code"] = 201
		ReData.m["message"] = "记录不存在"
		ReData.m["data"] = nil
	} else {
		affectRow := Dao.SetStatus(id, status)
		if affectRow > 0 {
			ReData.m["code"] = 200
			ReData.m["message"] = "设置成功"
			ReData.m["data"] = nil
		} else {
			ReData.m["code"] = 201
			ReData.m["message"] = "更新失败"
			ReData.m["data"] = nil
		}
	}
	common.ReJson(c, ReData.m)
	ReData.Unlock()
}
