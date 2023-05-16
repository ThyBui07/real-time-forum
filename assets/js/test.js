// window.onpopstate = () => {
//     console.log('popstate');
//     console.log("Page path is: " + window.location.pathname)
// }
// const APP = {
//     init() {
//         console.log('init');
//         APP.checkState();
//     },
//     checkState() {
//         console.log(location);
//         console.log("Page path is: " + window.location.pathname)
//         if (window.location.pathname == "/about") {
//             console.log('about')
//             tick();
//         }
//     }
// }


//let outputBox = document.querySelector('#output');

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

console.log('test.js')