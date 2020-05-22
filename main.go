package main

import (
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
	// user
	router.POST("/register", func(c *gin.Context) { addUser(c) })
	router.POST("/login", func(c *gin.Context) { queryUser(c) })
	router.POST("/user/update_plan_group_id", func(c *gin.Context) { updateUserPlanGroupID(c) })

	// action
	router.GET("/action/get_action_list", func(c *gin.Context) { getActionList(c) })
	router.POST("/action/add_user_action", func(c *gin.Context) { addUserAction(c) })
	router.POST("/action/update_user_action", func(c *gin.Context) { updateUserAction(c) })
	router.POST("/action/delete_user_action", func(c *gin.Context) { deleteUserAction(c) })

	// plan
	router.GET("/plan/get_cur_plan", func(c *gin.Context) { getCurPlan(c) })
	router.GET("/plan/get_plan_list", func(c *gin.Context) { getPlanList(c) })
	router.POST("/plan/add_plan", func(c *gin.Context) { addPlan(c) })
	router.POST("/plan/update_plan", func(c *gin.Context) { updatePlan(c) })
	router.POST("/plan/delete_plan", func(c *gin.Context) { deletePlan(c) })

	router.GET("/plan/get_plan_group_list", func(c *gin.Context) { getPlanGroupList(c) })
	router.GET("/plan/get_plan_group", func(c *gin.Context) { getPlanGroup(c) })
	router.POST("/plan/add_plan_group", func(c *gin.Context) { addPlanGroup(c) })
	router.POST("/plan/update_plan_group", func(c *gin.Context) { updatePlanGroup(c) })
	router.POST("/plan/delete_plan_group", func(c *gin.Context) { deletePlanGroup(c) })

	// data
	router.GET("/data/get_data_list", func(c *gin.Context) { getDataList(c) })
	router.GET("/data/get_data", func(c *gin.Context) { getData(c) })
	router.POST("/data/add_data", func(c *gin.Context) { addData(c) })

	router.GET("/data/get_run_data_list", func(c *gin.Context) { getRunDataList(c) })
	router.GET("/data/get_run_data", func(c *gin.Context) { getRunData(c) })
	router.POST("/data/add_run_data", func(c *gin.Context) { addRunData(c) })

	router.Run(":8080")
}
