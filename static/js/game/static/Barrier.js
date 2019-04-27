// import StaticEssence from './StaticEssence.js'

export default class Barrier {
    constructor(   
        xPos,
        yPos,

        xSize = 250,
        ySize = 250,
        URL = "/default_texture",    
        ) {
       
        // Координаты
        this.xPos = xPos;
        this.yPos = yPos;

        this.name = 'barrier';

        // Позиция прицела - только у плеера

        // Его размеры 
        this.xSize = xSize; // vh
        this.ySize = ySize; // vh

        this.centerX = this.xPos + this.xSize / 2;
        this.centerY = this.yPos + this.ySize / 2;
    }

    render(ctx) {
        ctx.beginPath();
        ctx.rect(this.xPos, this.yPos, this.xSize, this.ySize);
        ctx.fillStyle = "gray";
        ctx.fill();
        ctx.closePath();
    }

    logic() {}

    interact() {
    }

    // Геттеры

    get top() {
        return this.yPos;
    }

    get right() {
        return this.xPos + this.xSize;
    }

    get bottom() {
        return this.yPos + this.ySize;
    }

    get left() {
        return this.xPos;
    }

}