package main

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

// User 用户表
type User struct {
	UserID          uuid.UUID `gorm:"primary_key"`
	UserName        string
	UserPassword    string
	UserAvatarURL   string
	UserDetails     string
	UserDateOfBirth string
	UserHeight      string
	UserWeight      string
	UserSex         string
	UserAims        string
	PlanGroupID     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

func queryUser(
	c *gin.Context,
) {
	userName := c.PostForm("userName")
	userPassword := c.PostForm("userPassword")
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	var user User
	db.Where("user_name = ? AND user_password = ?", userName, userPassword).First(&user)
	if user.UserName != "" {
		userJSON, err := json.Marshal(user)
		if err == nil {
			c.PureJSON(200, gin.H{
				"status":   0,
				"userInfo": string(userJSON),
			})
			return
		}
		c.PureJSON(200, gin.H{
			"status": 1,
		})
		return
	}
	c.PureJSON(200, gin.H{
		"status": 1,
	})
	return
}

func addUser(
	c *gin.Context,
) {
	userName := c.PostForm("userName")
	userPassword := c.PostForm("userPassword")
	userAvatarURL := c.PostForm("userAvatarURL")
	userDetails := c.PostForm("userDetails")
	userDateOfBirth := c.PostForm("userDateOfBirth")
	userHeight := c.PostForm("userHeight")
	userWeight := c.PostForm("userWeight")
	userAims := c.PostForm("userAims")
	userSex := c.PostForm("userSex")
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&User{})
	// var user User
	// 获取第一个匹配的记录
	// db.Where("user_name = ?", userName).First(&user)
	// if user.UserName != "" {
	// 	c.PureJSON(200, gin.H{
	// 		"status": 1,
	// 	})
	// 	return
	// }
	id := uuid.NewV4()
	newUser := User{
		UserID:          id,
		UserName:        userName,
		UserPassword:    userPassword,
		UserAvatarURL:   userAvatarURL,
		UserDetails:     userDetails,
		UserDateOfBirth: userDateOfBirth,
		UserSex:         userSex,
		UserHeight:      userHeight,
		UserWeight:      userWeight,
		UserAims:        userAims,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	db.Create(newUser)
	userJSON, err := json.Marshal(newUser)
	if err == nil {
		c.PureJSON(200, gin.H{
			"status":   0,
			"userInfo": string(userJSON),
		})
		return
	}
	c.PureJSON(200, gin.H{
		"status": 1,
	})

}
func deleteUser() {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&User{})
	// 创建
	id := uuid.NewV4()
	db.Create(&User{
		UserID:          id,
		UserName:        "zgw",
		UserAvatarURL:   "https://i.loli.net/2020/05/04/Rf6ze9LFQMxgShu.jpg",
		UserDetails:     "眨巴",
		UserDateOfBirth: "1997-8-7",
		UserHeight:      "184",
		UserWeight:      "80",
		UserAims:        "body build",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})
}

func updateUser(
	c *gin.Context,
) {
	userID := c.PostForm("userID")
	userName := c.PostForm("userName")
	userPassword := c.PostForm("userPassword")
	userAvatarURL := c.PostForm("userAvatarURL")
	userDetails := c.PostForm("userDetails")
	userDateOfBirth := c.PostForm("userDateOfBirth")
	userHeight := c.PostForm("userHeight")
	userWeight := c.PostForm("userWeight")
	userAims := c.PostForm("userAims")
	userSex := c.PostForm("userSex")

	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()

	// 使用 map 更新多个属性，只会更新其中有变化的属性
	var user User
	db.Where("user_id = ?", userID).First(&user)
	db.Model(&user).Updates(map[string]interface{}{
		"UserName":        userName,
		"UserPassword":    userPassword,
		"UserAvatarURL":   userAvatarURL,
		"UserDetails":     userDetails,
		"UserDateOfBirth": userDateOfBirth,
		"UserHeight":      userHeight,
		"UserWeight":      userWeight,
		"UserSex":         userSex,
		"UserAims":        userAims,
		"UpdatedAt":       time.Now(),
	})
	userJSON, err := json.Marshal(user)
	if err == nil {
		c.PureJSON(200, gin.H{
			"status":   0,
			"userInfo": string(userJSON),
		})
		return
	}
	c.PureJSON(200, gin.H{
		"status": 1,
	})
	return
}

func updateUserPlanGroupID(
	c *gin.Context,
) {
	userID := c.PostForm("userID")
	userPlanGroupID := c.PostForm("userPlanGroupID")
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	var user User
	db.Where("user_id = ?", userID).First(&user)
	db.Model(&user).Updates(map[string]interface{}{
		"UserPlanGroupID": userPlanGroupID,
		"UpdatedAt":       time.Now(),
	})
	if err == nil {
		c.PureJSON(200, gin.H{
			"status": 0,
		})
	} else {
		c.PureJSON(200, gin.H{
			"status": 1,
		})
	}
}
