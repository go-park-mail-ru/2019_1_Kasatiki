export default class Viewport {
    constructor(
        x = 0, y = 0, w = 800, h = 800, // начальные и конечные координаты порта
        gameScreenW = 16, // ширина порта в тайлах
        gameScreenH = 16, // высота порта в тайлах
    ) {
        this.x = x;
        this.y = y;
        this.w = w;
        this.h = h;

        this.gameScreenH = gameScreenH;
        this.gameScreenW = gameScreenW;

        this.scale = 16;
        this.zoom = 0;
    }

    update(x ,y, zoom, tileSize) {
        if (zoom) {
            if (tileSize.val > 25) {
                this.scale += 2;
                tileSize.val = 800 / this.scale;
                this.zoom += 50;
                this.w = this.w + this.zoom;
                this.h = this.h + this.zoom;
            }
        } else if (!zoom) {
            if (tileSize.val < 50) {
                this.scale -= 2;
                tileSize.val = 800 / this.scale;
                this.zoom -= 50;
                this.w = this.w - this.zoom;
                this.h = this.h - this.zoom;
            }
        }

        // this.x += (x - this.x - this.w * 0.5) * 0.05;
        // this.y += (y - this.y - this.h * 0.5) * 0.05;

        this.x = x - this.w / 2;
        this.y = y - this.h / 2;
    }
}