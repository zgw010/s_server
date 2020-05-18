package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// BaseAction 表, 基础动作库
type BaseAction struct {
	ActionID      int64 `gorm:"primary_key"`
	ActionName    string
	ActionType    string `gorm:"default:''"`
	ActionDetails string `gorm:"default:''"`
	ActionImgURL  string `gorm:"default:''"`
	// ActionImgURLList []string
	ActionVideoURL string `gorm:"default:''"`
	ActionMoreURL  string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
}

func addBaseAction(
	actionID int64,
	actionName string,
	actionType string,
	actionDetails string,
	actionImgURL string,
	actionVideoURL string,
	actionMoreURL string,
) string {
	db, err := gorm.Open("mysql", "root:19970705qq@(47.100.43.162)/zgw_s?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&BaseAction{})
	db.Create(&BaseAction{
		ActionID:       actionID,
		ActionName:     actionName,
		ActionType:     actionType,
		ActionDetails:  actionDetails,
		ActionImgURL:   actionImgURL,
		ActionVideoURL: actionVideoURL,
		ActionMoreURL:  actionMoreURL,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	return ""
}

func addSomeBaseAction() {
	addBaseAction(
		1,
		"深蹲",
		"base-times",
		"深蹲，又称蹲举，在力量练习中，是个复合的、全身性的练习动作，它可以训练到大腿、臀部、大腿后肌，同时可以增强骨头、韧带和横贯下半身的肌腱。深蹲被认为是增长腿部和臀部力量和围度，以及发展核心力量（core strength）必不可少的练习。在等长收缩中，在以正确的方式深蹲时，下背部、上背部、腹部、躯干肌肉，以及肋间肌肉，以及肩部和手臂对于这个练习都是必不可少的。 深蹲在力量举（健力）中是一种竞争性的上举",
		"https://static1.gotokeep.com/picture/frame/1544524199280.jpg",
		"https://static1.keepcdn.com/chaos/0728/B072C009_main_s.mp4",
		"https://www.gotokeep.com/exercises/5763d3b011fc5077c3acf5be?gender=f",
	)
	addBaseAction(
		2,
		"硬拉",
		"base-times",
		"硬拉（英语：Deadlift），又称拉举、硬举，是一种负重训练，主要用于锻炼下背部即竖脊肌、臀部肌肉和大腿肌肉，是世界力量举重锦标赛（WWC, World Weightlifting Championship）的项目之一，另外两项是深蹲和卧推。",
		"https://static1.gotokeep.com/picture/frame/1541579014755.jpg",
		"https://static1.keepcdn.com/chaos/0728/A076C038_main_s.mp4",
		"https://www.gotokeep.com/exercises/5763d3b011fc5077c3acf589?gender=f",
	)
	addBaseAction(
		3,
		"卧推",
		"base-times",
		"仰卧推举（英语：bench press），简称卧推，又称推举，是上身训练中的一种。若以健美为目的，这种训练被用于增强胸肌、三角肌与三头肌。当做卧推时，需要先仰卧，双手将重物压低到胸部所处的水平线以下，接下来将重物向上推直到手臂伸直。此项训练主要着重于发展胸大肌，辅助完成此项动作的其他肌肉也同时得到了锻炼，其中包括前三角肌、前锯肌、啄肱肌、肩部肌肉、斜方肌以及肱三头肌。卧推是健力运动中三项举之一且被广泛地用于负重训练、健美以及其他类型的健身训练中用以发展胸部肌肉。",
		"https://static1.gotokeep.com/picture/frame/1523934255745.jpg",
		"https://static1.keepcdn.com/chaos/0728/A075C022_main_s.mp4",
		"https://www.gotokeep.com/exercises/5763d3b011fc5077c3acf5b0?gender=f",
	)
	addBaseAction(
		4,
		"跑步",
		"base-time",
		"跑步是日常方便的一种体育锻炼方法，是有氧呼吸的有效运动方式。",
		"https://static1.gotokeep.com/picture/frame/1500975405902.jpg",
		"https://static1.keepcdn.com/chaos/0728/B058C062_main_s.mp4",
		"https://www.gotokeep.com/exercises/5979cae611fc50467bd124e9?gender=f",
	)
	addBaseAction(
		5,
		"平板支撑",
		"base-onlytime",
		"跑步是日常方便的一种体育锻炼方法，是有氧呼吸的有效运动方式。",
		"https://static1.gotokeep.com/picture/frame/1500881848788.jpg",
		"https://static1.keepcdn.com/chaos/0728/A031C068_main_s.mp4",
		"https://www.gotokeep.com/exercises/55cc42d9fdac76af7fc278a9?gender=f",
	)
}
