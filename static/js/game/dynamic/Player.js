// import KeyboardControl from '../functions/KeyboardControl.js'

import Bullet from './bullet.js'
import DynamicEssence from "./DynamicEssence.js";

export default class Player extends DynamicEssence {
    
    constructor(
    ) {
        super(...arguments);

        this.name = 'player'

        this.defaultVelocity = this.velocity;
        this.buffs = [];

        this.centerX;
        this.centerY;

        this.score;

        this.inventory = {
            weapons : [],
            buffs : [],
        };

        this.inShop;
        this.currentShop;

        this.xPrev = this.xPos;
        this.yPrev = this.yPos;

        this.weapon = {
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
        };
    }

    render(ctx) {
        ctx.beginPath();
        ctx.rect(this.xPos, this.yPos, this.xSize, this.ySize);
        ctx.fillStyle = "#48F67F";
        ctx.fill();
        ctx.closePath();
    }

    logic(eventsMap) {
        const that = this;
        this._logicBuffs();

        this.centerY = this.yPos + this.ySize/2;
        this.centerX = this.xPos + this.xSize/2;
        // console.log(eventsMap['mouseX'], eventsMap['mouseY'])

        // this.teta = this.myMath.getTeta(this.centerX, this.centerY, eventsMap['mouseX'], eventsMap['mouseY']);
        // console.log('this.teta',this.teta);

        if(eventsMap['right']) {
            this.xPrev = this.xPos;
            this.vx = this.velocity;
            this.xPos += this.vx;
        } else if(eventsMap['left']) {
            this.xPrev = this.xPos;
            this.vx = - this.velocity;
            this.xPos += this.vx;
        } else if (!eventsMap['right'] || !eventsMap['left']) {
            this.vx = 0;
        }

        if(eventsMap['up']) {
            this.yPrev = this.yPos;
            this.vy = - this.velocity;
            this.yPos += this.vy;
        } else if(eventsMap['down']) {
            this.yPrev = this.yPos;
            this.vy = this.velocity;
            this.yPos += this.vy;
        } else if (!eventsMap['up'] || !eventsMap['down']) {
            this.vy = 0;
        }

        // console.log('vx ', this.vx, 'vy', this.vy);

        if (this.inShop) {
            // console.log(this.currentShop);
            if (eventsMap['interact']) {
                // console.log('open shop');
                this.currentShop.open();
            } else {
                eventsMap['interact'] = false;
                this.currentShop.close();
            }
        }

        // console.log(eventsMap);


        // this.vx = 0; this.vy = 0;

        // if (this.xPos <= 0 || this.xPos >= cvsWidth || this.yPos <= 0 || this.yPos >= cvsHeight) {
        //     this.interact();
        // }
    }

    interact(obj = {}) {
        // console.log(name);
        if (obj.name == 'barrier') {

            // if (this.right > obj.left) {   
            //     this.right = obj.left;
            // } else if (this.left < obj.right) {
            //     this.left = obj.right;
            // }

            // if (this.bottom > obj.top) {
            //     this.bottom = obj.top;
            // } else if (this.top < obj.bottom) {
            //     this.top = obj.bottom;
            // }

            // if (this.vx > 0) {
            //     if (this.right > obj.left) {   
            //         this.right = obj.left;
            //     }
            // } else if (this.vx < 0) {
            //     if (this.left < obj.right) {
            //         this.left = obj.right;
            //     }
            // } 
            
            // if (this.vy > 0) {
            //     if (this.bottom > obj.top) {
            //         this.bottom = obj.top;
            //     } 
            // } else if (this.vy < 0) {
            //     if (this.top < obj.bottom) {
            //         this.top = obj.bottom;
            //     }
            // }


            // if (this.left < obj.right) {
            //     this.xPos = this.xPrev + 1;
            // } else if (this.right > obj.left) {
            //     this.xPos = this.xPrev - 1;
            // } else if (this.top < obj.bottom) {
            //     this.yPos = this.yPrev + 1;
            // } else if (this.bottom > obj.top) {
            //     this.yPos = this.yPrev - 1;
            // }

            let vecX = this.centerX - obj.centerX;
            let vecY = this.centerY - obj.centerY;

            if (vecY * vecY > vecX * vecX) {
                if (vecY < 0) {
                    this.bottom = obj.top;
                } else if (vecY > 0){
                    this.top = obj.bottom;
                } 
            } else if (vecY * vecY < vecX * vecX) {
                if (vecX > 0) {
                    this.left = obj.right;
                } else if (vecX < 0) {
                    this.right = obj.left;
                }
            }

            // this.xPos = this.xPrev;
            // this.yPos = this.yPrev;
        } else if (obj.name == 'adv') { 
            this.hp -= 50;
        } else if (obj.name == 'shop') {
            this.inShop = true;
            this.currentShop = obj;
        } 
    }

    _addHp(hp) {
        if (this.hp + hp >= this.hpCapacity) {
            this.hp = this.hpCapacity;
        } else {
            this.hp += hp;
        }
    }

    _logicBuffs() {
        let buffs = this.buffs;
        this.buffs = [];
        // // console.log(buffs);
        buffs.forEach((buff) => {
            if (Date.now() - buff.startTime < buff.buff.time) {
                this.buffs[this.buffs.length] = buff;
            } else {
                if (buff.buff.name == 'increaseHpCapacity') {
                    this.hpCapacity = buff.buff.value;
                    if (this.hp > this.hpCapacity) {
                        this.hp = this.hpCapacity;
                    }
                } else if (buff.buff.name == 'increaseVelocity') {
                    this.velocity -= buff.buff.value;
                }
            }
        });
    }

    addBuff(person, buff) {
        if (buff.isTemporary) {
            this.buffs[this.buffs.length] = {
                buff: buff,
                startTime: Date.now(),
            }
            switch (buff.name) {
                case 'increaseHpCapacity':
                    // eslint-disable-next-line no-case-declarations
                    let prevHpCapacity = this.hpCapacity;
                    person.hpCapacity = this.hpCapacity;
                    person.hp *= (1 + this.hpCapacity / prevHpCapacity);
                    continue;
                case 'increaseVelocity':
                    person.velocity += buff.value;
                    // console.log(person.velocity);
                    continue;
            }
        } else {
            switch (buff.name) {
                case 'health':
                    this._addHp(buff.value);
                    continue;
            }
        }
    }
}