CKEDITOR.plugins.add('video', {
    icons: 'video',
    init: function (editor) {
        editor.addCommand('video', new CKEDITOR.dialogCommand('videoDialog'));

        editor.ui.addButton('video', {
            label: '视频',
            command: 'video',
            toolbar: 'insert',
            id: 'videotool'

        });
        CKEDITOR.dialog.add('videoDialog', this.path + 'dialogs/video.js');
    }
});