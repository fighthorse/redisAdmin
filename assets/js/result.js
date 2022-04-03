function ShowResult(dataRes) {
    if (dataRes === undefined) {
        return
    }
    $("#key_key").val(dataRes.keys);
    $("#key_type").val(dataRes.type);
    $("#key_ttl").val(dataRes.ttl);
    $("#key_value").val("");
    $("#key_page").val(0);

    $("#TableResultHtml").hide();
    $("#aloneKeyShow").show();
    if (dataRes.type === "msg"){
        layer.msg(dataRes.data)
    }else if (dataRes.type === "string") {
        $("#key_value").val(dataRes.value)
    } else if (dataRes.type === "list") {
        $("#total_page").val(dataRes.total);
        let str = '';
        str += '<tr>';
        str += '<td>新增数据</td><td colspan="2" ><textarea  class="form-control"  id="new_list_value" placeholder="输入值"></textarea></td>';
        str += '<td><a onclick="AddList(' + dataRes.keys + ')">新增</a></td>';
        str += '</tr><tr class="success" ><td>队列序号</td><td colspan="3">数据值</td></tr>';

        for (var i in dataRes.list) {
            let cc = dataRes.list[i];
            str += '<tr>';
            str += '<td style="width: 50px;">' + i + '</td><td><input class="form-control" id="list_index_' + cc.index + '" value="' + cc.value + '"></td>';
            str += '<td><a onclick="UpdateList(' + dataRes.keys + ',' + cc.index + ')">修改</a></td>';
            str += '<td><a onclick="DelList(' + dataRes.keys + ',' + cc.index + ')">删除</a></td>';
            str += '</tr>'
        }
        $("#TabResult").html(str);
        $("#key_len").html(dataRes.length)
        $("#TableResultHtml").show()
    } else if (dataRes.type === "hash") {
        $("#total_page").val(dataRes.total);
        let str = '';

        str += '<tr>';
        str += '<td ><textarea  class="form-control"  id="new_hash_key" placeholder="名称"></textarea></td><td colspan="2" ><textarea  class="form-control"  id="new_hash_value" placeholder="输入值"></textarea></td>';
        str += '<td><a onclick="AddHash(' + dataRes.keys + ')">新增</a></td>';
        str += '</tr><tr class="success" ><td>字典key</td><td colspan="3" >字典value</td></tr>';

        for (var i in dataRes.hash) {
            str += '<tr>';
            str += '<td style="width: 100px;">' + i + '</td><td><input class="form-control" id="list_index_' + i + '" value="' + dataRes.hash[i] + '"></td>';
            str += '<td><a onclick="UpdateHash(' + dataRes.keys + ',' + i + ')">修改</a></td>';
            str += '<td><a onclick="DelHash(' + dataRes.keys + ',' + i + ')">删除</a></td>';
            str += '</tr>'
        }

        $("#TabResult").html(str);
        $("#key_len").html(dataRes.length)
        $("#TableResultHtml").show()
    }else if (dataRes.type === "set") {
        $("#total_page").val(dataRes.total);
        let str = '';

        str += '<tr>';
        str += '<td colspan="3" ><textarea  class="form-control"  id="new_set_value" placeholder="输入值 逗号分隔 "></textarea></td>';
        str += '<td><a onclick="AddSet(' + dataRes.keys + ')">新增</a></td>';
        str += '</tr><tr class="success" ><td>#</td><td colspan="2" >集合value</td><td></td></tr>';

        let ii = 0
        for (var i in dataRes.hash) {
            str += '<tr>';
            ii++
            str += '<td style="width: 100px;">' + ii + '</td><td colspan="2" ><input class="form-control" id="list_index_' + i + '" value="' + dataRes.hash[i] + '" disabled></td>';
            str += '<td><a onclick="DelSET(' + dataRes.keys + ')">删除</a></td>';
            str += '</tr>'
        }

        $("#TabResult").html(str);
        $("#key_len").html(dataRes.length)
        $("#TableResultHtml").show()
    }else if (dataRes.type === "zset") {
        $("#total_page").val(dataRes.total);
        let str = '';

        str += '<tr>';
        str += '<td><input class="form-control" id="new_set_score" placeholder="输入分值" ></td><td colspan="2" ><input class="form-control"  id="new_set_value" placeholder="输入值"/></td>';
        str += '<td><a onclick="AddSet(' + dataRes.keys + ')">新增</a></td>';
        str += '</tr><tr class="success" ><td>score</td><td colspan="2" >value</td><td></td></tr>';

        for (var i in dataRes.zset) {
            let item = dataRes.zset[i]
            str += '<tr>';
            str += '<td style="width: 100px;">' + item.score + '</td><td colspan="2" ><input class="form-control" id="list_index_' + i + '" value="' + item.member + '" disabled></td>';
            str += '<td><a onclick="DelZSET(' + dataRes.keys + ','+ item.member +')">删除</a></td>';
            str += '</tr>'
        }

        $("#TabResult").html(str);
        $("#key_len").html(dataRes.length)
        $("#TableResultHtml").show()
    }else if (dataRes.type === "bitmap") {
        console.log(dataRes)
    }else{
        console.log(dataRes)
    }
}

function AddList(keys) {

}

function DelList(keys, index) {

}

function UpdateList(keys, index) {

}

function AddHash(keys) {

}

function DelHash(keys, index) {

}

function UpdateHash(keys, index) {

}

function AddSet(keys) {

}

function DelSET(keys) {

}

function DelZSET(keys, value) {

}
function switchPage() {
    var data = {
        "client": $("#SelectDB").val(),
        "db": $("#SelectDBIndex").val(),
        "key": $("#key_key").val(),
        "page": $("#key_page").val(),
        "type": $("#key_type").val(),
        "token": GetLocalToken(),
    };
    $.ajax({
        type: "POST",
        url: '/redis/getKey',
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
                layer.msg(response.message);
            }
        }
    });
}