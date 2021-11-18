$(function(){
    $("#BT_SendMsg").click(function(){

        alert("BT_SendMsg")
        if ($("#SendMsg").val().length == 0) {
            alert("请输入发送报文后再进行发送")
            return
        }
        var ip = $("#ServerIp").val();
        var port = $("#ServerPort").val();
        var msg = JSON.stringify(JSON.parse($("#SendMsg").val()));
        var url = "/Msg/SendMsg"

        //设置请求参数
        var param = {"ip":ip, "port":port, "msg":msg };
        var jsonObj = JSON.parse(msg)
        $("#SendMsg").val(JSON.stringify(jsonObj, null, 4));
        //发送Ajax请求
        $.post(url,param,function(res){
            if (res != "err"){
                $("#ResMsg").val(res);
            }else{
                alert(res);
            }
        });
    });

    $("#GetVersion").click(function(){
        var url = "/shell/GetVersion"
        //发送Ajax请求
        $.post(url,null,function(res){
            if (res == "ok"){
                alert("ok!");
            }else{
                $("#shell_Resp_msg").val(res);
            }
        });
    });

    $("#GetRequest").click(function(){
        var url = "/shell/GetRequest"
        //发送Ajax请求
        $.post(url,null,function(res){
            if (res == "ok"){
            alert("ok!");
            }else{
                $("#shell_Resp_msg").val(res);
            }
        });
    });

    $("#GetClientInfo").click(function(){
        var url = "/shell/GetClientInfo"
        //发送Ajax请求
            $.post(url,null,function(res){
            if (res == "ok"){
                alert("ok!");
            }else{
                $("#shell_Resp_msg").val(msg);
            }
        });
    });

    $("#GetHealthz").click(function(){
        var url = "/shell/GetHealthz"
        //发送Ajax请求
        $.post(url,null,function(res){
            if (res == "ok"){
                alert("ok!");
            }else{
                $("#shell_Resp_msg").val(res);
            }
        });
    });
});