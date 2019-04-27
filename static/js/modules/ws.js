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
        this.ws = new WebSocket('wss://' + wsUrl + '/ws');

        this.ws.onmessage = (evt) => {
            const messageBox = document.createElement("div");

            messageBox.className = 'chat__chatbox-message'
            let message = JSON.parse(evt.data);

            let messageAvatar = document.createElement('img');
            messageAvatar.className = 'chat__chatbox-message-avatar';
            messageAvatar.src = message.Url;

            let messageNickname = document.createElement('div');
            messageNickname.className = 'chat__chatbox-message-nickname';
            messageNickname.innerText = message.Nickname + ':';

            let messageText = document.createElement('div');
            messageText.className = 'chat__chatbox-message-text';
            messageText.innerText = message.Body;

            let messageTimestamp = document.createElement('div');
            messageTimestamp.className = 'chat__chatbox-message-timestamp';
            messageTimestamp.innerText = message.Timestamp;

            messageBox.appendChild(messageAvatar);
            messageBox.appendChild(messageNickname);
            messageBox.appendChild(messageText);
            messageBox.appendChild(messageTimestamp);

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