export default class Viewport {
    constructor(
        x = 0, y = 0, w = 800, h = 800, // начальные и конечные координаты порта
        gameScreenW = 16, // ширина порта в тайлах
        gameScreenH = 16, // высота порта в тайлах
        tileSize = 50,
    ) {
        this.x = x;
        this.y = y;
        this.w = w;
        this.h = h;

        this.gameScreenH = gameScreenH;
        this.gameScreenW = gameScreenW;

        this.scale = 16;
        this.zoom = 0;

        this.tileSize = tileSize;
        this.baseTileSize = tileSize;
    }

    update(x ,y, zoom) {
        if (zoom) {
            if (this.tileSize > 25) {
                this.scale += 2;
                this.tileSize = Math.ceil(800 / this.scale);
                this.zoom += 50;
                this.w += this.baseTileSize * 2;
                this.h += this.baseTileSize * 2;
            }
        } else if (!zoom) {
            if (this.tileSize < 50) {
                this.scale -= 2;
                this.tileSize = Math.ceil(800 / this.scale)
                this.zoom -= 50;
                this.w -= this.baseTileSize * 2;
                this.h -= this.baseTileSize * 2;
            }
        }

        // this.x += (x - this.x - this.w * 0.5) * 0.05;
        // this.y += (y - this.y - this.h * 0.5) * 0.05;

        this.x = x - this.w / 2;
        this.y = y - this.h / 2;
    }
}