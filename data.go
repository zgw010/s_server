package main

import (
	// "strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

// Data 表, 运动数据
type Data struct {
	DataID      uuid.UUID `gorm:"primary_key"`
	DataUserID  uuid.UUID
	DataName    string
	DataType    string
	DataTime    string
	DataDetails string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

func addData(c *gin.Context) bool {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&Data{})

	userID := c.PostForm("userID")
	dataName := c.PostForm("dataName")
	dataTime := c.PostForm("dataTime")
	dataType := c.PostForm("dataType")
	dataDetails := c.PostForm("dataDetails")
	uuidUserID, err := uuid.FromString(userID)
	db.Create(&Data{
		DataID:      uuid.NewV4(),
		DataUserID:  uuidUserID,
		DataName:    dataName,
		DataTime:    dataTime,
		DataType:    dataType,
		DataDetails: dataDetails,
		CreatedAt:   time.Now(),
	})
	return true
}

func getDataList(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()

	userID := c.Query("userID")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	t1, _ := time.Parse(time.RFC3339, startTime+"T00:00:00Z")
	t2, _ := time.Parse(time.RFC3339, endTime+"T23:59:59Z")
	var dataList []Data
	db.Where(
		"data_user_id = ? AND created_at BETWEEN ? AND ?",
		userID,
		t1,
		t2,
	).Find(&dataList)
	c.PureJSON(200, gin.H{
		"status": 0,
		"data":   dataList,
	})
}
func getData(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()

	userID := c.Query("userID")
	var data Data
	db.Where("data_user_id = ?", userID).First(&data)
	c.PureJSON(200, gin.H{
		"status": 0,
		"data":   data,
	})
}
