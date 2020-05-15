package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	// type Message struct {
	// 	Name string
	// 	Body string
	// 	Time int64
	// }
	// m := Message{"Alice", "Hello", 1294706395881547000}
	// b, err := json.Marshal(m)
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// fmt.Println("b", b)
	// fmt.Println("b", reflect.TypeOf(b))
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	//  gin.DisableConsoleColor()

	// 记录到文件。
	// f, _ := os.Create("gin.log")
	//  gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	// addSomeBaseAction()
	router.POST("/register", func(c *gin.Context) {
		userName := c.PostForm("userName")
		userPassword := c.PostForm("userPassword")
		userAvatarURL := c.PostForm("userAvatarURL")
		userDetails := c.PostForm("userDetails")
		userDateOfBirth := c.PostForm("userDateOfBirth")
		userHeight := c.PostForm("userHeight")
		userWeight := c.PostForm("userWeight")
		userAims := c.PostForm("userAims")
		var addRes = addUser(
			userName,
			userPassword,
			userAvatarURL,
			userDetails,
			userDateOfBirth,
			userHeight,
			userWeight,
			userAims,
		)
		if addRes == "ok" {
			c.PureJSON(200, gin.H{
				"status": 0,
			})
		} else {
			c.PureJSON(200, gin.H{
				"status": 1,
			})
		}
		// c.String(200, addRes)
	})

	router.POST("/login", func(c *gin.Context) {
		userName := c.PostForm("userName")
		userPassword := c.PostForm("userPassword")
		var queryRes = queryUser(
			userName,
			userPassword,
		)
		if queryRes != "" {
			c.PureJSON(200, gin.H{
				"status": 0,
				"cookie": queryRes,
			})
		} else {
			c.PureJSON(200, gin.H{
				"status": 1,
			})
		}
		// c.String(200, addRes)
	})

	router.POST("/action/add_user_action", func(c *gin.Context) {
		userName := c.PostForm("userName")
		actionName := c.PostForm("actionName")
		actionDetails := c.PostForm("actionDetails")
		addRes := addUserAction(
			userName,
			actionName,
			actionDetails,
		)
		if addRes != "" {
			c.PureJSON(200, gin.H{
				"status": 0,
				"data":   addRes,
			})
		} else {
			c.PureJSON(200, gin.H{
				"status": 1,
			})
		}
	})
	router.POST("/plan/add_plan", func(c *gin.Context) {
		userName := c.PostForm("userName")
		planName := c.PostForm("planName")
		planDetails := c.PostForm("planDetails")
		addRes := addPlan(
			userName,
			planName,
			planDetails,
		)
		if addRes != "" {
			c.PureJSON(200, gin.H{
				"status": 0,
				"data":   addRes,
			})
		} else {
			c.PureJSON(200, gin.H{
				"status": 1,
			})
		}
	})
	router.POST("/plan/add_plan_group", func(c *gin.Context) {
		userName := c.PostForm("userName")
		planGroupName := c.PostForm("planGroupName")
		planIDs := c.PostForm("planIDs")
		planGroupDetails := c.PostForm("planGroupDetails")
		addRes := addPlanGroup(
			userName,
			planGroupName,
			planIDs,
			planGroupDetails,
		)
		if addRes != "" {
			c.PureJSON(200, gin.H{
				"status": 0,
				"data":   addRes,
			})
		} else {
			c.PureJSON(200, gin.H{
				"status": 1,
			})
		}
	})
	router.GET("/plan/get_plan_group_list", func(c *gin.Context) {
		userID := c.Query("userID")
		planGroupList := getPlanGroupList(userID)
		fmt.Println("planGroupList", planGroupList)
		// if addRes != "" {
		// 	c.PureJSON(200, gin.H{
		// 		"status": 0,
		// 		"data":   addRes,
		// 	})
		// } else {
		// 	c.PureJSON(200, gin.H{
		// 		"status": 1,
		// 	})
		// }
		c.PureJSON(200, gin.H{
			"status": 0,
			"data":   planGroupList,
		})
	})
	router.GET("/plan/get_plan_group", func(c *gin.Context) {
		planGroupID := c.Query("planGroupID")
		planGroup := getPlanGroup(planGroupID)
		fmt.Println("planGroup", planGroup)
		c.PureJSON(200, gin.H{
			"status": 0,
			"data":   planGroup,
		})
	})
	router.GET("/plan/get_plan", func(c *gin.Context) {
		planGroupID := c.Query("planGroupID")
		plan := getCurPlan(planGroupID)
		fmt.Println("plan", plan)
		c.PureJSON(200, gin.H{
			"status": 0,
			"data":   plan,
		})
	})

	router.Run(":8080")
}
