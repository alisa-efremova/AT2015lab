<!DOCTYPE html>
<html lang="ru">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="/css/bootstrap.min.css">
    <link rel="stylesheet" href="/css/bootstrap-theme.min.css">
    <script type="text/javascript" src="/js/form_validation.js"></script>
  </head>
  <body>
    <div class="container">
      <div class="row">
        <h2>
          {{.title}}
        </h2>
      </div>
      <div class="row">
        <form action="/form" method="POST" id="regForm">
          <div class="form-group">
            <label for="inputNickname">Nickname</label>
            {{if .showAlertName}}
              <div class="alert alert-danger" role="alert">{{.alertMessage}}</div>
            {{end}}
            <div class="alert alert-danger" style="display: none;" id="nicknameErrMessage"></div>
            <input type="text" class="form-control" id="inputNickname" placeholder="Your Nickname" name="userNickname" value="{{.nickname}}">
            <p class="help-block">Nickname should contain only English letters, digits and underscores.</p>
          </div>
          <div class="form-group">
            <label for="inputEmail">Email address</label>
            {{if .showAlertEmail}}
              <div class="alert alert-danger" role="alert">{{.alertMessage}}</div>
            {{end}}
            <div class="alert alert-danger" style="display: none;" id="emailErrMessage"></div>
            <input type="text" class="form-control" id="inputEmail" placeholder="Your Email" name="userEmail" value="{{.email}}">
            <p class="help-block">Only GMail, Yandex Mail and Mail.ru email addresses allowed.</p>
          </div>
          <div class="form-group">
            <label for="inputPassword">Password</label>
            {{if .showAlertPassword}}
              <div class="alert alert-danger" role="alert">{{.alertMessage}}</div>
            {{end}}
            <div class="alert alert-danger" style="display: none;" id="passwordErrMessage"></div>
            <input type="password" class="form-control" id="inputPassword" placeholder="Your New Password" name="userPassword">
          </div>
          <div class="form-group">
            <input type="password" class="form-control" id="inputRepeatPassword" placeholder="Repeat Password" name="userPasswordRepeat">
            <p class="help-block">Password should have at least 6 characters with letters and digits</p>
          </div>
          <!--<div class="checkbox">
            <label>
              <input type="checkbox">Don't remember me
            </label>
          </div>-->
          <button type="submit" class="btn btn-default">Submit</button>
        </form>
      </div>
    </div>
  </body>
</html>
