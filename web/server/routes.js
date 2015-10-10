module.exports = function(app) {

	// initial load route
	app.get('*', function(req, res) {
		console.log('Request to \'/\'');
		res.sendfile('./public/index.html');
	});

	// websocket routes
	app.io.route('ready', function(req) {
	    req.io.emit('talk', {
	        message: 'ack'
	    });
	});

	app.io.route('move', function(req) {
	    console.log('move request, data=' + req.data);
	    app.redisPublisher.publish('move', req.data);
	});

	app.io.route('look', function(req) {
	    console.log('look request, data=' + req.data);
	    app.redisPublisher.publish('look', req.data);
	});

};
