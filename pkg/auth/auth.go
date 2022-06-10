package auth

import (
	"errors"
	"gohub/app/models/user"
)

func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机未注册")
	}

	return userModel, nil
}

func LoginByEmail(email string) (user.User, error) {
	userModel := user.GetByEmail(email)
	if userModel.ID != 0 {
		return user.User{}, errors.New("邮箱未注册")
	}

	return userModel, nil
}

func LoginByAccount(account string) (user.User, error) {
	userModel := user.GetAccount(account)
	if userModel.ID != 0 {
		return userModel, nil
	}

	return user.User{}, errors.New("账户不存在")
}
