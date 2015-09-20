var express = require('express.io');
var app = express();
app.http().io();

var port = 8080;

app.use(express.static(__dirname + '/public'));

require('./server/routes')(app);


// Setup the ready route, and emit talk event.
app.io.route('ready', function(req) {
    req.io.emit('talk', {
        message: 'io event from an io route on the server'
    })
})

// Send the client html.
app.get('/', function(req, res) {
    res.sendfile(__dirname + '/client.html')
})

app.listen(port);
console.log('Listening on port ' + port);
exports = module.exports = app;
