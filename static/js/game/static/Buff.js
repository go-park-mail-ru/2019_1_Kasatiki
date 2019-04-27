import StaticEssence from './StaticEssence.js';

export default class Buff {
    constructor(
        xPos,
        yPos,

        xSize = 50,
        ySize = 50,

        cfg = {},
    ) {
        this.xPos = xPos;
        this.yPos = yPos;

        // Его размеры 
        this.xSize = xSize; // vh
        this.ySize = ySize; // vh

        this.cfg = cfg;
        this.visited = false;
    }

    render(ctx) {
        ctx.beginPath();
        ctx.rect(this.xPos, this.yPos, this.xSize, this.ySize);
        ctx.fillStyle = this.cfg.color;
        ctx.fill();
        ctx.closePath();
    }

    interact(obj) {
        if (this.cfg.name == 'heal') {
            obj.hp += this.cfg.value;
        } else if (this.cfg.name == 'boost') {
            obj.velocity += this.cfg.value;
        }
    }
}