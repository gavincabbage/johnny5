var express = require('express.io');
var redis = require('redis');

var app = express();
app.http().io();
port = 8080;

app.redisSubscriber = redis.createClient();
app.redisPublisher = redis.createClient();

app.redisSubscriber.on('message', function(channel, message) {
    console.log('got a message on channel ' + channel);
    app.io.broadcast('talk', {
        message: 'redis: ' + message
    });
});

app.redisSubscriber.subscribe('mychan');

app.use(express.static(__dirname + '/public'));
require('./server/routes')(app);
app.listen(port);
console.log('Listening on port ' + port);
exports = module.exports = app;
