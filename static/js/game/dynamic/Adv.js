import DynamicEssence from "./DynamicEssence.js";

export default class Adv extends DynamicEssence {
    constructor() {
        super(...arguments);

        this.xPrev = this.xPos;
        this.yPrev = this.yPos;

        this.curTargetX;
        this.curTargetY;

        this.centerX;
        this.centerY;

        this.name = 'adv';

        this.advUrl = 'http://ya.ru';

        this.teta;
        this.deltaColor = null;

        this.color = 9400;
    }
            // // Геттеры

            // get top() {
            //     return this.yPos;
            // }
        
            // get right() {
            //     return this.xPos + this.xSize;
            // }
        
            // get bottom() {
            //     return this.yPos + this.ySize;
            // }
        
            // get left() {
            //     return this.xPos;
            // }

    render(ctx) {
        ctx.beginPath();
        ctx.rect(this.xPos, this.yPos, this.xSize, this.ySize);
        ctx.fillStyle = '#ff' + this.color;
        ctx.fill();
        ctx.closePath();
    }

    logic(curTargetX, curTargetY) {
        this.centerY = this.yPos + this.ySize/2;
        this.centerX = this.xPos + this.xSize/2;

        this.curTargetX = curTargetX;
        this.curTargetY = curTargetY;

        this.teta = Math.atan2(this.curTargetX - this.xPos, this.curTargetY - this.yPos);

        this.xPrev = this.xPos;
        this.yPrev = this.yPos;

        this.xPos += this.velocity * Math.sin(this.teta);
        this.yPos += this.velocity * Math.cos(this.teta);
    }

    interact(obj) {
        if (obj.name == 'bullet') {
            if (this.deltaColor === null) {
                this.deltaColor = Math.floor(this.color * obj.damage / this.hp);
                console.log('damage:',obj.damage,'delta:', this.deltaColor)
            }
            this.hp -= obj.damage;
            this.xPos = this.xPrev;
            this.yPos = this.yPrev;
            this.color -= this.deltaColor;
            return this.hp;
        } else if (obj.name == 'player') {
            this.hp = 0;
            obj.hp -= 25;
            window.open(this.advUrl);
            return this.hp;
        } else if (obj.name == 'barrier') {
            // this.xPos = this.xPrev;
            // this.yPos = this.yPrev;

            // prev

            if (this.curTargetX - this.center.x > 0) {
                if (this.right >= obj.left) {
                    // this.right = obj.left;
                    this.xPos -= this.velocity;
                } 
            } else if (this.curTargetX - this.center.x < 0) {
                if (this.left >= obj.right) {
                    // this.left = obj.right;
                    this.xPos += this.velocity;
                }
            }

            if (this.curTargetY - this.center.y > 0) {
                if (this.bottom >= obj.top) {
                    // this.bottom = obj.top;
                    this.yPos -= this.velocity;
                }
            } else if (this.curTargetY - this.center.y < 0) {
                if (this.top <= obj.bottom) {
                    // this.top = obj.bottom;
                    this.yPos += this.velocity;
                }
            }

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

        }
    }
}