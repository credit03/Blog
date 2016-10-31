<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <title>{{config "String" "globaltitle" ""}}</title>
    <link href="/static/css/style.css" rel="stylesheet">
    <link href="/static/plugins/flexslider/flexslider.css" rel="stylesheet">
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


                <div class="col-md-4">
                    <div class="row">

                        <div class="col-md-12">
                            <div class="panel">
                                <div class="panel-body">

                                    <div class="profile-pic text-center"><img alt="" src="{{.User.Pic}}"></div>
                                    <ul class="p-info">
                                        <li>
                                            <div class="title">文章分类</div>
                                            <div class="desk">{{.OpType}}</div>
                                            <div class="title">文章总数</div>
                                            <div class="desk">{{.Count}}</div>
                                        </li>

                                    </ul>
                                </div>
                            </div>
                        </div>


                    </div>
                </div>

                <div class="col-md-8">
                    <div class="panel">
                        <header class="panel-heading">文章列表({{.Count}})</header>
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

                                        {{if eq $v.ContentType 1}}
                                        {{template "item1" .}}
                                        {{else}}
                                        {{if eq $v.ContentType 2}}
                                        {{template "item2" .}}
                                        {{else}}
                                        {{if eq $v.ContentType 3}}
                                        {{template "item3" .}}
                                        {{end}}
                                        {{end}}
                                        {{end}}
                                        <p class="text-muted">{{$v.Describe}}</p>

                                        <p class="pull-right text-muted">
                                            <i class="fa fa-eye"></i>
                                            {{$v.BrowseCount}}&nbsp;&nbsp;&nbsp;
                                            <!-- <i class="glyphicon glyphicon-trash"></i>
                                             <a href="/myajax/category/delete?id={{.Id}}"
                                                title="{{.Title}}"
                                                onclick="return OnDelte(this) ">删除</a>&nbsp;&nbsp;&nbsp;
                                             <i class="glyphicon glyphicon-pencil"></i><a
                                                 href="/myajax/category/edit?id={{.Id}}">
                                             修改</a>&nbsp;&nbsp;&nbsp;-->

                                            {{getDateMH $v.Created}}&nbsp;&nbsp;&nbsp;

                                        </p>

                                        <div class="btn-group pull-right">
                                            <button type="button" class="btn btn-primary dropdown-toggle"
                                                    data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                                操作<span class="caret"></span></button>
                                            <ul class="dropdown-menu">
                                                <li role="separator" class="divider"></li>
                                                <li><a href="/myajax/category/delete?id={{.Id}}"
                                                       title="{{.Title}}"
                                                       onclick="return OnDelte(this) ">删除</a></li>
                                                <li><a
                                                        href="/myajax/category/edit?id={{.Id}}">
                                                    修改</a></li>
                                            </ul>
                                        </div>
                                    </div>
                                </li>
                                {{end}}
                            </ul>

                        </div>


                        <div class="text-center">
                            <div class="row">
                                {{if gt .PageLimit.IndexCount 1}}
                                <nav>
                                    <ul class="pagination">
                                        {{if ge .PageLimit.CurrentIndex 2}}
                                        <li>
                                            <a href='{{urlfor "CategoryController.PageIndex"}}?form=gategorymanage&op=Previous&index={{.PageLimit.CurrentIndex}}'><span
                                                    aria-hidden="true">&laquo;</span><span
                                                    class="sr-only">Previous</span></a>
                                        </li>
                                        {{end}}
                                        {{range $i,$v:=.PageLimit.IndexArr}}
                                        <li {{if eq $v $.PageLimit.CurrentIndex}}class="active" {{end}}><a
                                                href='{{urlfor "CategoryController.PageIndex"}}?form=gategorymanage&op=Index&index={{$v}}'
                                                selectCurr="{{$.PageLimit.CurrentIndex}}" value="{{$v}}"
                                                onclick="return Eq(this)">{{$v}}</a></li>
                                        {{end}}
                                        {{if lt .PageLimit.CurrentIndex .PageLimit.IndexCount}}
                                        <li>
                                            <a href='{{urlfor "CategoryController.PageIndex"}}?form=gategorymanage&op=Next&index={{.PageLimit.CurrentIndex}}'><span
                                                    aria-hidden="true">&raquo;</span><span
                                                    class="sr-only">Next</span></a>
                                        </li>
                                        {{end}}
                                    </ul>
                                </nav>
                                {{else}}

                                <h1>全部文章加载完毕</h1>
                                {{end}}
                            </div>
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

<script src="/static/plugins/flexslider/jquery.flexslider-min.js"></script>
<script src="/static/plugins/flexslider/flexslider.config.js"></script>

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

{{define "item1"}}
<article>
    <div class="row">
        <div class="col-lg-10">
            <div class="post-image">
                {{range .Images}}
                <img src="/{{.Url}}" alt="" style="max-width: 480px; height: 320px"/>
                {{end}}
            </div>
        </div>
    </div>
</article>
{{end}}

{{define "item2"}}


<div class="row">
    <div class="col-lg-10">
        <div>
            <!-- start flexslider -->
            <div class="postslider flexslider">
                <ul class="slides">
                    {{range .Images}}
                    <li>
                        <img src="/{{.Url}}" style="height: 320px" alt=""/>
                    </li>
                    {{end}}
                </ul>
            </div>
        </div>
        <!-- end flexslider -->
    </div>
</div>


{{end}}


{{define "item3"}}
<article class="panel">
    <div class="row">
        <div class="col-lg-10">
            <div class="post-video">
                <!-- 视频-->
                <iframe width="560" height="315"
                        src="{{.VideoUrl}}"
                        frameborder="1" allowfullscreen></iframe>
            </div>
        </div>
    </div>
</article>

{{end}}

