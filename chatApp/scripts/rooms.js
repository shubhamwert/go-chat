
async function fetchRooms() {
    const response = await fetch('http://localhost:8080/');
    const rooms = await response.json()
    // waits until the request completes...
    console.log(rooms);
    return rooms
  }

fetchRooms().then(rooms => {
    console.log(rooms)
    console.log(document.getElementById("rooms").innerHTML)
    Object.keys(rooms).forEach(r => {
        document.getElementById("rooms").innerHTML +=
        "<div class='room' onclick='loadChat("+ r+")'>" +r+" "+
        rooms[r]
        "</div>"
    });
        


});
function loadChat(roomNo) {
    console.log(roomNo)
    universal.socketName=roomNo
    resetChatWindow()
    initializeConnection()
}
