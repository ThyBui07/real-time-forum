window.onload = function() {
  console.log('Page has finished loading');
  hello();
};

function hello() {
  console.log('Reload');
  getAllPosts();
}

async function getAllPosts() {
  try {
    const response = await fetch('/posts', {
      method: "GET",
    });
    if (response.ok) {
      const data = await response.json();
      console.log(typeof data);
      let existing = JSON.parse(localStorage.getItem('posts'));
      if (isEmpty(existing)) {
        console.log('existing is empty, null, or undefined');
        existing = data;
      } else {
        if (!areObjectsEqual(existing, data)) {
          console.log('existing is not equal to data');
          existing = data;
        } else {
          console.log('existing is equal to data')
          return
        }
      }
      localStorage.setItem('posts', JSON.stringify(existing));
    } else {
      throw new Error('HTTP status ' + response.status);
    }
  }
  catch (error) {
    console.log(error)
  }
}

function isEmpty(value) {
  return value === null || value === undefined;
}

function areObjectsEqual(obj1, obj2) {
  return JSON.stringify(obj1) === JSON.stringify(obj2);
}
