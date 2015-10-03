angular.module('app')
.controller('MainController', ['$scope', '$log', 'socket', function($scope, $log, socket) {

	$scope.mainController = this;

	this.leftDistance = 0.0;
	this.frontDistance = 0.0;
	this.rightDistance = 0.0;

	this.moveDirections = ['forward', 'left', 'stop', 'right', 'back'];
	this.lookDirections = ['up', 'left', 'center', 'right', 'down'];

	socket.on('talk', function(data) {
		$log.log(data.message);
	});

	this.move = function(direction) {
		$log.log('move ' + direction);
		socket.emit('move', direction);
	};

	this.look = function(direction) {
		$log.log('look ' + direction);
		socket.emit('look', direction);
	};

	this.init = function() {
		$log.log('loaded MainController');
		socket.emit('ready');
	};

}]);
