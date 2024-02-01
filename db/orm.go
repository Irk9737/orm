package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"orm/business"
	"orm/common"
	"xorm.io/xorm"
)

type ValueOrm struct {
	Xorm *xorm.Engine
}

type Repository interface {
	InsertData(c *gin.Context)
}

var _ Repository = &ValueOrm{}

func (db *ValueOrm) GetEngine() (*xorm.Engine, error) {
	sqlStr := "root:010729@tcp(127.0.0.1:3306)/xorm?charset=utf8mb4&parseTime=true&loc=Local" // xorm: 数据库名称
	var err error
	// 1、创建数据库引擎
	db.Xorm, err = xorm.NewEngine("mysql", sqlStr)
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return nil, err
	}
	// 2、创建或同步表 Stu
	err = db.Xorm.Sync(new(business.Stu))
	if err != nil {
		fmt.Println("数据库同步失败:", err)
		return nil, err
	}

	return db.Xorm, nil
}

// InsertData 插入操作
func InsertData(c *gin.Context) {
	var s business.Stu
	err := c.Bind(&s)
	if err != nil {
		common.Response.Code = http.StatusBadRequest
		xormResponse.Message = "参数错误"
		xormResponse.Data = "error"
		c.JSON(http.StatusOK, xormResponse)
		return
	}
	// affected：受影响记录行数
	affected, err := x.Insert(s)
	if err != nil || affected <= 0 {
		fmt.Printf("insert failed, err:%v\n", err)
		xormResponse.Code = http.StatusBadRequest
		xormResponse.Message = "写入失败"
		xormResponse.Data = err
		c.JSON(http.StatusOK, xormResponse)
		return
	}
	xormResponse.Code = http.StatusOK
	xormResponse.Message = "写入成功"
	xormResponse.Data = "OK"
	c.JSON(http.StatusOK, xormResponse)
	fmt.Println(affected) // 打印结果
}

// xormUpdateData 删除操作
//func xormDeleteData(c *gin.Context) {
//	stuNum := c.Query("stu_num")
//	// 1、先查询
//	var stus []Stu
//	err := x.Where("stu_num=?", stuNum).Find(&stus)
//	if err != nil || len(stus) <= 0 {
//		xormResponse.Code = http.StatusBadRequest
//		xormResponse.Message = "数据不存在"
//		xormResponse.Data = "error"
//		c.JSON(http.StatusOK, xormResponse)
//		return
//	}
//	// 2、再删除
//	affected, err := x.Where("stu_num=?", stuNum).Delete(&Stu{})
//	if err != nil || affected <= 0 {
//		xormResponse.Code = http.StatusBadRequest
//		xormResponse.Message = "删除失败"
//		xormResponse.Data = "error"
//		c.JSON(http.StatusOK, xormResponse)
//		return
//	}
//	xormResponse.Code = http.StatusOK
//	xormResponse.Message = "删除成功"
//	xormResponse.Data = "OK"
//	c.JSON(http.StatusOK, xormResponse)
//	fmt.Println(affected) // 打印结果
//}

// xormUpdateData 修改操作
//func xormUpdateData(c *gin.Context) {
//	var s Stu
//	err := c.Bind(&s)
//	if err != nil {
//		xormResponse.Code = http.StatusBadRequest
//		xormResponse.Message = "参数错误"
//		xormResponse.Data = "error"
//		c.JSON(http.StatusOK, xormResponse)
//		return
//	}
//	// 1、先查询
//	var stus []Stu
//	err = x.Where("stu_num=?", s.StuNum).Find(&stus)
//	if err != nil || len(stus) <= 0 {
//		xormResponse.Code = http.StatusBadRequest
//		xormResponse.Message = "数据不存在"
//		xormResponse.Data = "error"
//		c.JSON(http.StatusOK, xormResponse)
//		return
//	}
//	// 2、再修改
//	affected, err := x.Where("stu_num=?", s.StuNum).Update(&Stu{Name: s.Name, Age: s.Age})
//	if err != nil || affected <= 0 {
//		xormResponse.Code = http.StatusBadRequest
//		xormResponse.Message = "修改失败"
//		xormResponse.Data = "error"
//		c.JSON(http.StatusOK, xormResponse)
//		return
//	}
//	xormResponse.Code = http.StatusOK
//	xormResponse.Message = "修改成功"
//	xormResponse.Data = "OK"
//	c.JSON(http.StatusOK, xormResponse)
//	fmt.Println(affected) // 打印结果
//}

// xormGetMulData 查询操作(多条记录)
//func xormGetMulData(c *gin.Context) {
//	name := c.Query("name")
//	var stus []Stu
//	err := x.Where("name=?", name).And("age>20").Limit(10, 0).Asc("age").Find(&stus)
//	if err != nil {
//		xormResponse.Code = http.StatusBadRequest
//		xormResponse.Message = "查询错误"
//		xormResponse.Data = "error"
//		c.JSON(http.StatusOK, xormResponse)
//		return
//	}
//	xormResponse.Code = http.StatusOK
//	xormResponse.Message = "查询成功"
//	xormResponse.Data = stus
//	c.JSON(http.StatusOK, xormResponse)
//}

// xormGetData 查询操作(单条记录)
//func xormGetData(c *gin.Context) {
//	stuNum := c.Query("stu_num")
//	var stus []Stu
//	err := x.Where("stu_num=?", stuNum).Find(&stus)
//	if err != nil {
//		xormResponse.Code = http.StatusBadRequest
//		xormResponse.Message = "查询错误"
//		xormResponse.Data = "error"
//		c.JSON(http.StatusOK, xormResponse)
//		return
//	}
//	xormResponse.Code = http.StatusOK
//	xormResponse.Message = "查询成功"
//	xormResponse.Data = stus
//	c.JSON(http.StatusOK, xormResponse)
//}
