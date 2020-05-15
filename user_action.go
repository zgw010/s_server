package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

// UserAction 表, 基础动作库
type UserAction struct {
	ActionID      uuid.UUID `gorm:"primary_key"`
	ActionName    string
	ActionType    string
	ActionDetails string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

func addUserAction(
	userName string,
	actionName string,
	actionDetails string,
) string {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&UserAction{})
	var user User
	db.Where("user_name = ?", userName).First(&user)
	userID := user.UserID
	if userID.String() == "" {
		return "not find user"
	}
	// fmt.Println("userID", userID)
	id := uuid.NewV5(userID, actionName)
	// fmt.Println("id", id)
	db.Create(&UserAction{
		ActionID:      id,
		ActionName:    actionName,
		ActionDetails: actionDetails,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		ActionType:    "user",
	})
	return "ok"
}
