var xhr = new XMLHttpRequest();
xhr.open('GET', 'https://www.majsoul.com/1/v0.6.2.w/extendRes/charactor/sigongxiasheng/smallhead.png');
xhr.responseType = 'arraybuffer';
xhr.send()
// byte array

xhr.onload = () => {
    console.log(xhr.response)

    var byteArray = xhr.response;
    var view = new DataView(byteArray);
    // DataView是为了遍历和修改byte array
    console.log(byteArray, view)
    if (byteArray) {
        for (var i = 0; i < byteArray.byteLength; i++) {
            // 雀魂release环境会对二进制返回数据做一次混淆
            view.setInt8(i, 73 ^ view.getInt8(i))
        }
        byteArray = view.buffer
        console.log(byteArray)
        var blob = new Blob([byteArray], {type: 'image/png'});
        // 前端渲染图片
        console.log(window.URL.createObjectURL(blob))
    } else {
        console.log('failed')
    }
}
