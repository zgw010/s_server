package main

import (
	// "strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

// RunData 表, 运动数据
type RunData struct {
	RunDataID     uuid.UUID `gorm:"primary_key"`
	RunDataUserID uuid.UUID
	Track         string `gorm:"type:text"`
	Calories      string
	Distance      string
	StartTime     string
	EndTime       string
	MaxSpeed      string
	AverageSpeed  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

func addRunData(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&RunData{})

	userID := c.PostForm("userID")
	track := c.PostForm(("track"))
	calories := c.PostForm(("calories"))
	distance := c.PostForm(("distance"))
	startTime := c.PostForm(("startTime"))
	endTime := c.PostForm(("endTime"))
	maxSpeed := c.PostForm(("maxSpeed"))
	averageSpeed := c.PostForm(("averageSpeed"))

	uuidUserID, err := uuid.FromString(userID)
	db.Create(&RunData{
		RunDataID:     uuid.NewV4(),
		RunDataUserID: uuidUserID,
		Track:         track,
		Calories:      calories,
		Distance:      distance,
		StartTime:     startTime,
		EndTime:       endTime,
		MaxSpeed:      maxSpeed,
		AverageSpeed:  averageSpeed,
		CreatedAt:     time.Now(),
	})
	c.PureJSON(200, gin.H{
		"status": 0,
	})
}

func getRunDataList(c *gin.Context) {
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
	var runDataList []RunData
	db.Where(
		"run_data_user_id = ? AND created_at BETWEEN ? AND ?",
		userID,
		t1,
		t2,
	).Order("updated_at desc").Find(&runDataList)
	c.PureJSON(200, gin.H{
		"status": 0,
		"data":   runDataList,
	})
}
func getRunData(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()

	runDataID := c.Query("runDataID")
	var runData RunData
	db.Where("run_data_id = ?", runDataID).First(&runData)
	c.PureJSON(200, gin.H{
		"status": 0,
		"data":   runData,
	})
}
