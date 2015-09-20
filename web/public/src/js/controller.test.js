describe('MainController', function() {

  var mainController, scope, $controller;

  beforeEach(module('app'));
  beforeEach(inject(function($rootScope, _$controller_) {
    scope = $rootScope.$new();
    $controller = _$controller_;
    mainController = $controller('MainController', {
      $scope : scope
    });
  }));

  it('main controller should be defined', function() {
    expect(mainController).toBeDefined();
  });

});
