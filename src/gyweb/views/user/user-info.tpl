<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <title>用户资料</title>
    <link href="/static/css/style.css" rel="stylesheet">
    {{template "inc/meta.tpl" .}}
</head>
<body>
<!-- left side start-->
<div class="left-side sticky-left-side">
    <!--logo and iconic logo start-->
    <div class="logo"><a href="/"><img src="/static/img/logo_icon.png" alt=""></a></div>
    <div class="logo-icon text-center"><a href="/"><img src="/static/img/logo_icon.png" style="width:40px;" alt=""></a>
    </div>
    <!--logo and iconic logo end-->
    <div class="left-side-inner">
        <!--sidebar nav start-->
        <ul class="nav nav-pills nav-stacked custom-nav js-left-nav">
            <li><a href="/myajax/usercenter"><i class="glyphicon glyphicon-home"></i> <span> 我的主页</span></a></li>
            <li><a onclick="back()"><i class="glyphicon glyphicon-share-alt"></i> <span> 返回上一级</span></a></li>
            <li><a href="/"><i class="glyphicon glyphicon-chevron-left"></i> <span> 返回首页</span></a></li>
        </ul>

        <!--sidebar nav end-->
    </div>
</div>
<div class="main-content">
    <div class="header-section"><a class="toggle-btn"><i class="glyphicon glyphicon-align-justify"></i></a>
    </div>

    <div class="wrapper">
        <div class="row">
            <div class="col-md-4">
                <div class="row">

                    <div class="col-md-12">
                        <div class="panel">
                            <div class="panel-body">

                                <div class="profile-pic text-center"><img alt=""
                                                                          src="{{.User.Pic}}">
                                </div>
                                <ul class="p-info">
                                    <li>
                                        <div class="title">帐号</div>
                                        <div class="desk">{{.User.Name}}</div>
                                    </li>
                                    <li>
                                        <div class="title">昵称</div>
                                        <div class="desk">{{ .User.NickName}}</div>
                                    </li>
                                    <li>
                                        <div class="title">性别</div>
                                        <div class="desk">{{if .User.Sex}}男{{else}}女{{end}}
                                        </div>
                                    </li>
                                    <li>
                                        <div class="title">年龄</div>
                                        <div class="desk">{{.User.Age}}</div>
                                    </li>
                                    <li>
                                        <div class="title">生日</div>
                                        <div class="desk">{{.User.Birth}}</div>
                                    </li>
                                    <li>
                                        <div class="title">地址</div>
                                        <div class="desk">{{.User.Address}}</div>
                                    </li>
                                    {{if .User.IsAdmin}}
                                    <li>
                                        <div class="title">用户组</div>
                                        <div class="desk">管理员</div>
                                    </li>
                                    {{else}}
                                    <li>
                                        <div class="title">用户组</div>
                                        <div class="desk">普通用户</div>
                                    </li>
                                    {{end}}
                                </ul>
                            </div>
                        </div>
                    </div>


                </div>
            </div>

            <div class="col-md-8">
                <div class="panel">
                    <header class="panel-heading">最近10篇文章</header>
                    <div class="panel-body">
                        <ul class="activity-list">
                            {{range $k,$v := .PostArr}}
                            <li style="   display: inline-block; width: 100%; margin-bottom: 30px;border-bottom: 1px solid #eff0f4;">
                                <div class="avatar"><img src="{{$v.Author.Pic}}" alt=""></div>
                                <div class="activity-desk">
                                    <h5><a href=""></a> <span><b
                                            style="color: #9b4a1b;font-size: 18px">{{$v.Author.NickName}}</b> &nbsp;&nbsp;<a
                                            href="/browse_category?id={{$v.Id}}"
                                            style="color:#2a323f">{{$v.Title}}</a></span>
                                    </h5>
                                    <p class="text-muted">{{$v.Describe}}</p>

                                    <p class="pull-right text-muted">
                                        <i class="fa fa-eye"></i>
                                        {{$v.BrowseCount}}&nbsp;&nbsp;&nbsp;
                                        {{getDateMH $v.Created}}&nbsp;&nbsp;&nbsp;

                                    </p>
                                </div>
                            </li>
                            {{end}}
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    <!--footer section end-->
</div>
</div>
<!-- main content end-->
{{template "inc/foot.tpl" .}}

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
    function back() {
        history.go(-1); //后退1页
    }
    function next() {
        history.go(+1); //前进1页
    }
    function refresh() {
        history.go(-0) //刷新
    }
</script>
</body>
</html>
