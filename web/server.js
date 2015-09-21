var express = require('express.io');
var redis = require('redis');

var app = express();
app.http().io();
var port = 8080;

var redisClient = redis.createClient();
var redis_port = 6379;


// Setup the ready route, and emit talk event.
app.io.route('ready', function(req) {

    redisClient.on("message", function(channel, message) {
        console.log("got a message on channel " + channel);
        if (channel === "mychan") {
            console.log(message);
            req.io.emit('talk', {
                message: 'redis: ' + message
            });
        }
    });


    req.io.emit('talk', {
        message: 'io event from an io route on the server'
    });
});


redisClient.subscribe("mychan");


app.use(express.static(__dirname + '/public'));
require('./server/routes')(app);
app.listen(port);
console.log('Listening on port ' + port);
exports = module.exports = app;
