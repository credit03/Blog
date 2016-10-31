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
            <li>
                <a class="dropdown-toggle" data-toggle="dropdown"> <i
                        class="glyphicon glyphicon-tasks "></i><span> 文章管理</span></a>

                <ul class="dropdown-menu">
                    <li><a href='/myajax/gategory_manage/{{.User.Id}}?op=-1&type=SectionType'>全部文章</a></li>
                    <li class="divider"></li>
                    <li><a href='/myajax/gategory_manage/{{.User.Id}}?op=1&type=SectionType'>科技版块</a></li>
                    <li class="divider"></li>
                    <li><a href='/myajax/gategory_manage/{{.User.Id}}?op=0&type=SectionType'>生活版块</a></li>
                    <li class="divider"></li>
                    <li><a href='/myajax/gategory_manage/{{.User.Id}}?op=2&type=SectionType'>其他版块</a></li>
                    <li class="divider"></li>
                    <li><a href='/myajax/gategory_manage/{{.User.Id}}?op=0&type=ContentType'>文本文章</a></li>
                    <li class="divider"></li>
                    <li><a href='/myajax/gategory_manage/{{.User.Id}}?op=1&type=ContentType'>图片文章</a></li>
                    <li class="divider"></li>
                    <li><a href='/myajax/gategory_manage/{{.User.Id}}?op=2&type=ContentType'>视频文章</a></li>
                </ul>
            </li>
            <!--
                        <li><a href="/myajax/gategory_manage/{{.User.Id}}"><i class="glyphicon glyphicon-tasks"></i>
                            <span> 文章管理</span></a></li>
            -->
            {{if .User.IsAdmin}}
            <li><a href="/category/section?op=1&type=ContentType"><i class="glyphicon glyphicon-heart"></i>
                <span> 头条设置</span></a></li>
            <li><a href="/myajax/user_manage/{{.User.Id}}"><i class="glyphicon glyphicon-user"></i>
                <span> 用户管理</span></a></li>
            {{end}}
            <li><a href="/"><i class="glyphicon glyphicon-chevron-left"></i> <span> 返回首页</span></a></li>
        </ul>

        <!--sidebar nav end-->
    </div>
</div>
<!-- left side end-->
