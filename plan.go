package main

import (
	"encoding/json"
	"fmt"

	// "reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

// Plan 表, 用户运动计划
type Plan struct {
	PlanID      uuid.UUID `gorm:"primary_key"`
	PlanUserID  uuid.UUID
	PlanName    string
	PlanDetails string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

// PlanGroup 表, 用户运动计划
type PlanGroup struct {
	PlanGroupID       uuid.UUID `gorm:"primary_key"`
	PlanGroupUserID   uuid.UUID
	PlanGroupName     string
	PlanGroupDetails  string `gorm:"type:text"`
	PlanGroupStep     int
	PlanGroupTimes    int32
	PlanCompletedTime string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

func getCurPlan(c *gin.Context) {
	userID := c.Query("userID")
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	var planGroup PlanGroup
	var user User
	var plan Plan

	db.Where("user_id = ?", userID).First(&user)
	planGroupID := user.UserPlanGroupID
	if planGroupID != "" {
		db.Where("plan_group_id = ?", planGroupID).First(&planGroup)
		planGroupDetails := planGroup.PlanGroupDetails
		planIDsByte := []byte(planGroupDetails)
		var planGroupDetailsJSON []map[string]string
		json.Unmarshal(planIDsByte, &planGroupDetailsJSON)
		// fmt.Println(planGroupDetailsJSON)
		db.Where("plan_id = ?", planGroupDetailsJSON[planGroup.PlanGroupStep]["PlanID"]).First(&plan)
		c.PureJSON(200, gin.H{
			"status":        0,
			"data":          plan,
			"planGroupData": planGroup,
		})
	} else {
		var planGroupList []PlanGroup
		db.Where("plan_group_user_id = ?", userID).Find(&planGroupList)
		if len(planGroupList) > 0 { // 有计划组，但未制定当前计划组
			c.PureJSON(200, gin.H{
				"status": 10001,
			})
		} else { // 用户未制定任何计划组
			c.PureJSON(200, gin.H{
				"status": 10002,
			})
		}
	}
}

func addPlan(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&Plan{})

	userID := c.PostForm("userID")
	planName := c.PostForm("planName")
	planDetails := c.PostForm("planDetails")
	uuidUserID, err := uuid.FromString(userID)
	id := uuid.NewV5(uuidUserID, planName)
	plan := Plan{
		PlanID:      id,
		PlanUserID:  uuidUserID,
		PlanName:    planName,
		PlanDetails: planDetails,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	db.Create(plan)
	c.PureJSON(200, gin.H{
		"status": 0,
		"data":   plan,
	})
}

func updatePlan(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()

	var plan Plan
	planID := c.PostForm("planID")
	planName := c.PostForm("planName")
	planDetails := c.PostForm("planDetails")
	db.Where("plan_id = ?", planID).First(&plan)
	db.Model(&plan).Updates(map[string]interface{}{
		"PlanName":    planName,
		"PlanDetails": planDetails,
	})
	c.PureJSON(200, gin.H{
		"status": 0,
		"data":   plan,
	})
}

func deletePlan(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()

	planID := c.PostForm("planID")
	var plan Plan
	db.Where("plan_id = ?", planID).First(&plan)
	db.Delete(&plan)
	c.PureJSON(200, gin.H{
		"status": 0,
	})
}

func getPlanList(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()

	userID := c.Query("userID")
	var planList []Plan
	db.Where("plan_user_id = ?", userID).Order("updated_at desc").Find(&planList)
	c.PureJSON(200, gin.H{
		"status": 0,
		"data":   planList,
	})
}

func getPlanGroupList(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	userID := c.Query("userID")

	var planGroupList []PlanGroup
	db.Where("plan_group_user_id = ?", userID).Find(&planGroupList)
	c.PureJSON(200, gin.H{
		"status": 0,
		"data":   planGroupList,
	})
}

func getPlanGroup(c *gin.Context) PlanGroup {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	var planGroup PlanGroup
	// db.Where("plan_group_id = ?", planGroupID).First(&planGroup)

	return planGroup
}

func addPlanGroup(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&PlanGroup{})

	userID := c.PostForm("userID")
	planGroupName := c.PostForm("planGroupName")
	planGroupDetails := c.PostForm("planGroupDetails")

	uuidUserID, err := uuid.FromString(userID)
	id := uuid.NewV5(uuidUserID, planGroupName)

	planGroup := PlanGroup{
		PlanGroupID:      id,
		PlanGroupUserID:  uuidUserID,
		PlanGroupName:    planGroupName,
		PlanGroupDetails: planGroupDetails,
		PlanGroupTimes:   0,
		PlanGroupStep:    0,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	db.Create(planGroup)
	c.PureJSON(200, gin.H{
		"status": 0,
		"data":   planGroup,
	})
}

func updatePlanGroup(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()

	var planGroup PlanGroup
	planGroupID := c.PostForm("planGroupID")
	planGroupName := c.PostForm("planGroupName")
	planGroupDetails := c.PostForm("planGroupDetails")
	updateType := c.PostForm("updateType")

	db.Where("plan_group_id = ?", planGroupID).First(&planGroup)
	if updateType == "complete" {
		var planGroupDetails = planGroup.PlanGroupDetails
		planGroupDetailsByte := []byte(planGroupDetails)
		var planGroupDetailsJSON []map[string]string
		json.Unmarshal(planGroupDetailsByte, &planGroupDetailsJSON)
		fmt.Println(planGroupDetailsJSON)
		if planGroup.PlanGroupStep >= len(planGroupDetailsJSON)-1 {
			db.Model(&planGroup).Updates(map[string]interface{}{
				"PlanGroupStep":     0,
				"PlanGroupTimes":    planGroup.PlanGroupTimes + 1,
				"PlanCompletedTime": time.Now().Format("2006-01-02"),
			})
		} else {
			db.Model(&planGroup).Updates(map[string]interface{}{
				"PlanGroupStep":     planGroup.PlanGroupStep + 1,
				"PlanCompletedTime": time.Now().Format("2006-01-02"),
			})
		}

	} else {
		db.Model(&planGroup).Updates(map[string]interface{}{
			"PlanGroupName":    planGroupName,
			"PlanGroupDetails": planGroupDetails,
		})
	}
	c.PureJSON(200, gin.H{
		"status": 0,
		"data":   planGroup,
	})
}

func deletePlanGroup(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()

	planGroupID := c.PostForm("planGroupID")
	var planGroup PlanGroup
	db.Where("plan_group_id = ?", planGroupID).First(&planGroup)
	db.Delete(&planGroup)
	c.PureJSON(200, gin.H{
		"status": 0,
	})
}
