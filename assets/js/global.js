window.onload = function() {
  console.log('Page has finished loading');
  hello();
};

function hello() {
  console.log('Reload');
}

async function getAllPosts() {
  try {
    const response = await fetch('/posts', {
      method: "GET",
    });
    if (response.ok) {
      const data = await response.json();
      console.log(data);
    } else {
      throw new Error('HTTP status ' + response.status);
    }
  }
  catch (error) {
    console.log(error)
  }
}