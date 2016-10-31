<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <title>{{config "String" "globaltitle" ""}}</title>
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/style.css" rel="stylesheet">
    <link href="/static/css/cubeportfolio.min.css" rel="stylesheet">
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
                <div class="col-md-4">
                    <div class="row">

                        <div class="col-md-12">
                            <div class="panel">
                                <div class="panel-body">

                                    <div class="profile-pic text-center"><img alt="" src="{{.User.Pic}}"></div>
                                    <ul class="p-info">
                                        <li>
                                            <div class="title">帐号</div>
                                            <div class="desk">{{ .User.Name}}</div>
                                        </li>
                                        <li>
                                            <div class="title">昵称</div>
                                            <div class="desk">{{ .User.NickName}}</div>
                                        </li>
                                        <li>
                                            <div class="title">性别</div>
                                            <div class="desk">{{if .User.Sex}}男{{else}}女{{end}}</div>
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

                        <div class="col-md-12">
                            <div class="panel">
                                <div class="panel-body">
                                    <div class="calendar-block ">
                                        <div class="cal1"></div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-md-8">
                    <!--

                                        <div class="row">
                                            <div class="col-md-12">
                                                <div class="panel">
                                                    <header class="panel-heading"> 热门文章 <span class="pull-right"> <a
                                                            href="/myajax/gategory_manage/{{.User.Id}}">更多</a></span></header>
                                                    <div class="panel-body">
                                                        {{if gt (.PostArr|len) 0}}
                                                        <ul class="activity-list">
                                                            {{range $k,$v := .PostArr}}
                                                            <li>
                                                                <div class="avatar"><img src="{{$v.Author.Pic}}" alt=""></div>
                                                                <div class="activity-desk">
                                                                    <h5><a href=""></a> <span> <b
                                                                            style="color: #9b4a1b">{{$v.Author.Name}}</b> &nbsp;&nbsp;<a
                                                                            href="/browse_category?id={{$v.Id}}"
                                                                            style="color:#2a323f">{{$v.Title}}</a></span>
                                                                    </h5>
                                                                    <p class="text-muted">{{$v.Describe}}</p>
                                                                    <p class="pull-right text-muted">
                                                                        <i class="fa fa-eye"></i>
                                                                        {{$v.BrowseCount}}&nbsp;&nbsp;&nbsp;

                                                                        {{getDateMH $v.Created}}&nbsp;&nbsp;&nbsp;</p>
                                                                </div>
                                                            </li>
                                                            {{end}}
                                                        </ul>
                                                        {{else}}
                                                        <h1>您暂没有热门文章</h1>
                                                        {{end}}
                                                    </div>
                                                </div>
                                            </div>
                    -->


                    <div class="row">
                        <div class="col-md-12">
                            <div class="panel">
                                <header class="panel-heading">个人统计</header>
                                <div class="panel-body">

                                    <table class="table table-bordered table-striped table-condensed">
                                        <thead>
                                        <tr>
                                            <th>编号</th>
                                            <th>名称</th>
                                            <th>总浏览次数</th>
                                            <th>截止时间</th>
                                        </tr>
                                        </thead>
                                        <tbody>
                                        <tr>
                                            <td>-</td>
                                            <td>
                                                <a href="/myajax/gategory_manage/{{.User.Id}}?op=-1&type=SectionType">发表文章总数：{{.CateCount}}篇</a>
                                            </td>
                                            <td>{{.BrowseCount}}</td>
                                            <td> {{getCurrentTime}}</td>

                                        </tr>
                                        <tr>
                                            <td colspan="4">Top5文章</td>
                                        </tr>
                                        <tr>
                                            <td>编号</td>
                                            <td>文章名称</td>
                                            <td>总浏览次数</td>
                                            <td>最后发表时间</td>
                                        </tr>

                                        {{if gt (.PostArr|len) 0}}

                                        {{range $k,$v := .PostArr}}
                                        <tr>
                                            <td>Top{{getTopNo $k}}</td>
                                            <td><a href="/browse_category?id={{$v.Id}}">{{$v.Title}}</a></td>
                                            <td><i class="fa fa-eye"></i>
                                                {{$v.BrowseCount}}
                                            </td>
                                            <td> {{getDateMH $v.Created}}</td>

                                        </tr>
                                        {{end}}
                                        {{else}}
                                        <td colspan="4">你没有发表过任何文章...</td>
                                        {{end}}
                                        </tbody>

                                    </table>


                                </div>
                            </div>
                        </div>
                    </div>


                    <div class="row">
                        <div class="col-md-12">
                            <div class="panel">
                                <header class="panel-heading">最近图片</header>
                                <div class="panel-body">
                                    <div id="grid-container" class="cbp-l-grid-projects">
                                        <ul>
                                            {{range $k,$t:=getImages .User.Name 0}}
                                            {{range $i,$v:=$t.Images}}
                                            <li class="cbp-item">
                                                <div class="cbp-caption">
                                                    <div class="cbp-caption-defaultWrap">
                                                        <img src="/{{$v.Url}}"
                                                             style="width: inherit;height: inherit"
                                                             alt=""/>
                                                    </div>
                                                    <div class="cbp-caption-activeWrap">
                                                        <div class="cbp-l-caption-alignCenter">
                                                            <div class="cbp-l-caption-body">
                                                                <a href="/{{$v.Url}}"
                                                                   class="cbp-lightbox cbp-l-caption-buttonRight"
                                                                   data-title="{{$t.Title}}<br>by {{$t.Author.Name}}">查看大图</a>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                                <div class="cbp-l-grid-projects-title">
                                                    {{$t.Title}}--{{$i}}
                                                </div>
                                            </li>
                                            {{end}}
                                            {{end}}
                                        </ul>
                                    </div>

                                </div>
                            </div>
                        </div>
                    </div>

                    <!--     #############################################    -->
                </div>
            </div>
            <!--body wrapper end-->
            <!--footer section start-->
            {{template "inc/notice-dialog.tpl" .}}
            <!--footer section end-->
        </div>

    </div>        <!-- main content end-->
</section>

{{template "inc/foot.tpl" .}}

<script src="/static/js/jquery.easing.1.3.js"></script>
<script src="/static/js/jquery.appear.js"></script>
<script src="/static/js/uisearch.js"></script>
<script src="/static/js/jquery.cubeportfolio.min.js"></script>
<script src="/static/js/custom.js"></script>

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
</script>
</body>
</html>
