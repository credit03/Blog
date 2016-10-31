CKEDITOR.dialog.add('videoDialog', function (editor) {

    var url = "";
    var data = "";
    var onTouch = false;


    var m = /^\s*(\d+)((px)|\%)?\s*$/i, z = /(^\s*(\d+)((px)|\%)?\s*$)|^$/i, r = /^\d+px$/,
        A = function () {
            var a = this.getValue(), b = this.getDialog(), c = a.match(m);
            c && ("%" == c[2] && n(b, !1), a = c[1]);
            b.lockRatio && (c = b.originalElement, "true" == c.getCustomData("isReady") && ("txtHeight" == this.id ? (a && "0" != a && (a = Math.round(a / c.$.height * c.$.width)), isNaN(a) || b.setValueOf("info", "txtWidth", a)) : (a && "0" != a && (a = Math.round(a / c.$.width * c.$.height)), isNaN(a) || b.setValueOf("info", "txtHeight", a))));
            e(b)
        }, e = function (a) {
            if (!a.originalElement || !a.preview)return 1;
            a.commitContent(4, a.preview);
            return 0
        }, w, n = function (a,
                            b) {
            if (!a.getContentElement("info", "ratioLock"))return null;
            var c = a.originalElement;
            if (!c)return null;
            if ("check" == b) {
                if (!a.userlockRatio && "true" == c.getCustomData("isReady")) {
                    var d = a.getValueOf("info", "txtWidth"), f = a.getValueOf("info", "txtHeight"), c = 1E3 * c.$.width / c.$.height, g = 1E3 * d / f;
                    a.lockRatio = !1;
                    d || f ? isNaN(c) || isNaN(g) || Math.round(c) != Math.round(g) || (a.lockRatio = !0) : a.lockRatio = !0
                }
            } else void 0 !== b ? a.lockRatio = b : (a.userlockRatio = 1, a.lockRatio = !a.lockRatio);
            d = CKEDITOR.document.getById(t);
            a.lockRatio ?
                d.removeClass("cke_btn_unlocked") : d.addClass("cke_btn_unlocked");
            d.setAttribute("aria-checked", a.lockRatio);
            CKEDITOR.env.hc && d.getChild(0).setHtml(a.lockRatio ? CKEDITOR.env.ie ? "■" : "▣" : CKEDITOR.env.ie ? "□" : "▢");
            return a.lockRatio
        }, B = function (a, b) {
            var c = a.originalElement;
            if ("true" == c.getCustomData("isReady")) {
                var d = a.getContentElement("info", "txtWidth"), f = a.getContentElement("info", "txtHeight"), g;
                b ? c = g = 0 : (g = c.$.width, c = c.$.height);
                d && d.setValue(g);
                f && f.setValue(c)
            }
            e(a)
        }, C = function (a, b) {
            function c(a, b) {
                var c =
                    a.match(m);
                return c ? ("%" == c[2] && (c[1] += "%", n(d, !1)), c[1]) : b
            }

            if (1 == a) {
                var d = this.getDialog(), f = "", g = "txtWidth" == this.id ? "width" : "height", e = b.getAttribute(g);
                e && (f = c(e, f));
                f = c(b.getStyle(g), f);
                this.setValue(f)
            }
        }, x, u = function () {
            var a = this.originalElement, b = CKEDITOR.document.getById(p);
            a.setCustomData("isReady", "true");
            a.removeListener("load", u);
            a.removeListener("error", h);
            a.removeListener("abort", h);
            b && b.setStyle("display", "none");
            this.dontResetSize || B(this, !1 === d.config.image_prefillDimensions);
            this.firstLoad &&
            CKEDITOR.tools.setTimeout(function () {
                n(this, "check")
            }, 0, this);
            this.dontResetSize = this.firstLoad = !1;
            e(this)
        }, h = function () {
            var a = this.originalElement, b = CKEDITOR.document.getById(p);
            a.removeListener("load", u);
            a.removeListener("error", h);
            a.removeListener("abort", h);
            a = CKEDITOR.getUrl(CKEDITOR.plugins.get("image").path + "images/noimage.png");
            this.preview && this.preview.setAttribute("src", a);
            b && b.setStyle("display", "none");
            n(this, !1)
        }, q = function (a) {
            return CKEDITOR.tools.getNextId() + "_" + a
        }, t = q("btnLockSizes"),
        y = q("btnResetSize"), p = q("ImagePreviewLoader"), E = q("previewLink"), D = q("previewImage");
    return {
        title: '打开视频URL',
        minWidth: 400,
        minHeight: 200,
        contents: [
            {
                id: "info", label: "URL", accessKey: "I", elements: [{
                type: "vbox", padding: 0, children: [{
                    type: "hbox", widths: ["280px", "110px"], align: "right", children: [{
                        id: "txtUrl", type: "text", label: "url", required: !0, onChange: function () {
                            var a = this.getDialog(), b = this.getValue();
                            if (0 < b.length) {
                                url = b;
                            }
                        }, validate: CKEDITOR.dialog.validate.notEmpty("请上传视频到服务器")
                    }]
                }]
            }]
            }]/*, {
         id: 'flv',
         label: '上传视频',
         elements: [
         {
         type: "file",
         id: 'videoupload',
         label: '打开视频路径',
         style: "height:40px",
         filebrowser: "uploadButton",
         },
         {
         type: "fileButton",
         id: "uploadButton",
         filebrowser: "info:txtUrl",
         label: "上传视频",
         "for": ["flv",
         "videoupload"]
         , onClick: function () {
         onTouch = true;
         }
         , validate: CKEDITOR.dialog.validate.notEmpty("请选择文件...")
         }
         ]
         }
         ],
         onShow: function () {
         var m = CKEDITOR.dom.element.createFromHtml("\<p/\>");
         editor.insertElement(m);
         }*/
        ,
        onOk: function () {

            /*  if (CKEDITOR.instances.editor) {
             data = CKEDITOR.instances.editor.getData();
             CKEDITOR.instances.editor.destroy();
             }

             CKEDITOR.config.toolbar_Basic =
             [
             {
             name: 'clipboard',
             items: ['Cut', 'Copy', 'Paste', 'PasteText', 'PasteFromWord', '-', 'Undo', 'Redo']
             },
             {name: 'editing', items: ['Scayt']},
             {name: 'links', items: ['Link', 'Unlink', 'Anchor']},
             {name: 'insert', items: ['Image', 'video1', 'Table', 'HorizontalRule', 'SpecialChar']},
             {name: 'tools', items: ['Maximize']},
             {name: 'document', items: ['Source']},
             '/',
             {name: 'basicstyles', items: ['Bold', 'Italic', 'Strike', '-', 'RemoveFormat']},
             {
             name: 'paragraph',
             items: ['NumberedList', 'BulletedList', '-', 'Outdent', 'Indent', '-', 'Blockquote']
             },
             {name: 'styles', items: ['Styles', 'Format']},
             {name: 'about', items: ['About']}
             ];

             CKEDITOR.replace('editor', CKEDITOR.config);
             CKEDITOR.instances.editor.updateElement();
             */
            //\x3cdiv class\x3d"cke_tpl_item"\x3e    \x3c/div\x3e
            var html = "";
            if (onTouch) {
      /*          var id = UUID(8, 10); // "47473046"
                CallVideoUrl(url, id);
                html = "\x3cdiv class='text-center'\x3e \x3ciframe  uuid='" + id + "' width='560' height='315' src='" + url + "' frameborder='1' allowfullscreen\x3e\x3c/iframe\x3e\x3c/div\x3e";
                onTouch = false;*/
            } else {
                var id = UUID(8, 10); // "47473046"
                VideoUrl(url, id);
                html = "\x3cdiv class='text-center'\x3e \x3ciframe uuid='" + id + "' width='560' height='315' src='" + url + "' frameborder='1' allowfullscreen\x3e\x3c/iframe\x3e\x3c/div\x3e";

            }
            var m = CKEDITOR.dom.element.createFromHtml(html);
            editor.insertElement(m);

        }
    }
});
