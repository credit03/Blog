<!--头部-->
{{define "header"}}
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta name=”renderer” content=”webkit|ie-comp|ie-stand”/>
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,Chrome=1"/>
    <meta http-equiv="X-UA-Compatible" content="IE=9"/>


    <!-- Bootstrap -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link rel="stylesheet" href="/static/css/animate.css">
    <link rel="stylesheet" href="/static/css/font-awesome.min.css">
    <link rel="stylesheet" href="/static/css/jquery.bxslider.css">
    <link rel="stylesheet" type="text/css" href="/static/css/normalize.css"/>
    <link rel="stylesheet" type="text/css" href="/static/css/demo.css"/>
    <link rel="stylesheet" type="text/css" href="/static/css/set1.css"/>
    <link href="/static/css/overwrite.css" rel="stylesheet">
    <link href="/static/css/style.css" rel="stylesheet">


    <!--[if IE]>
  <script src=”http://html5shiv.googlecode.com/svn/trunk/html5.js”></script>
  < ![endif]-->
    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
    <script src="http://apps.bdimg.com/libs/html5shiv/3.7/html5shiv.min.js"></script>
    <script src="http://apps.bdimg.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
    <script type="text/javascript">

        var _gaq = _gaq || [];
        _gaq.push(['_setAccount', 'UA-30181385-1']);
        _gaq.push(['_trackPageview']);

        (function () {
            var ga = document.createElement('script');
            ga.type = 'text/javascript';
            ga.async = true;
            ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
            var s = document.getElementsByTagName('script')[0];
            s.parentNode.insertBefore(ga, s);
        })();

        function openhref(href) {
            window.open(href, "_self");
        }

        function checkKernel() {
            if (window.XMLHttpRequest) {
                if (window.ActiveXObject) {//IE6以上有这个属性，是IE内核独有的，其他内核不具备
                    window.alert("亲，页面对您的浏览器不兼容;请升级到IE10以上或360极速模式，Chrome，火狐等非IE9内核以下的浏览器");
                }

            }
        }


    </script>

    {{end}}