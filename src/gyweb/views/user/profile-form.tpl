<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <title>基本资料</title>
    <link href="/static/css/style.css" rel="stylesheet">
    {{template "inc/meta.tpl" .}}
</head>
<body class="sticky-header">
<section> {{template "inc/left.tpl" .}}

    <div class="main-content">

        <div class="header-section"><a class="toggle-btn"><i class="glyphicon glyphicon-align-justify"></i></a>
            {{template "inc/user-info.tpl" .}}
        </div>
        <!-- header section end-->
        <!-- page heading start-->
        <!-- page heading end-->
        <!--body wrapper start-->
        <div class="wrapper">
            <div class="row">
                <div class="col-lg-12">
                    <section class="panel">
                        <header class="panel-heading"> 基本资料</header>
                        <div class="panel-body">
                            <form class="form-horizontal adminex-form" id="Useredprofile-form" method="post"
                                  action="/myajax/updateinfo/{{.Usered.Id}}">
                                {{ .xsrfdata }}
                                <div class="form-group">
                                    <label class="col-sm-2 col-sm-2 control-label">昵称</label>
                                    <div class="col-sm-10">
                                        <input type="text" name="nickname" value="{{.Usered.NickName}}"
                                               class="form-control form-check" placeholder="请填写昵称">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-sm-2 col-sm-2 control-label">性别</label>
                                    <div class="col-sm-10">
                                        <label class="radio-inline">
                                            <input type="radio" name="sex" value="1"
                                                   {{if .Usered.Sex}}checked{{end}}>
                                            男 </label>
                                        <label class="radio-inline">
                                            <input type="radio" name="sex" value="2"
                                                   {{if eq false .Usered.Sex}}checked{{end}}>
                                            女 </label>
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-sm-2 col-sm-2 control-label">年龄</label>
                                    <div class="col-sm-10">
                                        <input type="text" name="age" value="{{.Usered.Age}}"
                                               class="form-control form-check" placeholder="请填写年龄">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-sm-2 col-sm-2 control-label">生日</label>
                                    <div class="col-sm-10">
                                        <input type="text" name="birth" id="default-date-picker" value="{{.Usered.Birth}}"
                                               class="form-control form-check" placeholder="请填写生日">
                                    </div>
                                </div>
                                <div class="form-group">
                                    <label class="col-sm-2 col-sm-2 control-label">地址</label>
                                    <div class="col-sm-10">
                                        <input type="text" name="address" value="{{.Usered.Address}}"
                                               class="form-control form-check"
                                               placeholder="请填写地址">
                                    </div>
                                </div>

                                <div class="form-group">
                                    <label class="col-lg-2 col-sm-2 control-label"></label>
                                    <div class="col-lg-10">
                                        <button type="submit" class="btn btn-primary" onclick="return caheckitem()">提
                                            交
                                        </button>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </section>
                </div>
            </div>
        </div>
        <!--body wrapper end-->
        <!--footer section start-->
        <!--footer section end-->
    </div>
    <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
<script src="/static/js/jquery-ui-1.10.3.min.js"></script>
<script src="/static/js/datepicker-zh-CN.js"></script>
<script>
    $(function () {
        $('#default-date-picker').datepicker('option', $.datepicker.regional['zh-CN']);
        $('#default-date-picker').datepicker({
            dateFormat: 'yy-mm-dd',
            changeMonth: true,
            changeYear: true,
            yearRange: '-60:+0'
        });
    })

    function caheckitem() {
        $arr = $(".form-check")

        for (i = 0; i < $arr.length; i++) {
            if ($($arr[i]).val().length == 0) {
                alert($($arr[i]).attr("placeholder"));
                return false;

            }

        }
        return true;


    }


</script>
</body>
</html>
