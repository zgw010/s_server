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
	router.POST("/register", func(c *gin.Context) { addUser(c) })

	router.POST("/login", func(c *gin.Context) { queryUser(c) })

	router.POST("/user/update_plan_group_id", func(c *gin.Context) { updateUserPlanGroupID(c) })

	router.GET("/action/get_action_list", func(c *gin.Context) { getActionList(c) })
	router.POST("/action/add_user_action", func(c *gin.Context) { addUserAction(c) })

	router.GET("/plan/get_plan_list", func(c *gin.Context) { getPlanList(c) })

	router.POST("/plan/add_plan", func(c *gin.Context) { addPlan(c) })
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
	router.GET("/plan/get_cur_plan", func(c *gin.Context) { getCurPlan(c) })

	router.Run(":8080")
}
