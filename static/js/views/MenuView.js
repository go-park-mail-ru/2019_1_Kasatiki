import BaseView from './View.js';

import MenuComponent from '../components/MenuComponent/MenuComponent.js';
import ws from '../modules/ws.js';

const { NetworkHandler } = window;

export default class MenuView extends BaseView {
    constructor() {
        super(...arguments);
        this.MenuComponent = new MenuComponent();
    }

    show() {
        let that = this;
        NetworkHandler.doGet({
            // eslint-disable-next-line no-unused-vars
            callback(data) {
                console.log('menu view', data);
                if (typeof(data) == 'object') {
                    that.root.innerHTML = that.MenuComponent.render(data);

                    const profileSection = document.querySelector('.menu__profile');
                    const buttonsSection = document.getElementById('menu__profile-buttons-section');
                    profileSection.addEventListener('mouseover', showButtons, false);
                    profileSection.addEventListener('mouseout', hideButtons, false);

                    // const sendButton = document.querySelector('.chat__submit');
                    const chatForm = document.querySelector('.chat__form');
                    const chatInput = document.querySelector('.chat__input');
                    console.log(sendButton, chatInput);
                    chatForm.onsubmit('click', () => {
                        let message = chatInput.value;

                        ws.send(message);
                    })
        
                    function showButtons(e) {
                        // console.log('on');
                        console.log();
                        buttonsSection.style.display = 'flex';
                    }

                    function hideButtons(e) {
                        buttonsSection.style.display = 'none';
                    }

                } else {
                    that.root.innerHTML = that.MenuComponent.render(false);
                }   


            },
            path: '/api/isauth',
        });
    }
}