import BaseView from './View.js';

import ChatComponent from '../components/ChatComponent/ChatComponent.js';
// import NetworkHandler from '../modules/NetworkHandler.js';

const { NetworkHandler } = window;

/**
 * Класс с отрисовкой формы логина.
 */
export default class ChatView extends BaseView {
    constructor() {
        super(...arguments);
        this.ChatComponent = new ChatComponent();
    }

    set DOMelement(DOMelement) {
        this.root = DOMelement;
    }

    getMoreMessages() {
        let chatBox = document.querySelector('.chat__chatbox');
        let chatButton = document.querySelector('#chat__get-more');

        const that = this;

        NetworkHandler.doGet({
			callback(data) {
                data.forEach( mess => {
                    const messageBox = document.createElement("div");

                    messageBox.className = 'chat__chatbox-message'
                    let message = JSON.parse(mess);
        
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
        
                    chatBox.insertBefore(messageBox, chatButton.nextSibling); 
                })
			},
			path: '/chat/limit?offset=20',
		});
    }

    show(isPage) {
        if (isPage === false) {
            this.root.innerHTML = this.ChatComponent.render(false);
        } else {
            this.root.innerHTML = this.ChatComponent.render(true);
        }
        
        const chatForm = document.querySelector('.chat__form');
        const chatInput = document.querySelector('.chat__input');
        let chatButton = document.querySelector('#chat__get-more');

        chatButton.addEventListener('click', this.getMoreMessages);
        
        console.log(chatForm);
        this.router.ws.setChatbox(document.querySelector('.chat__chatbox'));

        let that = this;
        chatForm.addEventListener('click', () => {
            let message = chatInput.value;

            if (message !== '') {
                that.router.ws.send(message);

                chatInput.value = '';
            }
        })
    }
}