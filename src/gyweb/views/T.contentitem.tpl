{{define "textitem"}}
<article class="panel">
    <div class="row">
        <div class="col-lg-10">
            <div class="post-image">
                <div class="post-heading category-pic">
                    <h3><img alt="" src="{{.Author.Pic}}">&nbsp;<a href="/browse_category?id={{.Id}}">{{ .Title}}</a>
                    </h3>
                </div>
            </div>

            <div class="col-lg-10">
                {{template "conitem" .}}
            </div>
        </div>
    </div>
</article>

{{end}}

{{define "aimgitem"}}

<article>
    <div class="row">
        <div class="col-lg-10">
            <div class="post-image">
                <div class="post-heading category-pic">
                    <h3><img alt="" src="{{.Author.Pic}}">&nbsp;<a href="/browse_category?id={{.Id}}">{{ .Title}}</a>
                    </h3>
                </div>
                {{range .Images}}
                <img src="/{{.Url}}" alt="" style="max-width: 480px; height: 320px"/>
                {{end}}
            </div>

            <div class="col-lg-10">
                {{template "conitem" .}}
            </div>
        </div>
    </div>
</article>


{{end}}

{{define "moreimgitem"}}

<article>
    <div class="row">
        <div class="col-lg-10">
            <div class="post-slider">
                <div class="post-heading category-pic">
                    <h3><img alt="" src="{{.Author.Pic}}">&nbsp;<a href="/browse_category?id={{.Id}}">{{ .Title}}</a>
                    </h3>
                </div>
                <div class="col-lg-11">
                    <!-- start flexslider -->
                    <div id="post-slider" class="postslider flexslider">
                        <ul class="slides">
                            {{range .Images}}
                            <li>
                                <img src="/{{.Url}}" style="max-width:480px;height: 320px" alt=""/>
                            </li>
                            {{end}}
                        </ul>
                    </div>
                </div>
                <!-- end flexslider -->
            </div>
        </div>
        <div class="col-lg-10">
            {{template "conitem" .}}
        </div>
    </div>
</article>

{{end}}


{{define "videoitem"}}
<article class="panel">
    <div class="row">
        <div class="col-lg-10">
            <div class="post-video">
                <div class="post-heading category-pic">
                    <h3><img alt="" src="{{.Author.Pic}}">&nbsp;<a href="/browse_category?id={{.Id}}">{{ .Title}}</a>
                    </h3>
                </div>
                <!-- 视频-->
                <iframe width="560" height="315"
                        src="{{.VideoUrl}}"
                        frameborder="1" allowfullscreen></iframe>
            </div>
            <div class="col-lg-10">
                {{template "conitem" .}}
            </div>
        </div>
    </div>
</article>

{{end}}

{{define "conitem"}}
<p>
    {{.Describe}}
</p>
<div class="container">
    <ul class="list-inline">
        <li><i class="fa fa-calendar"></i> {{getDateMH .Created}}</li>
        <li><i class="fa fa-user"></i><a href="/user/borwseinfo?uid={{.Author.Id}}"> {{.Author.NickName}}</a></li>
        {{if eq .SectionType 0}}
        <li><i class="fa fa-folder-open"></i><a href='{{urlfor "CategoryController.Section"}}?op=0'> 生活</a></li>
        {{else}}
        {{if eq .SectionType 1}}
        <li><i class="fa fa-folder-open"></i><a href='{{urlfor "CategoryController.Section"}}?op=1'> 科技</a></li>
        {{else}}
        {{if eq .SectionType 2}}
        <li><i class="fa fa-folder-open"></i><a href='{{urlfor "CategoryController.Section"}}?op=2'> 其他</a></li>
        {{end}}
        {{end}}
        {{end}}

        <li><i class="glyphicon glyphicon-eye-open"></i><a href="/browse_category?id={{.Id}}"> 查看详细</a></li>
    </ul>
</div>
{{end}}