<!--导航栏-->
{{define "navbar"}}
<nav class="navbar navbar-default navbar-fixed-top" role="navigation">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse"
                    data-target=".navbar-collapse.collapse">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="/"><span>Cre</span></a>
        </div>
        <div class="navbar-collapse collapse">
            <div class="menu">
                <ul class="nav nav-tabs" role="tablist">
                    {{range $data:=.NavArr}}
                    {{if .IsShow}}
                    {{template "liloop" .}}
                    {{end}}
                    {{end}}
                    <li style="margin-left: 15px">
                        {{if .Islogin}}
                        <button type="button" class="btn btn-primary" onclick="openhref(value)" value="{{.LoginOP}}">退出
                        </button>
                        {{else}}
                        <button type="button" class="btn btn-primary" onclick="openhref(value)" value="{{.LoginOP}}">
                            登录/注册
                        </button>
                        {{end}}

                    </li>
                </ul>
            </div>
        </div>

    </div>
</nav>
{{end}}

{{define "liloop"}}
<li {{if .Active}}class="active" {{end}}><a href='{{if .Active}}#{{else}}{{.Url}}{{end}}'>{{.Name}}</a></li>
{{end}}