//document.addEventListener('polymer-ready', function() {
//    var navicon = document.getElementById('navicon');
//    var drawerPanel = document.getElementById('drawerPanel');
//    navicon.addEventListener('click', function() {
//        drawerPanel.togglePanel();
//    });
//});

var app = document.querySelector('#application');

page('/', index);
page('/login', login);
page('/register', registration);

function index () {
    app.route = 'index';
}

function registration() {
    app.route = 'register';
}

function login() {
    app.route = 'login';
}

page();
