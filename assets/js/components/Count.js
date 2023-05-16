import Observer from "../lib/Observer.js";
console.log('counnt.js')
class Count extends Observer {
  createMarkup(state) {
    console.log('state in count:', state)
    return `<span>
    user count: ${state.users.length}
    </span>`;
  }

  render(state, selector = "app") {
    const markup = this.createMarkup(state);
    const parent = document.getElementById(selector);

    parent.innerHTML = markup;
  }

  // This method will be called by the Subject(state) whenever it updates.
  // Notice how it prompts a re-render.
  update(state) {
    this.render(state, "user-count-container");
  }
}

export default Count;
