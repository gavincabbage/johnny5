module.exports = function(app) {

	// api routes

	// frontend route
	app.get('*', function(req, res) {
		console.log('Request to \'/\'');
		res.sendfile('./public/index.html');
	});

};
