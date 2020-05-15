package main

import (
	"fmt"
	"time"

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
	UserAims        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time `gorm:"default:'NULL'"`
}

func queryUser(
	userName string,
	userPassword string,
) string {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	// 查询用户
	var user User
	// 获取第一个匹配的记录
	db.Where("user_name = ? AND user_password >= ?", userName, userPassword).First(&user)
	fmt.Println("user", user)
	fmt.Println("user", user.UserName)
	if user.UserName != "" {
		return user.UserID.String()
	}
	return ""
}

func addUser(
	userName string,
	userPassword string,
	userAvatarURL string,
	userDetails string,
	userDateOfBirth string,
	userHeight string,
	userWeight string,
	userAims string,
) string {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&User{})
	var user User
	// 获取第一个匹配的记录
	db.Where("user_name = ?", userName).First(&user)
	if user.UserName != "" {
		return ""
	}
	id := uuid.NewV4()
	db.Create(&User{
		UserID:          id,
		UserName:        userName,
		UserPassword:    userPassword,
		UserAvatarURL:   userAvatarURL,
		UserDetails:     userDetails,
		UserDateOfBirth: userDateOfBirth,
		UserHeight:      userHeight,
		UserWeight:      userWeight,
		UserAims:        userAims,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})
	return "ok"
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
func updateUser() {
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
		DeletedAt:       time.Now(),
	})
}
