console.log('test2.js')

let outputBox = document.querySelector('#output');

// window.addEventListener('DOMContentLoaded', (event) => {
//   outputBox.innerHTML = "initializing...";
//   tick();
// });

// function tick() {
//     console.log('tick')
//   fetch('/about')
//     .then((response) => {
//       if (!response.ok) {
//         throw new Error("error response");
//       }
//       return response.text();
//     })
//     .then((text) => {
//         console.log(text)
//       outputBox.innerHTML = text;
//     })
//     .catch((error) => {
//       outputBox.innerHTML = "network error";
//     });
// }