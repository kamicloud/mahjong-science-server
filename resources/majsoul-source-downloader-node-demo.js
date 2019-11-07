const https = require('https');
const fs = require('fs');
const _ = require('lodash')
const path = require('path')
const dict = require('./majsoul-source-mapping')
const resMapping = require('./resversion')

// TODO: 有时读不到文件，会遇到超时问题

let downloadFile = (path, type, prefix, next) => {
    console.log(path)
    let url = `https://www.majsoul.com/1/${prefix}/${path}.${type}`;

    https.get(url, (resp) => {
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
            if (next) {
                next()
            }
        });

    }).on("error", (err) => {
        console.log(url + " Error: " + err.message);
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

var prev = null;

let download = (url, prefix, next) => {
    fs.exists(url, (exists) => {
        if (!exists) {
            ensureFolder(path.dirname(url))
            downloadFile(url.substr(0, url.length - 4), url.substr(url.length - 3), prefix, next)
        } else if (next) {
            next()
        }

    })
}

Object.keys(resMapping.res).map(key => {
    if (!key.startsWith('extendRes')) {
        return
    }
    let prefix = resMapping.res[key].prefix
    let url = key;
    if (!prev) {
        prev = () => {
            download(url, prefix, null)
        }
    } else {
        let temp = prev
        prev = () => {
            download(url, prefix, temp)
        }
    }
})
prev()

function sleep(milliSeconds){
    var startTime =new Date().getTime();
    while(new Date().getTime()< startTime + milliSeconds);
}
