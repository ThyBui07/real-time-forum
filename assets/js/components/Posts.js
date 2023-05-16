import Observer from "../lib/Observer.js";
console.log('posts.js')
class Posts extends Observer {
  createMarkup(state) {
    const posts = JSON.parse(state.posts);
    return posts.map(post => `
    <div class="card flex-md-row mb-4 h-md-250">
      <div class="card-body d-flex flex-column align-items-start">
        <h3 class="mb-0">
          <a class="text-dark" href="#">${post.Title}</a>
        </h3>
        <div class="mb-1 text-muted">${post.Date} by ${post.AuthorName}</div>
        <p class="card-text mb-auto">${post.Content}</p>
        <a href="#">Continue reading</a>
      </div>
    </div>
  `).join('')
  }

  render(state, selector = "app") {
    const markup = this.createMarkup(state);
    const parent = document.getElementById(selector);

    parent.innerHTML = markup;
  }

  // This method will be called by the Subject(state) whenever it updates.
  // Notice how it prompts a re-render.
  update(state) {
    this.render(state, "posts-container");
  }
}

export default Posts;
