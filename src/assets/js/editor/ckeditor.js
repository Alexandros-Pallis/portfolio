import ClassiceEditor from '@ckeditor/ckeditor5-build-classic'

export default () => {
    ClassiceEditor
        .create(document.querySelector('#editor'))
        .then(editor => {
            console.log(editor)
        })
        .catch(error => {
            console.error(error)
        })
}
