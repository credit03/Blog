<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <title>{{config "String" "globaltitle" ""}}</title>
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/style.css" rel="stylesheet">

    <link href="/static/css/clndr.css" rel="stylesheet">
    <link href="/static/css/table-responsive.css" rel="stylesheet">
    {{template "inc/meta.tpl" .}}
</head>
<body class="sticky-header">
<section> {{template "inc/left.tpl" .}}

    <div class="main-content">

        <div class="header-section"><a class="toggle-btn"><i class="glyphicon glyphicon-align-justify"></i></a>
            {{template "inc/user-info.tpl" .}}
        </div>

        <div class="wrapper">
            <div class="row">

                <div class="col-md-12">
                    <div class="panel">
                        <header class="panel-heading">用户管理 ({{.Count}})</header>
                        <div class="btn-group pull-right" style="margin: 10px 50px">
                            <button type="button" class="btn btn-primary dropdown-toggle"
                                    data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                添加用户<span class="caret"></span></button>
                            <ul class="dropdown-menu">
                                <li><a href="#" data-id="{{.Id}}">添加</a></li>
                                <li><a href="#" data-id="{{.Id}}">更多功能</a></li>
                                <li role="separator" class="divider"></li>
                            </ul>
                        </div>

                        <div class="panel-body">

                            <table class="table table-bordered table-striped table-condensed">
                                <thead>
                                <tr>
                                    <th>头像</th>
                                    <th>ID</th>
                                    <th>帐号</th>
                                    <th>昵称</th>
                                    <th>性别</th>
                                    <th>年龄</th>
                                    <th>地址</th>
                                    <th>用户组</th>
                                    <th>操作</th>
                                </tr>
                                </thead>
                                <tbody>
                                {{range .UserArr}}
                                <tr>
                                    <td><img
                                            style="  border: 5px solid #F1F2F7;border-radius: 50%; -webkit-border-radius: 50%; height: 32px; width: 32px;"
                                            src="{{.Pic}}"></td>
                                    <td>{{.Id}}</td>
                                    <td>{{.Name}}</td>
                                    <td>{{.NickName}}</td>
                                    <td>{{if .Sex}}男{{else}}女{{end}}</td>
                                    <td>{{.Age}}</td>
                                    <td>{{.Address}}</td>
                                    <td>{{if .IsAdmin}}管理员{{else}}普通用户{{end}}</td>
                                    <td>
                                        <div class="btn-group">
                                            <button type="button" class="btn btn-primary dropdown-toggle"
                                                    data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                                操作<span class="caret"></span></button>
                                            <ul class="dropdown-menu">
                                                <li><a href="/user/borwseinfo?uid={{.Id}}" data-id="{{.Id}}">查看</a></li>
                                                <li><a href="/myajax/editinfo/{{.Id}}" data-id="{{.Id}}">修改资料</a></li>
                                                <li role="separator" class="divider"></li>
                                            </ul>
                                        </div>
                                    </td>
                                </tr>
                                {{end}}
                                </tbody>

                            </table>


                        </div>


                    </div>
                </div>
            </div>
        </div>
        <!--body wrapper end-->
        <!--footer section start-->
        {{template "inc/notice-dialog.tpl" .}}
        <!--footer section end-->
    </div>
    </div>
    <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
<script src="/static/js/calendar/clndr.js"></script>
<script src="/static/js/calendar/evnt.calendar.init.js"></script>
<script src="/static/js/calendar/moment-2.2.1.js"></script>
<script src="/static/js/underscore-min.js"></script>
<script>
    $(function () {
        $('#noticeModal').on('show.bs.modal', function (e) {
            $('#notice-box').html($(e.relatedTarget).attr('data-content'))
        });
    })

    function Eq(t) {
        if ($(t).attr("selectCurr") == $(t).attr("value")) {
            return false;
        }
        return true;
    }

    function OnDelte(va) {
        var b = window.confirm("是否删除该文章 \"" + va.title + "\"");
        return b;
    }
</script>
</body>
</html>
