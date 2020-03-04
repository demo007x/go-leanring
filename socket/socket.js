let socket = new WebSocket("ws://localhost:7777")

socket.onclose = function () {
    console.log("socket closed")
}

socket.onopen = function () {
    console.log("opend")
}

socket.onerror = function (error) {
    console.error(error)
}

socket.onmessage = function (msg) {
    console.log("receive message:", msg);
}

socket.send("hello")