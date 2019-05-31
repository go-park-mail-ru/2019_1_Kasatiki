export default class Player {
    constructor(
        x, y, ctx
    ) {
        this.x = x;
        this.y = y;

        this.id = 1;

        this.v = 5;

        this.c = 0.2;
        this.ctx = ctx;
    }

    draw(viewport, tileSize) {
        this.ctx.fillStyle = '#45FF70';
        this.ctx.fillRect(Math.round(this.x - viewport.x - viewport.zoom), Math.round(this.y - viewport.y - viewport.zoom), tileSize.val, tileSize.val);
    }

    // move(keyMap) {
    //     if (keyMap.up) {
    //         this.y -= this.v;
    //     }
    //     if (keyMap.down) {
    //         this.y += this.v;
    //     }
    //     if (keyMap.left) {
    //         this.x -= this.v;
    //     }
    //     if (keyMap.right) {
    //         this.x += this.v;
    //     }
    // }
}