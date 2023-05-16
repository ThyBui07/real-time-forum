import State from "../lib/State.js";
import Count from "../components/Count.js";
import Posts from "../components/Posts.js";
class homePage extends HTMLElement {
	constructor () {
		let posts = localStorage.getItem('posts');
		console.log('posts in homepage:', posts)
		// Always call super first in constructor
		super();
		// Instantiate classes.
		this.AppState = new State();
		this.AppState.update({ posts })
		this.posts = new Posts();
		this.AppState.addObserver(this.posts);
	}
	// addPost() {
	// 	var posts = localStorage.getItem('posts');
	// 	posts = JSON.parse(posts);
	// 	console.log(posts);
	// 	posts.push({
	// 		"ID": 3,
	// 		"AuthorID": 1,
	// 		"Title": "My Whatver Post",
	// 		"Content": "Hello, !",
	// 		"Date": "2023-05-27",
	// 		"Categories": "Programming, Golang",
	// 		"CommentCount": 0
	// 	})
	// 	console.log(posts);
	// 	localStorage.setItem('posts', JSON.stringify(posts));
	// }
	connectedCallback () {
		//console.log(window.globalPosts)
		this.render();
		this.posts.render(this.AppState.get(), "posts-container");
		//this.renderPosts();
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
				This is the home page
			</p>
			<div id="welcome"></div>
			<div id="posts"></div>
			<div id="user-count-container"></div>
			<div id="posts-container"></div>
			`;
	}

}


customElements.define('home-page', homePage);