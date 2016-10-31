
<!-- 模态框（Modal） -->
<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel"
     aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
                <h4 class="modal-title" id="myModalLabel"></h4>
            </div>
            <div class="modal-body">选择显示位置</div>
            <form method="post" action="#">
                <label class="checkbox-inline">
                    <input type="radio" name="optionsRadiosinline" id="optionsRadios3" value="1" checked>
                    转播图 1
                </label>
                <label class="checkbox-inline">
                    <input type="radio" name="optionsRadiosinline" id="optionsRadios4" value="2"> 转播图 2
                </label>
                <label class="checkbox-inline">
                    <input type="radio" name="optionsRadiosinline" id="optionsRadios5" value="3"> 转播图 3
                </label>

                <div class="modal-footer">
                    <button type="button" class="btn btn-primary" onclick="return SubTop()">提交更改</button>
                    <button type="button" class="btn btn-primary" data-dismiss="modal">关闭</button>
                </div>
                <input type="text" hidden="hidden" name="categoryid" value="" id="categoryid"/>
            </form>
        </div><!-- /.modal-content -->
    </div><!-- /.modal-dialog -->
</div>


<!-- 按钮触发模态框 -->
<!-- 模态框（Modal） -->
<div class="modal fade" id="myLoad" tabindex="-1" role="dialog" aria-labelledby="myLoadLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                <h4 class="modal-title" id="myLoadLabel">进度</h4>
            </div>
            <div class="modal-body"><img src="/static/img/cbp-loading.gif"><span id="myLoadmsg">正在设置</span></div>
            <div class="modal-footer">
                <button type="button" class="btn btn-primary" data-dismiss="modal">关闭</button>
            </div>
        </div><!-- /.modal-content -->
    </div><!-- /.modal -->
</div>