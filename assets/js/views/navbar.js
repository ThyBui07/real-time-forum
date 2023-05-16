// Extend the HTMLElement class to create the web component
class NavBar extends HTMLElement {

	/**
	 * The class constructor object
	 */
	constructor () {
		// Always call super first in constructor
		super();

	}

	/**
	 * Runs each time the element is appended to or moved in the DOM
	 */
	connectedCallback () {
		console.log('navbar connected!');
		this.render();
	}

	/**
	 * Runs when the element is removed from the DOM
	 */
	disconnectedCallback () {
		console.log('navbar disconnected');
	}

	render() {
		// Render HTML
		this.innerHTML =
			`<header class="blog-header py-3">
            <div class="row flex-nowrap justify-content-between align-items-center">
              <div class="col-4 pt-1">
                <a class="blog-header-logo text-dark" href="/">Cat Forum</a>
              </div>
            </div>
          </header>`;
	}

}

customElements.define('nav-bar', NavBar);


{/* <div class="col-4 d-flex justify-content-end align-items-center">
<a class="btn btn-sm btn-outline-secondary mr-2" href="/signup">Sign up</a>
<a class="btn btn-sm btn-outline-dark mr-2" href="/login">Log in</a>
<a class="btn btn-sm btn-dark mr-2" href="#">Profile</a>

</div> */}