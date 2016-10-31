<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>修改密码</title>
    {{template "inc/meta.tpl" .}}
</head>
<body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
    <!-- main content start-->
    <div class="main-content">
        <!-- header section start-->
        <div class="header-section">
            <!--toggle button start-->
            <a class="toggle-btn"><i class="fa fa-bars"></i></a> {{template "inc/user-info.tpl" .}}
        </div>
        <!-- header section end-->
        <!-- page heading start-->
        <!-- page heading end-->
        <!--body wrapper start-->
        <div class="wrapper">
            <div class="row">
                <div class="col-lg-12">
                    <section class="panel">
                        <header class="panel-heading"> 修改密码</header>
                        <div class="panel-body">
                            <form class="form-horizontal adminex-form" id="userprofilepwd-form">
                                <div class="form-group">
                                    <label class="col-sm-2 col-sm-2 control-label">旧密码</label>
                                    <div class="col-sm-10">
                                        <input type="password" name="oldpwd" id="oldpwd" class="form-control"
                                               placeholder="请填旧密码">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-sm-2 col-sm-2 control-label">新密码</label>
                                    <div class="col-sm-10">
                                        <input type="password" name="newpwd" id="newpwd" class="form-control"
                                               placeholder="请填写新密码（字母与数字组合）">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-sm-2 col-sm-2 control-label">确认密码</label>
                                    <div class="col-sm-10">
                                        <input type="password" name="confpwd" id="confpwd" class="form-control"
                                               placeholder="请填写确认（字母与数字组合）">
                                    </div>
                                </div>

                                <div class="form-group">
                                    <label class="col-lg-2 col-sm-2 control-label"></label>
                                    <div class="col-lg-10">
                                        <input type="hidden" name="id" id="uid" value="{{.User.Id}}">
                                        <button type="button" class="btn btn-primary" onclick="return SubPwd()">提 交
                                        </button>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </section>
                </div>
            </div>
        </div>

        <!--footer section end-->
    </div>
    <!-- main content end-->
</section>

<!-- 按钮触发模态框 -->
<!-- 模态框（Modal） -->
<div class="modal fade" id="myLoad" tabindex="-1" role="dialog" aria-labelledby="myLoadLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                <h4 class="modal-title" id="myLoadLabel">设置结果</h4>
            </div>
            <div class="modal-body"><img src="/static/img/cbp-loading.gif"><span id="myLoadmsg">正在设置</span></div>
            <div class="modal-footer">
                <button type="button" class="btn btn-primary" onclick="OpenUserCenter()">返回个人中心</button>
                <button type="button" class="btn btn-primary" data-dismiss="modal">关闭</button>
            </div>
        </div><!-- /.modal-content -->
    </div><!-- /.modal -->
</div>
<script type="text/javascript">
    function SubPwd() {


        $arr = $(".form-control")
        for (i = 0; i < $arr.length; i++) {
            if ($($arr[i]).val().length == 0) {
                alert($($arr[i]).attr("placeholder"));
                return false;
            }

        }

        if ($("#newpwd").val() != $("#confpwd").val()) {
            alert("两者密码不一致")
            return false
        }

        var patt1 = new RegExp("^[0-9]*$");

        if (patt1.test($("#newpwd").val())) {
            alert("密码不能纯数字哦。。。")
            return false
        }

        var url = "/myajax/updateuserpwd";
        var args = {
            "uid": $('#uid').val(),
            "oldpwd": $('#oldpwd').val(),
            "newpwd": $('#newpwd').val(),
            "confpwd": $("#confpwd").val()
        }

        $.getJSON(url, args, function (data) {
            $('#myLoad').modal('show');
            var msg = data.Msg;
            var ok = data.Ok;
            if (ok == 200) {
                $('#myLoadLabel').text("设置结果：" + ok);
                $('#myLoadmsg').html("<h1>" + msg + "</h1>" +
                        "<b>3秒后自动跳转到个人中心</b>");
                setTimeout("OpenUserCenter()", 3000)
            } else {
                $('#myLoadLabel').text("设置结果：" + ok);
                $('#myLoadmsg').html("<h1>" + msg + "</h1>" +
                        "<b>3秒后自动关闭</b>");
                setTimeout(" $('#myLoad').modal('hide')", 3000)
            }

        });
        return false
    }

    function OpenUserCenter() {
        $('#myLoad').modal('hide');
        window.open("/myajax/usercenter", "_self");
    }

</script>
{{template "inc/foot.tpl" .}}
</body>
</html>
