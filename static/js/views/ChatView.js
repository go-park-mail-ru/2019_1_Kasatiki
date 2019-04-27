import BaseView from './View.js';

import ChatComponent from '../components/ChatComponent/ChatComponent.js';

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

        console.log("ASDSA");

            let messageBox = document.createElement("div");

            messageBox.className = 'chat__chatbox-message'

            let messageText = document.createElement('div');
            messageText.className = 'chat__chatbox-message-text';
            messageText.innerText = 'HELLO';

            messageBox.appendChild(messageText);
    
        console.log(chatButton);
        chatBox.insertBefore(messageBox, chatButton.nextSibling);
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

        // chatButton.addEventListener('click', this.getMoreMessages);
        
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