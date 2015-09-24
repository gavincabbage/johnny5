var express = require('express.io');
var redis = require('redis');

var app = express();
app.http().io();
var port = 8080;

var redisClient = redis.createClient();
var redis_port = 6379;

app.io.route('ready', function(req) {
    req.io.emit('talk', {
        message: 'ack'
    });
});

redisClient.on('message', function(channel, message) {
    console.log('got a message on channel ' + channel);
    app.io.broadcast('talk', {
        message: 'redis: ' + message
    });
});

redisClient.subscribe('mychan');

app.use(express.static(__dirname + '/public'));
require('./server/routes')(app);
app.listen(port);
console.log('Listening on port ' + port);
exports = module.exports = app;
