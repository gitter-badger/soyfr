(function(window, document, undefined) {
    'use strict';

    // Select auto-binding template and use as the top level of our app
    var app = document.querySelector('#app');
    app.addEventListener('template-bound', function() {
        var pages = document.querySelector('#pages');

        var loginAction = function () {
            app.route = 'login';
        };

        var userListAction = function () {
            app.route = 'users'
        };

        var userProfileAction = function () {
            app.route = 'test'
        };

        var indexAction = function () {
            loginAction();
        };

        var routes = {
            '/' : indexAction,
            '/login': loginAction,
            '/user': userListAction,
            '/user/:username': userProfileAction
        };

        var router = Router(routes);
        router.configure({html5history: true});
        router.init();

        // Listen for pages to fire their change-route event
        // Instead of letting them change the route directly,
        // handle the event here and change the route for them
        document.addEventListener('change-route', function(e) {
            if (e.detail) {
                router.setRoute(e.detail);
            }
        });

        // Similar to change-route, listen for when a page wants to go
        // back to the previous state and change the route for them
        document.addEventListener('change-route-back', function() {
            history.back();
        });

        // Set duration for core-animated-pages transitions
        CoreStyle.g.transitions.duration = '0.2s';
    });

})(window, document);
