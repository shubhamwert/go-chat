console.log("Hello World")
var universal =
{
    socketName: "0"
};

let socket = null
function initializeConnection() {
    resetChatWindow()
    if (socket != null) {
        socket.close()

    }
    let data = {
        "username": document.getElementById("username").value,
        "password": document.getElementById("password").value
    }

    socket = new WebSocket("ws://127.0.0.1:8080/room/" + universal.socketName + "/connect");
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
        console.log(JSON.stringify(data))
        socket.send(JSON.stringify(data))

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
        msgs = JSON.parse(event.data);
        msgs.forEach(msg => {
            if (msg.From == data["username"]) {
                document.getElementById("messagesArea").innerHTML +=

                "<div class='messageSelf'>" +
                "<div class='from'>" + msg.From + "</div>" +
                "<div class='text'>" + msg.message + "</div>" +

                "</div>"

            } else {
                document.getElementById("messagesArea").innerHTML +=
                    "<div class='message'>" +
                    "<div class='from'>" + msg.From + "</div>" +
                    "<div class='text'>" + msg.message + "</div>" +

                    "</div>"
            }

        });

    }
    return socket
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
        document.getElementById("messageSend").value = ""
    }
}
function resetChatWindow() {
    document.getElementById("messagesArea").innerHTML = ""

}