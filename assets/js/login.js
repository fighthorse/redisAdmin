function CheckLoginStatus() {
    token = GetLocalToken();
    if (token == null || token == "") {
        NeedLogin();
        return
    }
    // check token
    CheckToken(token)
}

function Login() {
    var data = {
        "name": $("#UserName").val(),
        "pwd": $("#Password").val(),
    };
    $.ajax({
        type: "POST",
        url: '/login/submit',
        data: data,
        success: function (response) {
            if (response.code !== 0) {
                layer.msg(response.message);
                return
            }
            layer.msg(response.message);
            var data = response.data;
            SetLocalToken(data.token);
            HasLogin();
            InitDBSelect();
        }
    });
}

function CheckToken(token) {
    var data = {
        "token": token,
    };
    $.ajax({
        type: "POST",
        url: '/login/check',
        data: data,
        success: function (response) {
            if (response.code !== 0) {
                layer.msg(response.message);
                ClearLocalToken();
                NeedLogin();
                return
            }
            HasLogin()
        }
    });
}

function NeedLogin() {
    $("#UnLoginForm").show();
    $("#workForm").hide()
}

function HasLogin() {
    $("#UnLoginForm").hide();
    $("#workForm").show()
}

function GetLocalToken() {
    if (typeof (Storage) !== "undefined") {
        // 针对 localStorage/sessionStorage 的代码
        return getLocalStorage("redis_token", "token_expires")
    } else {
        // 抱歉！不支持 Web Storage ..
        return getCookie("redis_token")
    }
}

function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i].trim();
        if (c.indexOf(name) == 0) return c.substring(name.length, c.length);
    }
    return "";
}

function setCookie(cname, cvalue, exdays) {
    let d = new Date();
    d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
    let expires = "expires=" + d.toGMTString();
    document.cookie = cname + "=" + cvalue + "; " + expires;
}

function delCookie(cname) {
    document.cookie = ""
}

function getLocalStorage(cname, expiresname) {
    let token = localStorage.getItem(cname);
    if (token == null) {
        return token
    }
    let expires = localStorage.getItem(expiresname);
    if (expires == null) {
        return null
    }
    let d = new Date();
    let dd = GMTToStr(d.toGMTString());
    let ss = GMTToStr(expires);
    // console.log(dd, ss)
    if (isFirstBig(dd, ss)) {
        ClearLocalToken();
        return ""
    }
    return token
}

function isFirstBig(date1, date2) {
    var oDate1 = new Date(date1);
    var oDate2 = new Date(date2);
    return oDate1.getTime() > oDate2.getTime();
}

function delLocalStorage(cname, expiresname) {
    localStorage.removeItem(cname);
    localStorage.removeItem(expiresname)
}

function setLocalStorage(cname, cvalue, expiresname, exdays) {
    let d = new Date();
    d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
    localStorage.setItem(expiresname, d.toGMTString());
    return localStorage.setItem(cname, cvalue)
}


function GMTToStr(time) {
    let date = new Date(time);
    let Str = date.getFullYear() + '-';
    let month = date.getMonth() + 1;
    if (month < 10) {
        Str += '0' + month + '-';
    } else {
        Str += +month + '-';
    }
    let dt = date.getDate();
    if (dt < 10) {
        Str += '0' + dt + ' ';
    } else {
        Str += +dt + ' ';
    }
    let hours = date.getHours();
    if (hours < 10) {
        Str += '0' + hours + ':'
    } else {
        Str += hours + ':'
    }

    let minutes = date.getMinutes();
    if (minutes < 10) {
        Str += '0' + minutes + ':'
    } else {
        Str += minutes + ':'
    }

    let seconds = date.getSeconds();
    if (seconds < 10) {
        Str += '0' + seconds
    } else {
        Str += seconds
    }
    return Str
}


function SetLocalToken(token) {
    if (typeof (Storage) !== "undefined") {
        // 针对 localStorage/sessionStorage 的代码
        // 存储
        // localStorage.setItem("lastname", "Gates");
        // 获取
        // localStorage.getItem("lastname")
        // 删除
        //localStorage.removeItem("lastname")
        return setLocalStorage("redis_token", token, "token_expires", 1)
    } else {
        // 抱歉！不支持 Web Storage ..
        // set kvx形式
        //document.cookie="username=John Doe";
        // get
        // var x = document.cookie;
        return setCookie("redis_token", token, 1)
    }
}

function ClearLocalToken() {
    if (typeof (Storage) !== "undefined") {
        // 针对 localStorage/sessionStorage 的代码
        // 存储
        // localStorage.setItem("lastname", "Gates");
        // 获取
        // localStorage.getItem("lastname")
        // 删除
        //localStorage.removeItem("lastname")
        return delLocalStorage("redis_token", "token_expires")
    } else {
        // 抱歉！不支持 Web Storage ..
        // set kvx形式
        //document.cookie="username=John Doe";
        // get
        // var x = document.cookie;
        return delCookie("redis_token")
    }
}