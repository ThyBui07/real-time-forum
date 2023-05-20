// create document click that watches the links only
document.addEventListener("click", (e) => {
	const { target } = e;
	if (!target.matches("a")) {
		return;
	}
	e.preventDefault();
	const href = e.target.href;
	urlRoute(href);
});

window.addEventListener('popstate', function(event) {
	console.log('popstate fired!');
  const newURL = event.currentTarget.location.href;
  var newHash = newURL.split('#')[1] || '';

  if (newHash === '/') {
	newHash = '';
  }
  const href = '/' + newHash;
  console.log('href: ', href)
  urlRoute(href);

  });


// create an object that maps the url to the template, title, and description
const urlRoutes = {
	"/login": {
		template: "<log-in></log-in>"
	},
	404: {
		template: "<p>404 Page Not Found</p>"
	},
	"/": {
		template: "<home-page></home-page>"
	},
	"/signup": {
		template: "<sign-up></sign-up>"
	}
};

// create a function that watches the url and calls the urlLocationHandler
const urlRoute = (route) => {
	console.log('urlRoute function called')
	console.log('route: ', route)
	window.history.pushState({}, "",route);
	urlLocationHandler()
};

// create a function that handles the url location
const urlLocationHandler = async () => {
	var location = window.location.pathname;
	console.log('location in url handler: ', location)

	var sessionID =  document.cookie
	.split(';')
	.find(cookie => cookie.startsWith('session-name'));

	 if (sessionID != undefined) {
		sessionID = sessionID.split('=', 2)[1]
	 }
	// if there is no session ID, redirect to login page
	if (sessionID == undefined && location == "/") {
		window.location.hash = 'login';
		return;
	  }
	if (sessionID != undefined && (location == "/login" || location == "/signup")) {
		location = "/";
	}
	// get the route object from the urlRoutes object
	const route = urlRoutes[location] || urlRoutes["404"];
	// get the html from the template
	const html =route.template
	// set the content of the content div to the html
	document.getElementById("app").innerHTML = html;
};

// add an event listener to the window that watches for url changes
// window.onpopstate = urlLocationHandler;

// call the urlLocationHandler function to handle the initial url
window.route = urlRoute;
urlLocationHandler();