angular.module('app')
.config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {

	$routeProvider
		.when('/', {
			templateUrl: 'views/main.html',
			controller: 'MainController',
			controllerAs: 'mainController'
		});

	$locationProvider.html5Mode(true);

}]);
