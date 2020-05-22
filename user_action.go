package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

// UserAction 表, 基础动作库
type UserAction struct {
	ActionID      uuid.UUID `gorm:"primary_key"`
	ActionUserID  uuid.UUID `gorm:"primary_key"`
	ActionName    string
	ActionType    string
	ActionDetails string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

// Action 结构体，用来放用户动作和基础动作
type Action struct {
	ActionID       string
	ActionName     string
	ActionType     string
	ActionDetails  string
	ActionImgURL   string
	ActionVideoURL string
	ActionMoreURL  string
}

func addUserAction(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&UserAction{})

	userID := c.PostForm("userID")
	actionName := c.PostForm("actionName")
	actionDetails := c.PostForm("actionDetails")
	actionType := c.PostForm("actionType")
	uuidUserID, err := uuid.FromString(userID)
	id := uuid.NewV5(uuidUserID, actionName)
	var userAction UserAction
	userAction = UserAction{
		ActionID:      id,
		ActionUserID:  uuidUserID,
		ActionName:    actionName,
		ActionDetails: actionDetails,
		ActionType:    actionType,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	db.Create(userAction)
	c.PureJSON(200, gin.H{
		"status": 0,
		"data":   userAction,
	})
}

func updateUserAction(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&UserAction{})

	actionID := c.PostForm("actionID")
	actionName := c.PostForm("actionName")
	actionDetails := c.PostForm("actionDetails")
	actionType := c.PostForm("actionType")
	// uuidActionID, err := uuid.FromString(actionID)
	var userAction UserAction
	db.Where("action_id = ?", actionID).First(&userAction)
	db.Model(&userAction).Updates(map[string]interface{}{
		"ActionName":    actionName,
		"ActionType":    actionType,
		"ActionDetails": actionDetails,
	})
	c.PureJSON(200, gin.H{
		"status": 0,
		"data":   userAction,
	})
}

func deleteUserAction(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&UserAction{})

	actionID := c.PostForm("actionID")
	var userAction UserAction
	db.Where("action_id = ?", actionID).First(&userAction)
	db.Unscoped().Delete(&userAction)
	c.PureJSON(200, gin.H{
		"status": 0,
	})
}

func getActionList(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	var actionList []Action
	var userActionList []UserAction
	var baseActionList []BaseAction
	userID := c.Query("userID")
	actionType := c.Query("actionType")
	if actionType == "user" {
		db.Where("action_user_id = ?", userID).Order("updated_at desc").Find(&userActionList)
		c.PureJSON(200, gin.H{
			"status": 0,
			"data":   userActionList,
		})
	} else if actionType == "base" {
		db.Order("updated_at desc").Find(&baseActionList)
		c.PureJSON(200, gin.H{
			"status": 0,
			"data":   baseActionList,
		})
	} else {
		db.Order("updated_at desc").Find(&baseActionList)
		db.Where("action_user_id = ?", userID).Order("updated_at desc").Find(&userActionList)
		for _, v := range userActionList {
			action := Action{
				v.ActionID.String(),
				v.ActionName,
				v.ActionType,
				v.ActionDetails,
				"",
				"",
				"",
			}
			actionList = append(actionList, action)
		}
		for _, v := range baseActionList {

			actionIDString := strconv.FormatInt(v.ActionID, 10)
			action := Action{
				actionIDString,
				v.ActionName,
				v.ActionType,
				v.ActionDetails,
				v.ActionImgURL,
				v.ActionVideoURL,
				v.ActionMoreURL,
			}
			actionList = append(actionList, action)
		}
		// fmt.Println("actionList", actionList)
		c.PureJSON(200, gin.H{
			"status": 0,
			"data":   actionList,
		})
	}
}
