import BaseView from './View.js';

import ShopComponent from '../components/ShopComponent/ShopComponent.js';

// const { NetworkHandler } = window;

export default class SignupView{
    constructor(
        root,
        weapons
    ) {
        this.root = root;
        this.weapons = weapons;
        this.ShopComponent = new ShopComponent(this.root, this.weapons);

        document.addEventListener('click', (event) => {          
            if  (event.target.className === 'shop__menu-item' || event.target.parentElement.className === 'shop__menu-item') {
                if (event.target.dataset.section === undefined) {
                    this.weapon(event.target.parentElement.dataset.section);
                } else {
                    this.weapon(event.target.dataset.section);
                }
            }
        })
    }

    show() {
        this.ShopComponent.render();
    }

    weapon(weapon) {
        this.ShopComponent.renderWeaponInfo(weapon)
    }

    setWeapon(weaponId) {
        return weaponId;
    }

    hide() {
        this.root.innerHTML = '';
    }
}