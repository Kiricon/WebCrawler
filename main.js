const http = require('http');

try {
    http.get({
        host: 'www.google.com',
    }, (response) => {
        response.on('data', (data) => {
            console.log(data.toString().match(/href="(.*?)"/g).);
        });

        response.on('error', (error) => {
            console.log(error);
        });
    });
} catch(e) {
    console.log(e);
}
