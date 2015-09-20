angular.module('app')
.controller('MainController', ['$scope', '$log', 'socket', function($scope, $log, socket) {

	$scope.mainController = this;

	socket.emit('ready');
	
	socket.on('talk', function(data) {
		$log.log(data.message);
	});

	this.init = function() {
		$log.log('loaded MainController');
	};

}]);
