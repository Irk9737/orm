package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
	"xorm.io/xorm"
)

// xorm 依赖: go get xorm.io/xorm
// 通过 xorm 进行数据库的 CRUD 操作

var x *xorm.Engine
var xormResponse XormResponse

type XormValue struct {
	Xorm *xorm.Engine
}

type OrmValue struct {
	xVal XormValue
}

// Stu 定义结构体(xorm 支持双向映射)：没有表会进行创建
type Stu struct {
	Id      int64     `xorm:"pk autoincr" json:"id"`
	StuNum  string    `xorm:"unique" json:"stu_num"`
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

// XormResponse 应答 Client 请求
type XormResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func init() {
	sqlStr := "root:pwd@tcp(127.0.0.1:3306)/xorm?charset=utf8mb4&parseTime=true&loc=Local" // xorm: 数据库名称
	var err error
	// 1、创建数据库引擎
	x, err = xorm.NewEngine("mysql", sqlStr)
	if err != nil {
		fmt.Println("数据库连接失败:", err)
	}
	// 2、创建或同步表 Stu
	err = x.Sync(new(Stu))
	if err != nil {
		fmt.Println("数据库同步失败:", err)
	}
}

func main() {
	r := gin.Default()
	// 数据库的 CRUD ---> gin 的 POST GET PUT DELETE 方法
	r.POST("xorm/insert", XormInsertData[Stu])
	r.GET("xorm/get", XormGetData[Stu])
	r.GET("xorm/mulget", XormGetMulData[Stu])
	r.PUT("xorm/update", XormUpdateData)
	r.DELETE("xorm/delete", XormDeleteData[Stu])
	r.Run(":8080")
}

// XormUpdateData 删除操作
func XormDeleteData[T any](c *gin.Context) {
	stuNum := c.Query("stu_num")
	// 1、先查询
	//var stus []Stu
	var typs []T
	err := x.Where("stu_num=?", stuNum).Find(&typs)
	if err != nil || len(typs) <= 0 {
		HandleResponse(c, http.StatusBadRequest, "数据不存在", "error")
		return
	}
	// 2、再删除
	affected, err := x.Where("stu_num=?", stuNum).Delete(&Stu{})
	if err != nil || affected <= 0 {
		HandleResponse(c, http.StatusBadRequest, "删除失败", "error")
		return
	}
	HandleResponse(c, http.StatusOK, "删除成功", "OK")
	fmt.Println(affected) // 打印结果
}

// XormUpdateData 修改操作
func XormUpdateData(c *gin.Context) {
	var s Stu
	err := c.Bind(&s)
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, "参数错误", "error")
		return
	}
	// 1、先查询
	var stus []Stu
	err = x.Where("stu_num=?", s.StuNum).Find(&stus)
	if err != nil || len(stus) <= 0 {
		HandleResponse(c, http.StatusBadRequest, "数据不存在", "error")
		return
	}
	// 2、再修改
	affected, err := x.Where("stu_num=?", s.StuNum).Update(&Stu{Name: s.Name, Age: s.Age})
	if err != nil || affected <= 0 {
		HandleResponse(c, http.StatusBadRequest, "修改失败", "error")
		return
	}
	HandleResponse(c, http.StatusOK, "修改成功", "OK")
	fmt.Println(affected) // 打印结果
}

// XormGetMulData 查询操作(多条记录)
func XormGetMulData[T any](c *gin.Context) {
	name := c.Query("name")
	var typs []T
	err := x.Where("name=?", name).And("age>20").Limit(10, 0).Asc("age").Find(&typs)
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, "查询错误", "error")
		return
	}
	HandleResponse(c, http.StatusOK, "查询成功", typs)
}

func QueryData[T any](c *gin.Context) ([]T, error) {
	stuNum := c.Query("stu_num")
	var types []T
	err := x.Where("stu_num=?", stuNum).Find(&types)
	return types, err
}

// XormGetData 查询操作(单条记录)
func XormGetData[T any](c *gin.Context) {
	typs, err := QueryData[T](c)
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, "查询错误", "error")
		return
	}
	HandleResponse(c, http.StatusOK, "查询成功", typs)
}

// XormInsertData 插入操作
func XormInsertData[T any](c *gin.Context) {
	var typ T
	err := c.Bind(&typ)
	if err != nil {
		HandleResponse(c, http.StatusBadRequest, "参数错误", "error")
		return
	}
	// affected：受影响记录行数
	affected, err := x.Insert(typ)
	if err != nil || affected <= 0 {
		fmt.Printf("insert failed, err:%v\n", err)
		HandleResponse(c, http.StatusBadRequest, "写入失败", err)
		return
	}
	HandleResponse(c, http.StatusOK, "写入成功", "OK")
	fmt.Println(affected) // 打印结果
}

func HandleResponse(c *gin.Context, code int, message string, data interface{}) {
	xormResponse.Code = code
	xormResponse.Message = message
	xormResponse.Data = data
	c.JSON(http.StatusOK, xormResponse)
}
