describe('MainController', function() {

    var mainController, scope, $controller;

    var mockSocket = {
        emit: function(eventName, data, callback) {}
    };

    beforeEach(module('app'));
    beforeEach(inject(function($rootScope, _$controller_) {
        scope = $rootScope.$new();
        $controller = _$controller_;
        mainController = $controller('MainController', {
            $scope : scope
        });
    }));
    beforeEach(function () {
        module(function ($provide) {
            $provide.value('socket', mockSocket);
        });
    });



    it('main controller should be defined', function() {
        expect(mainController).toBeDefined();
    });

});
