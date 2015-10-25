package main

import "github.com/gin-gonic/gin"
import "net/http"

func RenderRegisterForm(ctx *gin.Context, result *RegisterResult, user *SiteUser) {
  tplData := gin.H{
    "title": "Automata Theory - Lab 1, form validation",
    "alertMessage": "",
    "showAlertName": false,
    "showAlertEmail": false,
    "nickname": "",
    "email": "",
  }

  if result != nil && result.status != UserValid {
    tplData["alertMessage"] = result.message
    switch result.status {
    case UserInvalidNickname:
      tplData["showAlertName"] = true
    case UserInvalidEmail:
      tplData["showAlertEmail"] = true
    case UserWeakPassword:
      tplData["showAlertPassword"] = true
    case UserPasswordMismatch:
      tplData["showAlertPassword"] = true
    }
  }

  if user != nil {
    tplData["nickname"] = user.nickname
    tplData["email"] = user.email
  }

  ctx.HTML(http.StatusOK, "reg-form.tpl", tplData)
}

func RenderUserPage(ctx *gin.Context, user *SiteUser) {
  ctx.HTML(http.StatusOK, "reg-results.tpl", gin.H{
    "title": "Automata Theory - Lab 1, form validation",
    "userNickname": user.nickname,
    "userEmail": user.email,
  })
}

func RenderLifeGamePage(ctx *gin.Context) {
  ctx.HTML(http.StatusOK, "life-game.tpl", nil)
}

func main() {
  cache := NewSiteUsersCache()
  validator := new(RegisterFormValidator)

  router := gin.Default()
  router.Static("/css", "../site-content/css")
  router.Static("/js", "../site-content/js")
  router.LoadHTMLGlob("../site-content/tpl/*.tpl")
  router.GET("/form", func(ctx *gin.Context) {
    RenderRegisterForm(ctx, nil, nil)
  })
  router.GET("/life", func(ctx *gin.Context) {
    RenderLifeGamePage(ctx)
  })
  router.POST("/form", func(ctx *gin.Context) {
    user := &SiteUser{
      nickname: ctx.PostForm("userNickname"),
      email: ctx.PostForm("userEmail"),
      password: ctx.PostForm("userPassword"),
      passwordRepeat: ctx.PostForm("userPasswordRepeat"),
    }
    checkResult := validator.Check(user)
    if checkResult.status == UserValid {
      cache.AddUser(user)
      RenderUserPage(ctx, user)
    } else {
      //RenderRegisterForm(ctx, &checkResult, user)
      RenderRegisterForm(ctx, nil, nil)
    }
  })

  router.Run(":8080")
}
