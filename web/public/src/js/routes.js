angular.module('app')
.config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {

	$routeProvider
		.when('/', {
			templateUrl: 'templates/main.html',
			controller: 'MainController',
			controllerAs: 'mainController'
		});

	$locationProvider.html5Mode(true);

}]);
