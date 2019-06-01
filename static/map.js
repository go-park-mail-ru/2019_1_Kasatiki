export default class Map {
    constructor(
        mapSize = 100,
    ) {
        this.mapSize = mapSize;
        
        this.map = [];

        this.canvas = document.querySelector('.game');
        this.ctx = this.canvas.getContext('2d');
    }

    setCanvasWH(w, h) {
        this.canvas.width = w;
        this.canvas.height = h;
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
            this.ctx.strokeRect(barriers[i].object.x - viewport.x, barriers[i].object.y - viewport.y, barriers[i].object.xsize, barriers[i].object.ysize);
        }
    }


    _drawTile(tile, x, y, tileSize) {
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
        this.ctx.fillRect(x, y, tileSize, tileSize);
    }

    drawMap(viewport) {
        let x_min = Math.floor((viewport.x) / viewport.baseTileSize );
        let y_min = Math.floor((viewport.y) / viewport.baseTileSize );
        let x_max = Math.ceil((viewport.x + viewport.w) / viewport.baseTileSize);
        let y_max = Math.ceil((viewport.y + viewport.h) / viewport.baseTileSize);
    
        if (x_min < 0) {
            x_min = 0;
            // this.viewport.x = 0;
            // x_max = this.viewport.w / viewport.tileSize;
        }
        if (y_min < 0) {
            y_min = 0;
            // this.viewport.y = 0;
            // y_max = this.viewport.w / viewport.tileSize;
        }
        if (x_max > this.mapSize) {
            x_max = this.mapSize;
        }
        if (y_max > this.mapSize) {
            y_max = this.mapSize;
        }
    
        // console.log(x_min, y_min, x_max, y_max,
        //     'vp:', viewport.x, viewport.y, viewport.w, viewport.h,
        //     'ts', viewport.tileSize);
    
        for (let i = x_min; i < x_max; i++) {
            for (let j = y_min; j < y_max; j++) {
                let tile_x = Math.floor(i * viewport.tileSize - viewport.x * viewport.tileSize / viewport.baseTileSize);
                let tile_y = Math.floor(j * viewport.tileSize - viewport.y * viewport.tileSize / viewport.baseTileSize);
    
                this._drawTile(this.map[j][i], tile_x, tile_y, viewport.tileSize);
            }
        }
    }
}