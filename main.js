const http = require('http');

http.get({
    host: 'valenciana.me',
}, (response) => {
    response.on('data', (data) => {
        console.log(data.toString());
    });
});
