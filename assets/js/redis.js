// init check
$(document).ready(function () {
    CheckLoginStatus();
    InitDBSelect()
});

function Search() {
    InitDBIndexSelect()
}

function InitDBSelect() {
    let token = GetLocalToken();
    if (token == null || token === "") {
        return
    }
    var data = {
        "token": GetLocalToken(),
    };
    $.ajax({
        type: "POST",
        url: '/redis/init',
        data: data,
        success: function (response) {
            if (response.code === -126) {
                NeedLogin();
                return
            }
            if (response.code !== 0) {
                layer.msg(response.message);
                return
            }
            let dataRes = response.data;
            let str = "";
            for (let i in dataRes) {
                str += "<option value='" + i + "'>" + i + ":" + dataRes[i].addr + ")</option>";
            }
            $("#SelectDB").html(str)
        }
    });
}

function InitDBIndexSelect() {
    let token = GetLocalToken();
    if (token == null || token === "") {
        return
    }
    var data = {
        "client": $("#SelectDB").val(),
        "db": $("#SelectDBIndex").val(),
        "key": $("#SelectKey").val(),
        "token": GetLocalToken(),
    };
    $.ajax({
        type: "POST",
        url: '/redis/search',
        data: data,
        success: function (response) {
            if (response.code === -126) {
                NeedLogin();
                return
            }
            if (response.code !== 0) {
                layer.msg(response.message);
                return
            }
            var data = response.data;
            console.log(data);
            let str = addItem(data); //init
            $("#KeysListHtml").html(str)
        }
    });
}

function ClearInput() {

}

function NextUlShow(event) {
    stopBubble(event);
    let show = $(event).data("data-show");
    let child = $(event).attr("data-child");
    if (show === "1") {
        $(event).data("data-show", "0");
        $(event).children("ul").hide()
    } else {
        $(event).data("data-show", "1");
        let nums = $(event).attr("data-nums");
        if (nums === "1") {
            let key = $(event).attr("data-key");
            let level = $(event).attr("data-level");
            SearchNowKey(event, key, level);
            return
        }
        var index = layer.load(2, {time: 10 * 1000}); //又换了种风格，并且设定最长等待10秒
        KeyIndexSelect(event);

        //关闭
        layer.close(index);
        $(event).children("ul").show()
    }
}

function stopBubble(e) {
//如果提供了事件对象，则这是一个非IE浏览器
    if (e && e.stopPropagation)
    //因此它支持W3C的stopPropagation()方法
        e.stopPropagation();
    else
    //否则，我们需要使用IE的方式来取消事件冒泡
        window.event.cancelBubble = true;
}

function KeyIndexSelect(event) {
    var data = {
        "client": $("#SelectDB").val(),
        "db": $("#SelectDBIndex").val(),
        "key": $(event).attr("data-key"),
        "level": $(event).attr("data-level"),
        "token": GetLocalToken(),
    };
    $.ajax({
        type: "POST",
        url: '/redis/search',
        data: data,
        success: function (response) {
            if (response.code === -126) {
                NeedLogin();
                return
            }
            if (response.code !== 0) {
                layer.msg(response.message);
                return
            }
            var data = response.data;
            console.log(data);
            let str = addItem(data);
            $(event).children("ul").remove();
            $(event).append(str);
            $("#aloneKeyShow").hide()
        }
    });
}

function Del() {

}

function Add() {
    let val = $("#addCfgHtml").data("data-show");
    if (val === 1) {
        $("#addCfgHtml").data("data-show", 0);
        $("#addCfgHtml").hide()
    } else {
        $("#addCfgHtml").data("data-show", 1);
        $("#addCfgHtml").show()
    }
}

