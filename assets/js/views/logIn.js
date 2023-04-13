// Extend the HTMLElement class to create the web component
class logIn extends HTMLElement {

	/**
	 * The class constructor object
	 */
	constructor () {
		// Always call super first in constructor
		super();

	}

    count() {
        console.log('counting');
    }
	/**
	 * Runs each time the element is appended to or moved in the DOM
	 */
	connectedCallback () {
		this.render();
        this.count();
	}

	/**
	 * Runs when the element is removed from the DOM
	 */
	disconnectedCallback () {
	}

	render() {
		// Render HTML
		this.innerHTML =
			`<p>
				this is log in page
			</p>
            `;
	}

}

customElements.define('log-in', logIn);