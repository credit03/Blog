CKEDITOR.plugins.add('video1', {
    icons: 'video1',
    init: function (editor) {
        editor.addCommand('video1', new CKEDITOR.dialogCommand('videoDialog1'));

        editor.ui.addButton('video1', {
            label: '视频',
            command: 'video1',
            toolbar: 'insert',
        });
        CKEDITOR.dialog.add('videoDialog1', this.path + 'dialogs/video1.js');
    }
});