import StaticEssence from './StaticEssence.js';
import ShopComponent from '../../components/ShopComponent/ShopComponent.js'
import ShopView from '../../views/ShopView.js'


export default class Shop extends StaticEssence{
    constructor() {
        super(...arguments)

        this.playerInShop = false;
        this.shopOpenStatus = false;

        this.ctx;
        this.player;

        this.root = document.body;

        this.name = 'shop';
        this.shopDOM = new ShopComponent();

        addEventListener('click', (event) => {
            if  (event.target.className === 'weapon__about-main-purchase') {
                this.setWeapon(event.target.dataset.section);
            }
        });

        this.weapons = [
            {
                id : 0,
                name : 'Revolver',
                icon : "../../../icons/revolver.svg",
                cost : 200,
                fireRate : 500,
                damage : 25,
                velocity : 5,
                bulletSize : 2,
                bulletColor : 'red',
                about : 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam quis tempus magna. Nunc eget porttitor turpis. Sed sagittis lacus vel ligula vehicula, id rhoncus ipsum gravida.',
            },
            {
                id : 1,
                name : 'UZI',
                icon : "../../../icons/uzi.svg",
                cost : 1000,
                fireRate : 80,
                damage : 10,
                velocity : 10,
                bulletSize : 2,
                bulletColor : 'red',
                about : 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam quis tempus magna. Nunc eget porttitor turpis. Sed sagittis lacus vel ligula vehicula, id rhoncus ipsum gravida.',
            },
            {
                id : 2,
                name : 'Ak-74',
                icon : "../../../icons/ak-74.svg",
                cost : 1200,
                fireRate : 200,
                damage : 27,
                velocity : 10,
                bulletSize : 2,
                bulletColor : 'red',
                about : 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam quis tempus magna. Nunc eget porttitor turpis. Sed sagittis lacus vel ligula vehicula, id rhoncus ipsum gravida.',
            },
            {
                id : 3,
                name : 'M-16',
                icon : "../../../icons/m-16.svg",
                cost : 1150,
                fireRate : 180,
                damage : 22,
                velocity : 8,
                bulletSize : 2,
                bulletColor : 'red',
                about : 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam quis tempus magna. Nunc eget porttitor turpis. Sed sagittis lacus vel ligula vehicula, id rhoncus ipsum gravida.',
            },
            {
                id : 4,
                name : 'Shotgun',
                icon : "../../../icons/shotgun.svg",
                cost : 900,
                fireRate : 1000,
                damage : 7,
                velocity : 13,
                bulletSize : 2,
                bulletColor : 'red',
                about : 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam quis tempus magna. Nunc eget porttitor turpis. Sed sagittis lacus vel ligula vehicula, id rhoncus ipsum gravida.',
            },
            {
                id : 5,
                name : 'RPG',
                icon : "../../../icons/rpg.svg",
                cost : 10000,
                fireRate : 2000,
                damage : 100,
                velocity : 15,
                bulletSize : 30,
                bulletColor : 'red',
                about : 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam quis tempus magna. Nunc eget porttitor turpis. Sed sagittis lacus vel ligula vehicula, id rhoncus ipsum gravida.',
            },
            {
                id : 6,
                name : 'Laser',
                icon : "../../../icons/laser.svg",
                cost : 20000,
                fireRate : 0,
                damage : 100,
                velocity : 20,
                bulletSize : 20,
                bulletColor : 'aqua',
                about : 'FUCKING LASER! HOW DO YOU LIKE THIS ELON MUSK?!',
            },
        ]

        if (document.querySelector('.shop') === null) {
            this.shop = document.createElement('div');
            this.shop.className = 'shop';
            document.body.appendChild(this.shop);
        } else {
            this.shop = document.querySelector('.shop');
        }

        this.shopView = new ShopView(this.shop, this.weapons);
    }

    render(ctx) {
        this.ctx = ctx;
        ctx.beginPath();
        ctx.rect(this.xPos, this.yPos, this.xSize, this.ySize);
        ctx.fillStyle = "#C733FF";
        ctx.fill();
        ctx.closePath();
        
        if (this.playerInShop && !this.shopOpenStatus) {
            ctx.fillStyle = "#000";
            ctx.font = "italic 20pt Arial";
            ctx.fillText('Press E to Shop', 600, 300);
        }
    }

    logic() {

    }

    interact(player) {
        this.player = player;
        this.playerInShop = true;
    }

    setWeapon(weaponId) {
        const that = this;
        if (this.player !== undefined) {
            console.log(this.player);
            if (that.player.score >= this.weapons[weaponId].cost) {
                that.player.weapon = this.weapons[weaponId];
                that.player.score -= this.weapons[weaponId].cost;
            }
        }
    }

    open() {
        if (this.playerInShop) {
            this.shopOpenStatus = true;

            this.ctx.fillStyle = 'E3E3E3';
            this.ctx.fillRect(250, 300, 650, 600);

            if (this.shop.innerHTML === '') {
                this.shopView.show();
            }
        }
    }

    close() {
        this.shopOpenStatus = false;
        // console.log('Shop closed');
        this.player = undefined;
        this.shopView.hide();
    }
}