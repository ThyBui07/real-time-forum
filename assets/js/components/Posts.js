import Observer from "../lib/Observer.js";
console.log('posts.js')
class Posts extends Observer {
  createMarkup(state) {
    const posts = JSON.parse(state.posts);
    return posts.map(post => `
    <div class="card flex-md-row mb-4">
      <div class="card-body d-flex flex-column align-items-start">
        <h3 class="mb-0">
          <a class="text-dark" href="#">${post.Title}</a>
        </h3>
        <div class="mb-1 text-muted">${post.Date} by ${post.AuthorName}</div>
        <p class="card-text mb-2">${post.Content}</p>
        <div class="media text-muted pt-3">
          <img src="../../img/avatar.jpeg" alt="" class="mr-2 rounded" style="width: 32px; height: 32px;">
          <p class="media-body pb-3 mb-0 small lh-125 bg-light p-2">
            <strong class="d-block text-gray-dark">@username</strong>
            Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus.
          </p>
        </div>

        <div class="media text-muted pt-3">
          <img src="../../img/avatar.jpeg" alt="" class="mr-2 rounded" style="width: 32px; height: 32px;">
          <p class="media-body pb-3 mb-0 small lh-125 bg-light p-2">
            <strong class="d-block text-gray-dark">@username</strong>
            Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus.
          </p>
        </div>
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
