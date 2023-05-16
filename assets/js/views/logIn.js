class logIn extends HTMLElement {
  constructor() {
    super();
  }

 async signIn(event) {
    event.preventDefault();
    const email = document.getElementById("inputEmail").value;
    const password = document.getElementById("inputPassword").value;
    try {
      const response = await fetch("/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: email,
          password: password
        }),
      });

      if (response.ok) {
        const data = await response.json();
        console.log(data);
        localStorage.setItem('userInfo', JSON.stringify(data.User))
        localStorage.setItem('posts', JSON.stringify(data.Posts))

        // window.globalPosts = data.Posts;
        // console.log(window.globalPosts)
        window.history.pushState({}, "","/");
        window.location.hash = '/';
        //window.history.pushState(stateObj, "", "/homepage");
      } else {
        throw new Error('HTTP status ' + response.status);
      }
    } catch (error) {
      console.log(error);
    }
  }

  connectedCallback() {
    this.render();
    this.addEventListener("submit", this.signIn);
  }

  disconnectedCallback() {}

  render() {
    this.innerHTML = `
    <div class="login-body">
        <form class="form-signin">
        <h1 class="h3 mb-3 fs-2 fw-bold text-center">Sign In</h1>
        <label for="inputEmail" class="sr-only">Email address</label>
        <input type="email" id="inputEmail" class="form-control" placeholder="Email address" required autofocus>
        <label for="inputPassword" class="sr-only">Password</label>
        <input type="password" id="inputPassword" class="form-control" placeholder="Password" required>
        <button class="btn btn-lg btn-outline-dark btn-block" type="submit">Sign in</button>
        <p class="mt-2 text-center">Don't have an account? <a href="/signup">Sign up</a></p>
        </form>
  </div>
            `;
  }
}

customElements.define("log-in", logIn);
