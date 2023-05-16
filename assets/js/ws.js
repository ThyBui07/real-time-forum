//console.log('ws')

window.onload = function () {
// Apply our listener functions to the submit event on both forms
// we do it this way to avoid redirects
// document.getElementById("chatroom-selection").onsubmit = changeChatRoom;
// document.getElementById("chatroom-message").onsubmit = sendMessage;

// Check if the browser supports WebSocket
    if (window["WebSocket"]) {
    //console.log("supports websockets");
     // Connect to websocket
     conn = new WebSocket("ws://" + document.location.host + "/ws");
    } else {
    alert("Not supporting websockets");
    }
}