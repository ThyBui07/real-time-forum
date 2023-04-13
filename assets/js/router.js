// create document click that watches the nav links only
document.addEventListener("click", (e) => {
	const { target } = e;
	if (!target.matches("header a")) {
		return;
	}
	e.preventDefault();
	urlRoute();
});

// create an object that maps the url to the template, title, and description
const urlRoutes = {
	"/": {
		template: "<home-page></home-page>"
	},
	404: {
		template: "<p>404 Page Not Found</p>"
	},
	"/greet": {
		template: "<p>Hi there!</p>"
	
	},
	"/login": {
		template: "<log-in></log-in>"
	}
};

// create a function that watches the url and calls the urlLocationHandler
const urlRoute = (event) => {
	event = event || window.event; // get window.event if event argument not provided
	event.preventDefault();
	// window.history.pushState(state, unused, target link);
	window.history.pushState({}, "", event.target.href);
	urlLocationHandler();
};

// create a function that handles the url location
const urlLocationHandler = async () => {
	const location = window.location.pathname; // get the url path
	const auth = localStorage.getItem("auth");
	// if the path length is 0, set it to primary page route
	// if (location.length == 0 && auth != 1) {
	// 	//redirect to login page
	// 	location = "/login";
	// }

	if (location.length == 0) {
		//redirect to login page
		location = "/";
	}
	// get the route object from the urlRoutes object

	const route = urlRoutes[location] || urlRoutes["404"];
	// get the html from the template
	const html =route.template
	// set the content of the content div to the html
	document.getElementById("output").innerHTML = html;
};

// add an event listener to the window that watches for url changes
window.onpopstate = urlLocationHandler;
// call the urlLocationHandler function to handle the initial url
window.route = urlRoute;
// call the urlLocationHandler function to handle the initial url
urlLocationHandler();