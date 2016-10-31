<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>修改头像</title>
    <link href="/static/css/style.css" rel="stylesheet">
    {{template "inc/meta.tpl" .}}
    <link rel="stylesheet" type="text/css" href="/static/js/bootstrap-cropper/cropper.min.css"/>
</head>
<body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
    <!-- main content start-->
    <div class="main-content">
        <!-- header section start-->
        <div class="header-section">
            <!--toggle button start-->
            <a class="toggle-btn"><i class="fa fa-bars"></i></a> {{template "inc/user-info.tpl" .}}
        </div>

        <div class="wrapper">
            <div class="row">
                <div class="col-lg-12">
                    <section class="panel">
                        <header class="panel-heading"> 设置头像</header>
                        <div class="panel-body">
                            <form class="form-horizontal adminex-form" id="avatar-form" method="post"
                                  enctype="multipart/form-data"
                                  action="/myajax/updateuserpic/{{.User.Id}}">
                                {{ .xsrfdata }}
                                <div class="col-md-9">
                                    <div style="height:500px;"><img id="image" src="{{.User.Pic}}"/></div>
                                </div>
                                <div class="col-md-3">
                                    <div class="btn-group">
                                        <button type="button" class="btn btn-primary js-cropper-big" data-method="zoom"
                                                data-option="0.1" title="Zoom In"><span class="docs-tooltip"
                                                                                        data-toggle="tooltip" title=""
                                                                                        data-original-title="$().cropper(&quot;zoom&quot;, 0.1)"> <span
                                                class="fa fa-search-plus"></span> </span></button>
                                        <button type="button" class="btn btn-primary js-cropper-small"
                                                data-method="zoom" data-option="-0.1" title="Zoom Out"><span
                                                class="docs-tooltip" data-toggle="tooltip" title=""
                                                data-original-title="$().cropper(&quot;zoom&quot;, -0.1)"> <span
                                                class="fa fa-search-minus"></span> </span></button>
                                    </div>
                                    <label class="btn btn-primary btn-upload" for="inputImage"
                                           title="Upload image file">
                                        <input type="file" class="sr-only" id="inputImage" name="file" accept="image/*">
                                        <span class="docs-tooltip" data-toggle="tooltip"
                                              title="Import image with Blob URLs"> <span class="fa fa-upload"></span> 上传 </span>
                                    </label>
                                    <button type="submit" class="btn btn-primary" onclick="return checkfile()">保存
                                    </button>
                                    <input type="hidden" class="form-control" value="{{.User.Pic}}"
                                           id="avatar" name="avatar" placeholder="x">
                                    <input type="hidden" class="form-control" id="dataX" name="dataX" placeholder="x">
                                    <input type="hidden" class="form-control" id="dataY" name="dataY" placeholder="y">
                                    <input type="hidden" class="form-control" id="dataWidth" name="dataWidth"
                                           placeholder="width">
                                    <input type="hidden" class="form-control" id="dataHeight" name="dataHeight"
                                           placeholder="height">
                                    <input type="hidden" class="form-control" id="dataRotate" name="dataRotate"
                                           placeholder="rotate">
                                    <input type="hidden" class="form-control" id="dataScaleX" name="dataScaleX"
                                           placeholder="scaleX">
                                    <input type="hidden" class="form-control" id="dataScaleY" name="dataScaleY"
                                           placeholder="scaleY">
                                </div>
                            </form>
                        </div>
                    </section>
                </div>
            </div>
        </div>

        <!--footer section end-->
    </div>
    <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
<script type="text/javascript" src="/static/js/bootstrap-cropper/cropper.min.js"></script>
<script>


    function checkfile() {
        var f = document.getElementById("inputImage").files;
        if (f == null || f[0].size == 0 || f[0].type.length() == 0) {
            alert("请选择图片");
            return false;
        }
        return true;
    }


    $(function () {
        var $image = $('#image');
        var $dataX = $('#dataX');
        var $dataY = $('#dataY');
        var $dataHeight = $('#dataHeight');
        var $dataWidth = $('#dataWidth');
        var $dataRotate = $('#dataRotate');
        var $dataScaleX = $('#dataScaleX');
        var $dataScaleY = $('#dataScaleY');
        $image.cropper({
            viewMode: 1,
            dragMode: 'move',
            aspectRatio: 1,
            autoCropArea: 0.6,
            restore: false,
            guides: false,
            highlight: false,
            //cropBoxMovable: false,
            cropBoxResizable: false,
            built: function () {
                croppable = true;
            },
            crop: function (e) {
                $dataX.val(Math.round(e.x));
                $dataY.val(Math.round(e.y));
                $dataHeight.val(Math.round(e.height));
                $dataWidth.val(Math.round(e.width));
                $dataRotate.val(e.rotate);
                $dataScaleX.val(e.scaleX);
                $dataScaleY.val(e.scaleY);
            }
        });

        $('.js-cropper-big').on('click', function () {
            $image.cropper("zoom", 0.1)
        });

        $('.js-cropper-small').on('click', function () {
            $image.cropper("zoom", -0.1)
        });

        //$image.cropper("getData")

        //http://studygolang.com/articles/2581

        /*$('#avatar-form').validate({
         ignore: '',
         rules: {},
         messages: {},
         submitHandler: function (form) {
         $(form).ajaxSubmit({
         type: 'POST',
         dataType: 'json',
         success: function (data) {
         dialogInfo(data.message)
         if (data.code) {
         //setTimeout(function(){window.location.href="/project/test/"+$('#projectid').val()}, 2000);
         } else {

         }
         setTimeout(function () {
         $('#dialogInfo').modal('hide');
         }, 1000);
         }
         });
         }
         });*/


        // Import image
        var $inputImage = $('#inputImage');
        var URL = window.URL || window.webkitURL;
        var blobURL;

        if (URL) {
            $inputImage.change(function () {
                var files = this.files;
                var file;

                if (!$image.data('cropper')) {
                    return;
                }

                if (files && files.length) {
                    file = files[0];

                    if (/^image\/\w+$/.test(file.type)) {
                        blobURL = URL.createObjectURL(file);
                        $image.one('built.cropper', function () {

                            // Revoke when load complete
                            URL.revokeObjectURL(blobURL);
                        }).cropper('reset').cropper('replace', blobURL);
                        //$inputImage.val('');
                    } else {
                        window.alert('Please choose an image file.');
                    }
                }
            });
        } else {
            $inputImage.prop('disabled', true).parent().addClass('disabled');
        }

    })
</script>
</body>
</html>
