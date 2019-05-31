import Map from './map.js';
import Player from './player.js';

export default class Objects {
    constructor(
        cfg,
        viewport
    ) {
        this.objs = {
            'bullets' : [],
            'advs'    : [],
            'buffs'   : [],
            'enemy'   : [],
        };
        
        this.cfg = cfg;

        this.map = new Map();
        this.map.setCanvasWH(viewport.gameScreenW * this.map.tileSize, viewport.gameScreenH * this.map.tileSize);

        this.player = new Player(this.map.mapSize * this.map.tileSize / 2, this.map.mapSize * this.map.tileSize / 2, this.map.ctx);
        this.enemy = {
            x : 0,
            y : 0,

            draw : (viewport) => {
                this.map.ctx.fillStyle = '#F52B00';
                this.map.ctx.fillRect(Math.round(this.x - viewport.x), Math.round(this.y - viewport.y), this.map.tileSize, this.map.tileSize.val);
            }
        }
    }

    drawObjs(viewport) {
        // Стираем все что есть на канвасе
        this.map.ctx.clearRect(0, 0, this.map.canvas.width, this.map.canvas.height);

        // Отрисовываем карту
        this.map.drawMap(viewport, {val : this.map.tileSize}, {val : this.map.mapSize});

        // this.objs.forEach( obj => {
        //     obj.forEach( ess => {
        //         this.map.ctx.fillStyle = cfg.color(ess[obj]);
        //         this.map.ctx.fillRect(ess.object.x - this,map.viewport.x, ess.object.y - this,map.viewport.x, this.map.tileSize, this.map.tileSize);
        //     })
        // })
    }

    drawPlayers(viewport) {
        this.player.draw(viewport, {val : this.map.tileSize});
        this.enemy.draw(viewport);
    }
}