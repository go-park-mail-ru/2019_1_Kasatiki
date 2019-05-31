export default class Map {
    constructor(
        mapSize = 100,
        tileSize = 50,
    ) {
        this.mapSize = mapSize;
        this.tileSize = tileSize;
        this.baseTileSize = tileSize;
        
        this.map = [];

        this.canvas = document.querySelector('.game');
        this.ctx = this.canvas.getContext('2d');
    }

    setCanvasWH(w, h) {
        this.canvas.width = this.tileSize * w;
        this.canvas.height = this.tileSize * h;
    }

    _createMap() {
        for (let i = 0; i < this.mapSize; i++) {
            for (let j = 0 ; j < this.mapSize; j++) {
                if (i == 0 || i == this.mapSize-1 || j == 0 || j == this.mapSize -1) {
                    this.map[this.mapSize*i + j] = 1;
                } else {
                    this.map[this.mapSize*i + j] = Math.floor( Math.random() * 1.999999 );
                }
            }
        }
    }


    _drawBarriers() {
        for (let i = 0; i < barriers.length; i++) {
            this.ctx.strokeStyle = '#000000';

            // if (barriers[i].object.x > viewport.x && barriers[i].object.y > viewport.y && barriers[i].object.x < (viewport.x + viewport.w) && barriers[i].object.y < (viewport.y + viewport.h)) {
            this.ctx.strokeRect(barriers[i].object.x - viewport.x, barriers[i].object.y - viewport.y, barriers[i].object.xsize, barriers[i].object.ysize);
            // }

            // console.log(barriers[i].object.x, barriers[i].object.y, barriers[i].object.xsize, barriers[i].object.ysize);
        }
    }


    _drawTile(tile, x, y) {
        let color;
        switch (tile) {
            case 0:
                color = '#979797';
                break;
            case 1:
                color = '#5B5B5B';
                break;
        }

        this.ctx.fillStyle = color;
        this.ctx.fillRect(x, y, this.tileSize, this.tileSize);
    }

    drawMap(viewport, tileSize, mapSize) {
        var x_min = Math.floor((viewport.x) / tileSize.val * tileSize.val / 50);
        var y_min = Math.floor((viewport.y) / tileSize.val * tileSize.val / 50);
        var x_max = Math.ceil((viewport.x + viewport.w) / tileSize.val * tileSize.val / 50);
        var y_max = Math.ceil((viewport.y + viewport.h) / tileSize.val * tileSize.val / 50);
    
        if (x_min < 0) {
            x_min = 0;
            // this.viewport.x = 0;
            // x_max = this.viewport.w / tileSize.val;
        }
        if (y_min < 0) {
            y_min = 0;
            // this.viewport.y = 0;
            // y_max = this.viewport.w / tileSize.val;
        }
        if (x_max > mapSize.val) {
            x_max = mapSize.val;
        }
        if (y_max > mapSize.val) {
            y_max = mapSize.val;
        }
    
        // console.log(x_min, y_min, x_max, y_max,
        //     'vp:', viewport.x, viewport.y, viewport.w, viewport.h,
        //     'ts', tileSize.val);
    
        for (let i = x_min; i < x_max; i++) {
            for (let j = y_min; j < y_max; j++) {
                let tile_x = Math.floor(i * tileSize.val - viewport.x);
                let tile_y = Math.floor(j * tileSize.val - viewport.y);
    
                this._drawTile(this.map[j][i], tile_x, tile_y);
            }
        }
    }
}