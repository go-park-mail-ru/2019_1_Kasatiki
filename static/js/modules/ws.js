const ws = new WebSocket('ws://172.20.10.10:8080/ws');
// const ws = new WebSocket('https://advhater.ru/ws');

ws.onopen = () => {
    console.log('ws success connect');
    

    ws.onmessage = (evt) => {
        // console.log('ws message:', message);

        const messageBox = document.createElement("div");
        messageBox.className = 'chat__chatbox-message'
        messageBox.innerText = evt.data + '\n';

        console.log(evt, evt.data);
        const chatbox = document.querySelector('.chat__chatbox');


        chatbox.appendChild(messageBox);

        // ws.onclose = function (evt) {
        //     const messageBox = document.createElement("div.chat__chatbox-message");
        //     messageBox.innerHTML = "<b>Connection closed.</b>";
            
        //     chatbox.appendChild(messageBox);
        // };
        // ws.onmessage = function (evt) {
        //     const messageBox = document.createElement("div.chat__chatbox-message");
        //     messageBox.innerText = evt.data;

        //     console.log(evt, evt.data);

        //     chatbox.appendChild(messageBox);
        // };
    }
}

export default ws;