function AddCfg() {
    var data = {
        "name": $("#AddName").val(),
        "addr": $("#AddAddr").val(),
        "pwd": $("#AddPwd").val(),
        "token": GetLocalToken(),
    };
    $.ajax({
        type: "POST",
        url: '/redis/addCfg',
        data: data,
        success: function (response) {
            if (response.code === -126) {
                NeedLogin();
                return
            }
            if (response.code !== 0) {
                layer.msg(response.message);
                return
            }
            layer.msg(response.message);
            let str = "<option value='" + data.name + "'>" + data.name + ":" + data.addr + ")</option>";
            $("#SelectDB").append(str)
        }
    });
}

function SearchNowKey(event, key, level) {
    var data = {
        "client": $("#SelectDB").val(),
        "db": $("#SelectDBIndex").val(),
        "key": key,
        "level": level,
        "token": GetLocalToken(),
    };
    $.ajax({
        type: "POST",
        url: '/redis/searchNowKey',
        data: data,
        success: function (response) {
            if (response.code === -126) {
                NeedLogin();
                return
            }
            if (response.code !== 0) {
                layer.msg(response.message);
                return
            }
            var dataRes = response.data;
            if (dataRes.type != "") {
                ShowResult(dataRes)
            } else {
                let str = addItem(dataRes);
                $(event).children("ul").remove();
                $(event).append(str)
            }
        }
    });
}

function HandleType() {
    var data = {
        "client": $("#SelectDB").val(),
        "db": $("#SelectDBIndex").val(),
        "type": $("#Handletype").val(),
        "key": $("#HandleKey").val(),
        "value": $("#HandleValue").val(),
        "ttl": $("#HandleTTl").val(),
        "token": GetLocalToken(),
    };
    $.ajax({
        type: "POST",
        url: '/redis/handle',
        data: data,
        success: function (response) {
            if (response.code === -126) {
                NeedLogin();
                return
            }
            if (response.code !== 0) {
                layer.msg(response.message);
                return
            }
            layer.msg(response.message);
            var dataRes = response.data;
            console.log(dataRes);

        }
    });
}

function SearchKey() {
    var data = {
        "client": $("#SelectDB").val(),
        "db": $("#SelectDBIndex").val(),
        "type": $("#Handletype").val(),
        "key": $("#HandleKey").val(),
        "value": $("#HandleValue").val(),
        "ttl": $("#HandleTTl").val(),
        "token": GetLocalToken(),
    };
    $.ajax({
        type: "POST",
        url: '/redis/searchKey',
        data: data,
        success: function (response) {
            if (response.code === -126) {
                NeedLogin();
                return
            }
            if (response.code !== 0) {
                layer.msg(response.message);
                return
            }
            layer.msg(response.message);
            var dataRes = response.data;
            console.log(dataRes);

        }
    });
}

var str = '';

function formatJson(msg) {
    var rep = "~";
    var jsonStr = JSON.stringify(msg, null, rep);
    var str = "";
    for (var i = 0; i < jsonStr.length; i++) {
        var text2 = jsonStr.charAt(i);
        if (i > 1) {
            var text = jsonStr.charAt(i - 1);
            if (rep != text && rep == text2) {
                str += "<br/>"
            }
        }
        str += text2;
    }
    jsonStr = "";
    for (var i = 0; i < str.length; i++) {
        var text = str.charAt(i);
        if (rep == text)
            jsonStr += "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;";
        else {
            jsonStr += text;
        }
        if (i == str.length - 2)
            jsonStr += "<br/>"
    }
    return jsonStr;
}

function addItem(item) {
    let list = item.keys;
    let child = item.child;
    let d = new Date();
    d.setTime(d.getTime() + 300);
    let time = d.toGMTString();
    let str = "<ul class=\"list-group\" >";
    for (let i in list) {
        str += "<li class=\"list-group-item\" onclick=\"NextUlShow(this)\" data-nums='" + list[i] + "' data-child='" + child + "' data-time='" + time + "' data-show=\"0\" data-key='" + i + "' data-level=" + item.level + " >";
        let total = list[i];
        if (child === true) {
            str += "<span class=\"badge\">" + total + "</span>"
        }
        str += i;
        str += "</li>"
    }
    str += "</ul>";
    return str
}