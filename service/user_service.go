package service

import (
	"GoLangStudy/domain"
	"GoLangStudy/global"
	"GoLangStudy/model/entity"
	"GoLangStudy/model/request"
	"gorm.io/gorm"
)

const USER_NAME = "admin"
const PASS_WORD = "123456"

func UserLogin(loginPo domain.LoginPo) bool {
	userName := loginPo.UserName
	password := loginPo.Password
	if userName == USER_NAME && password == PASS_WORD {
		return true
	} else {
		return false
	}
}

func Register(param request.RegisterParam) (*entity.User, error) {
	user := entity.User{
		Phone:    param.Phone,
		Password: param.Password,
		NickName: param.NickName,
	}
	global.GvaMysqlClient.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		return nil
	})
	return &user, nil
}
