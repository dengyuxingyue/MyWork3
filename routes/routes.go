package routes

import (
	"work/api"
	"work/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	//建立路由组
	v1 := r.Group("api/v1")
	{
		//注册
		v1.POST("user/register", api.UserRegister)
		//登录
		v1.POST("user/login", api.UserLogin)

		//进行身份认证
		//分组嵌套
		authed := v1.Group("/")
		//群组中间件
		authed.Use(middleware.JWT())
		//执行路由时，会先验证token
		{
			//创建一条备忘录
			authed.POST("task", api.CreateTask)
			//展示一条备忘录
			authed.GET("task/:id", api.ShowTask)
			//展示用户所有的备忘录
			authed.GET("tasks", api.ShowAlltasks)
			//更新一条
			authed.PUT("task/:id", api.UpdateTask)
			//关键词搜索备忘录
			authed.POST("search", api.SearchTask)
			//删除一条备忘录
			authed.DELETE("task/:id", api.DeleteTask)
			//查看所有未完成的备忘录
			authed.POST("tasks/incompleted", api.GetInCompletedTasks)
			//查看所有已完成的备忘录
			authed.POST("tasks/completed", api.GetCompletedTasks)
		}
	}
	return r
}
