package controller

import (
	"bookManager/dao/mysql"
	"bookManager/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBookHandler 创建书籍添加数据
func CreateBookHandler(c *gin.Context) {
	p := new(model.Book)
	// 参数校验
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	mysql.DB.Create(p)
	c.JSON(200, gin.H{"msg": "success"})
}

// GetBookListHandler 查看书籍列表
func GetBookListHandler(c *gin.Context) {
	// 查找书籍
	var books []model.Book
	mysql.DB.Find(&books)
	c.JSON(200, gin.H{"books": books})
}

// GetBookDetailHandler 查看指定书籍 http://127.0.0.1:8000/book/1
func GetBookDetailHandler(c *gin.Context) {
	// 获取 bookId
	bookIdStr := c.Param("id")
	bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
	if err != nil {
		fmt.Print("ParseInt fail : ", err)
	}
	// 查询单条数据
	book := model.Book{Id: bookId}
	mysql.DB.Find(&book)
	c.JSON(200, gin.H{"book": book})
}

// UpdateBookHandler 修改书籍信息
func UpdateBookHandler(c *gin.Context) {
	p := new(model.Book)
	// 参数校验
	if err := c.ShouldBind(p); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
	}
	oldBook := &model.Book{Id: p.Id}
	var newBook model.Book
	if p.Name != "" {
		newBook.Name = p.Name
	}
	if p.Desc != "" {
		newBook.Desc = p.Desc
	}
	mysql.DB.Model(&oldBook).Updates(newBook)
	c.JSON(200, gin.H{"book": newBook})
}
