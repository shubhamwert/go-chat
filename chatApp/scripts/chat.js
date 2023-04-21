console.log("Hello World    ")
let socket = new WebSocket("ws://127.0.0.1:8080/room/0/msg");
console.log("Attempting Connection...");

socket.onopen = () => {
    console.log("Successfully Connected");

};

socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
    socket.send("Client Closed!")
};

socket.onerror = error => {
    console.log("Socket Error: ", error);
};
socket.onmessage = (event) => {
    console.log("Message Recieved")
    msg = JSON.parse(event.data);
    document.getElementById("messagesArea").innerHTML +=
        "<div class='message'>" +
        "<div class='from'>" + msg.From + "</div>" +
        "<div class='text'>" + msg.message + "</div>" +

        "</div>"

}

function sendMessage() {
    let msg = document.getElementById("messageSend").value;
    console.log(msg)
    if (msg.length > 0) {

        msg = {
            'message': msg,
            'From': 'me'

        }
        socket.send(JSON.stringify(msg))
    }
}