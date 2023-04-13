// Extend the HTMLElement class to create the web component
class BigGreeting extends HTMLElement {

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
		//console.log('connected!', this);
		this.render();
	}

	/**
	 * Runs when the element is removed from the DOM
	 */
	disconnectedCallback () {
		console.log('disconnected', this);
	}

	render() {
		// Render HTML
		this.innerHTML =
			`<p>
				<button>Hi there again, this is big greetingn page!</button>
                <greeting-message></greeting-message>
			</p>`;
	}

}

customElements.define('big-greeting', BigGreeting);