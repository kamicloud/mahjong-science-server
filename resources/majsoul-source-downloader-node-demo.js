const https = require('https');
const fs = require('fs');
const _ = require('lodash')
const path = require('path')
const dict = require('./majsoul-source-mapping')

// TODO: 有时读不到文件，会遇到超时问题

let downloadFile = (path, type) => {
    https.get(`https://www.majsoul.com/1/v0.6.2.w/${path}.${type}`, (resp) => {
        let data = '';
        let arrayBuffer = new ArrayBuffer(8);
        let chunkArray = [];

        // A chunk of data has been recieved.
        resp.on('data', (chunk) => {
            chunkArray = chunkArray.concat(chunk)
        });


        // The whole response has been received. Print out the result.
        resp.on('end', () => {
            let buffer = Buffer.concat(chunkArray);
            let view = new DataView(toArrayBuffer(buffer));

            let i = 0;
            for (; i < buffer.byteLength; i++) {
                view.setInt8(i, 73 ^ view.getInt8(i))
            }

            fs.open(`${path}.${type}`, 'w+', (err, fd) => {
                if (!err) {
                    fs.write(fd, toBuffer(view.buffer), () => {});
                    fs.close(fd, () => {})
                } else {
                    console.log(err)
                }
            })
        });

    }).on("error", (err) => {
        console.log(path + "Error: " + err.message);
    });
}

function toArrayBuffer(buf) {
    var ab = new ArrayBuffer(buf.length);
    var view = new Uint8Array(ab);
    for (var i = 0; i < buf.length; ++i) {
        view[i] = buf[i];
    }
    return ab;
}
function toBuffer(ab) {
    var buf = Buffer.alloc(ab.byteLength);
    var view = new Uint8Array(ab);
    for (var i = 0; i < buf.length; ++i) {
        buf[i] = view[i];
    }
    return buf;
}
function ensureFolder(url) {
    if (fs.existsSync(url)) {
        return true;
    }

    ensureFolder(path.dirname(url));
    fs.mkdirSync(url);
}

_.map(_.get(dict, 'chest.chest_shop.rows_'), 'icon').map(url => {
    ensureFolder(path.dirname(url))
    downloadFile(url.substr(0, url.length - 4), url.substr(url.length - 3))
})