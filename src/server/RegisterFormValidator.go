package main

import (
		"regexp"
		"strings"
)

const (
  UserValid = iota
  UserInvalidNickname
  UserInvalidEmail
  UserWeakPassword
  UserPasswordMismatch
)

type RegisterResult struct {
  status int32
  message string
}

type RegisterFormValidator struct {
}

func (self *RegisterFormValidator) Check(user *SiteUser) RegisterResult {
	result := RegisterResult{
		status: UserValid,
	}

	self.CheckNickname(user.nickname, &result)
	self.CheckEmail(user.email, &result)
	self.CheckPassword(user.password, &result)
	self.CheckRepeatPassword(user.password, user.passwordRepeat, &result)

  return result
}

func (self *RegisterFormValidator) CheckNickname(nickname string, result *RegisterResult) {
	if result.status != UserValid {
		return
	}

	result.status = UserInvalidNickname

  if len(nickname) == 0 {
    result.message = "nickname is empty"
  } else if len(nickname) < 3 {
		result.message = "nickname is too short, enter at least 3 symbols"
	} else if matched, err := regexp.MatchString("^[a-zA-Z0-9_]{3,}$", nickname); err == nil && !matched {
		result.message = "nickname contains not allowed symbols"
	} else {
		result.status = UserValid
	}
}

func (self *RegisterFormValidator) CheckEmail(email string, result *RegisterResult) {
	if result.status != UserValid {
		return
	}

	result.status = UserInvalidEmail

	if len(email) == 0 {
		result.message = "email is empty"
	} else if !strings.Contains(email, "@") {
		result.message = "email shoud contain @"
	} else {
		var emailParts []string = strings.Split(email, "@")
		if len(emailParts) != 2 {
			result.message = "email should contain only one @"
		} else if matched, err := regexp.MatchString("^[a-zA-Z0-9]{1,}$", emailParts[0]); err == nil && !matched {
			result.message = "email contains not allowed symbols"
		} else if matched, err := regexp.MatchString("^(gmail.com|yandex.ru|mail.ru)$", emailParts[1]); err == nil && !matched {
			result.message = "domain is not from allowed list"
		} else {
			result.status = UserValid
		}
	}
}

func (self *RegisterFormValidator) CheckPassword(password string, result *RegisterResult) {
	if result.status != UserValid {
		return
	}

	result.status = UserWeakPassword

	if len(password) == 0 {
		result.message = "empty password"
		} else if len(password) < 6 {
			result.message = "password is too short, enter at least 6 symbols"
		} else if matched, err := regexp.MatchString("^[a-zA-Z0-9]{6,}$", password); err == nil && !matched {
			result.message = "password contains not allowed symbols"
		} else if matched, err := regexp.MatchString("[a-zA-Z]", password); err == nil && !matched {
			result.message = "password should contain at least 1 latin letter"
		} else if matched, err := regexp.MatchString("[0-9]", password); err == nil && !matched {
			result.message = "password should contain at least 1 digit"
		} else {
			result.status = UserValid
		}
}

func (self *RegisterFormValidator) CheckRepeatPassword(newPassword string, repeatPassword string, result *RegisterResult) {
	if result.status != UserValid {
		return
	}

	if newPassword == repeatPassword {
		result.status = UserValid
	} else {
		result.status = UserPasswordMismatch
		result.message = "passwords are not matching"
	}
}
