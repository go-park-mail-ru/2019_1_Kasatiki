// // const ws = new WebSocket('https://advhater.ru/ws');

// ws.onopen = () => {
//     console.log('ws success connect');
    
//     const chatbox = document.querySelector('.chat__chatbox');

//     ws.onmessage = (evt) => {
//         // console.log('ws message:', message);

//         const messageBox = document.createElement("div");
//         messageBox.className = 'chat__chatbox-message'
//         messageBox.innerText = evt.data + '\n';

//         console.log(evt, evt.data);

//         chatbox.appendChild(messageBox);   
//     }
// }

export default class Ws {
    constructor(
        chatbox,
        wsUrl = '172.20.10.10:8080',
    ) {
        this.chatbox = chatbox;
        this.ws = new WebSocket('ws://' + wsUrl + '/ws');

        this.ws.onmessage = (evt) => {
            const messageBox = document.createElement("div");

            messageBox.className = 'chat__chatbox-message'
            messageBox.innerText = evt.data + '\n';

            console.log(evt, evt.data);

            this.chatbox.appendChild(messageBox);   
        }
    }

    setChatbox(cb) {
        this.chatbox = cb;
    }

    send(data) {
        this.ws.send(data);
    }
}

// export default ws;