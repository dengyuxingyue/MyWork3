package service

//用户通过提交form表单，经post方法实例化UserService
//routes->form->实例化service->Login()||Register操作model与数据库交互
import (
	"work/model"
	"work/pkg/utils"
	"work/serializer"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "已经有这个人了，无需在进行注册",
		}
	}

	user.UserName = service.UserName

	//密码加密
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功了",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	//验证用户存在
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在",
			}
		}

		//如果用户不存在，是其它不可抗拒的因素导致的错误
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	//对密码进行验证
	if !user.CheckPassword(service.Password) {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "token签发错误",
		}
	}

	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: "登录成功",
	}

}
