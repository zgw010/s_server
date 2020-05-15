package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

// Plan 表, 用户运动计划
type Plan struct {
	PlanID        uuid.UUID `gorm:"primary_key"`
	PlanUserID    uuid.UUID
	PlanName      string
	PlanDetails   string `gorm:"type:text"`
	PlanCompleted bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

// PlanGroup 表, 用户运动计划
type PlanGroup struct {
	PlanGroupID      uuid.UUID `gorm:"primary_key"`
	PlanGroupUserID  uuid.UUID
	PlanGroupName    string
	PlanIDs          string
	PlanGroupDetails string `gorm:"type:text"`
	PlanGroupStep    int32
	PlanGroupTimes   int32
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time
}

func addPlan(
	userName string,
	planName string,
	planDetails string,
) string {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&Plan{})
	var user User
	db.Where("user_name = ?", userName).First(&user)
	userID := user.UserID
	if userID.String() == "" {
		return "not find user"
	}
	id := uuid.NewV5(userID, planName)

	planDetailsByte := []byte(planDetails)
	var planDetailsJSON []map[string]string
	json.Unmarshal(planDetailsByte, &planDetailsJSON)
	fmt.Println("planDetailsJSON", planDetailsJSON[0]["actionType"])
	fmt.Println("planDetailsJSON", planDetailsJSON[0] != nil && planDetailsJSON[0]["actionType"] != "")

	db.Create(&Plan{
		PlanID:        id,
		PlanUserID:    userID,
		PlanName:      planName,
		PlanDetails:   planDetails,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		PlanCompleted: false,
	})
	return "ok"
}

func addPlanGroup(
	userName string,
	planGroupName string,
	planIDs string,
	planDetails string,
) string {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&PlanGroup{})
	var user User
	db.Where("user_name = ?", userName).First(&user)
	userID := user.UserID
	if userID.String() == "" {
		return "not find user"
	}
	id := uuid.NewV5(userID, planGroupName)

	planIDsByte := []byte(planIDs)
	var planIDsJSON []string
	json.Unmarshal(planIDsByte, &planIDsJSON)
	fmt.Println("planDetailsJSON", reflect.TypeOf(planIDsJSON[0]))
	fmt.Println("planDetailsJSON", planIDsJSON[0])
	db.Create(&PlanGroup{
		PlanGroupID:      id,
		PlanGroupUserID:  userID,
		PlanGroupName:    planGroupName,
		PlanIDs:          planIDs,
		PlanGroupDetails: planDetails,
		PlanGroupTimes:   0,
		PlanGroupStep:    0,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	})
	return "ok"
}

func getPlanGroupList(
	userID string,
) []PlanGroup {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	var planGroupList []PlanGroup
	db.Where("plan_group_user_id = ?", userID).Find(&planGroupList)
	return planGroupList
}

func getPlanGroup(
	planGroupID string,
) PlanGroup {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	var planGroup PlanGroup
	db.Where("plan_group_id = ?", planGroupID).First(&planGroup)
	planIDs := planGroup.PlanIDs

	planIDsByte := []byte(planIDs)
	var planIDsJSON []string
	json.Unmarshal(planIDsByte, &planIDsJSON)
	fmt.Println("planDetailsJSON", planIDsJSON[0])

	return planGroup
}

func getCurPlan(
	planGroupID string,
) Plan {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	var planGroup PlanGroup
	db.Where("plan_group_id = ?", planGroupID).First(&planGroup)
	fmt.Println("planGroup", planGroup)
	planIDs := planGroup.PlanIDs
	planIDsByte := []byte(planIDs)
	var planIDsJSON []string
	json.Unmarshal(planIDsByte, &planIDsJSON)
	fmt.Println("planDetailsJSON", planIDsJSON[planGroup.PlanGroupStep])

	var plan Plan
	db.Where("plan_id = ?", planIDsJSON[planGroup.PlanGroupStep]).First(&plan)

	return plan
}
