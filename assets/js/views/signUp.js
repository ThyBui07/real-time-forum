// Extend the HTMLElement class to create the web component
class signUp extends HTMLElement {
    /**
     * The class constructor object
     */
    constructor() {
      // Always call super first in constructor
      super();
    }
    count() {
      console.log("counting");
    }
    /**
     * Runs each time the element is appended to or moved in the DOM
     */
    connectedCallback() {
      this.render();
      this.count();
    }
    /**
     * Runs when the element is removed from the DOM
     */
    disconnectedCallback() {}

    render() {
      // Render HTML
	// Nickname,Age,Gender,First Name,Last Name,E-mail,Password
      this.innerHTML = `
      <div class="login-body">
          <form class="form-signin">
          <h1 class="h3 mb-3 fs-2 fw-bold text-center">Sign Up</h1>
          <label for="inputEmail" class="sr-only">Email address</label>
          <input type="email" id="inputEmail" class="form-control" placeholder="Email address" required autofocus>
          <label for="inputPassword" class="sr-only">Password</label>
          <input type="password" id="inputPassword" class="form-control" placeholder="Password" required>
          <button class="btn btn-lg btn-outline-dark btn-block" type="submit">Sign up</button>
            <p class="mt-2 text-center">Already have an account? <a href="/">Log in</a></p>
          </form>
    </div>
              `;
    }
  }
  customElements.define("sign-up", signUp);