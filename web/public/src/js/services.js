(function() {

	var app = angular.module('app');

	app.factory('appServices', ['$http', function($http) {
		return {
			exampleHttpGetService : function() {
				return $http({
					url: '/api/regions',
					method: 'GET',
					headers: {
						'Content-Type': 'application/json',
					}
				});
			}
		};
	}]);

	app.factory('socket', ['$rootScope', function($rootScope) {
		var socket = io.connect();
		return {
			on: function(eventName, callback) {
				socket.on(eventName, function() {
					var args = arguments;
					$rootScope.$apply(function() {
						callback.apply(socket, args);
					});
				});
			},
			emit: function(eventName, data, callback) {
				socket.emit(eventName, data, function() {
					var args = arguments;
					$rootScope.$apply(function() {
						if (callback) {
							callback.apply(socket, args);
						}
					});
				});
			}
		};
	}]);

})();
