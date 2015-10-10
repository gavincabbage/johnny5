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
    app.io.broadcast(channel, {
        message: message
    });
    var channelPath = channel.split(".");
    console.log(channelPath);
});

var redisSubscriptions = ['mychan', 'distance.left', 'distance.right', 'distance.center'];
for (ndx in redisSubscriptions) {
    var chan = redisSubscriptions[ndx];
    console.log('subscribing to ' + chan);
    app.redisSubscriber.subscribe(chan);
}

app.use(express.static(__dirname + '/public'));
require('./server/routes')(app);
app.listen(port);
console.log('Listening on port ' + port);
exports = module.exports = app;
