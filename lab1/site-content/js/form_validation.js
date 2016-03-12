window.onload = addListeners;

function addListeners()
{
    $("inputNickname").addEventListener("change", checkNickname, false);
    $("inputEmail").addEventListener("change", checkEmail, false);
    $("inputPassword").addEventListener("change", checkPassword, false);
    $("inputRepeatPassword").addEventListener("change", checkRepeatPassword, false);

    $("regForm").addEventListener('submit', function(evt) {
        if (!(checkNickname() && checkEmail() && checkPassword() && checkRepeatPassword()))
        {
            evt.preventDefault();
        }
    })

    function checkNickname()
    {
        var nickname = $("inputNickname").value;
        var nicknamePattern = /^[a-zA-Z0-9_]{3,}$/i;
        var message;

        if (nickname.length < 3)
        {
            message = "nickname is too short, enter at least 3 symbols";
        }
        else if (nicknamePattern.test(nickname) == false)
        {
            message = "nickname contains not allowed symbols";
        }

        proccessValidationResult("nicknameErrMessage", message);
        return (message == undefined);
    }

    function checkEmail()
    {
      var email = $("inputEmail").value;
      var mailboxNamePattern = /^[a-zA-Z0-9]{1,}$/i;
      var domainNamePattern = /^(gmail.com|yandex.ru|mail.ru)$/i;
      var message;

      var emailParts = email.split('@');
      if (emailParts.length != 2)
      {
          message = "email should contain one @";
      }
      else
      {
          if (mailboxNamePattern.test(emailParts[0]) == false)
          {
              message = "email contains not allowed symbols";
          }
          else if (domainNamePattern.test(emailParts[1]) == false)
          {
              message = "domain is not from allowed list";
          }
      }

      proccessValidationResult("emailErrMessage", message);
      return (message == undefined);
    }

    function checkPassword()
    {
        var password = $("inputPassword").value;
        var lettersPattern = /[a-zA-Z]/i;
        var digitsPattern = /[0-9]/i;
        var message;

        if (password.length < 6)
        {
            message = "password is too short, enter at least 6 symbols";
        }
        else if (lettersPattern.test(password) == false)
        {
            message = "password should contain at least 1 latin letter";
        }
        else if (digitsPattern.test(password) == false)
        {
            message = "password should contain at least 1 digit";
        }

        proccessValidationResult("passwordErrMessage", message);
        return (message == undefined);
    }

    function checkRepeatPassword()
    {
        var message;

        if ($("inputPassword").value != $("inputRepeatPassword").value)
        {
            message = "passwords are not matching";
        }

        proccessValidationResult("passwordErrMessage", message);
        return (message == undefined);
    }

    function proccessValidationResult(msgBlockId, message)
    {
        var errMessageBlock = $(msgBlockId);
        if (message != undefined)
        {
            errMessageBlock.innerHTML = message;
            errMessageBlock.style.display = 'block';
        }
        else
        {
            errMessageBlock.style.display = 'none';
        }
    }
}

function $(id)
{
    return document.getElementById(id);
}
