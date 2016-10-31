<!--notification menu start -->
<div class="menu-right">
  <ul class="notification-menu">
    <li> <a href="javascript:;" class="btn btn-default dropdown-toggle" data-toggle="dropdown"> <img src="{{.User.Pic}}" alt="{{.User.Name}}" />  <span class="caret"></span> </a>
      <ul class="dropdown-menu dropdown-menu-usermenu pull-right">
        <li><a href="/myajax/editinfo/{{.User.Id}}"><i class="fa fa-cog"></i> 基本资料</a></li>
		<li><a href="/myajax/eidtuserpic/{{.User.Id}}/"><i class="fa fa-camera"></i> 更换头像</a></li>
		<li><a href="/myajax/eidtuserpwd"><i class="fa fa-cog"></i> 更换密码</a></li>
        <li><a href="/user/loginpage?type=logout&uname={{.User.Name}}"><i class="fa fa-sign-out"></i> 退出</a></li>
      </ul>
    </li>
  </ul>
</div>
<!--notification menu end -->
