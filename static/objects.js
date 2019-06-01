import Map from './map.js';
import Player from './player.js';

export default class Objects {
    constructor(
        cfg,
        vp
    ) {
        this.objs = {
            'bullets' : [],
            'advs'    : [],
            'buffs'   : [],
            'enemy'   : [],
        };
        
        this.cfg = cfg;

        this.map = new Map();
        this.map.setCanvasWH(vp.gameScreenW * vp.tileSize, vp.gameScreenH * vp.tileSize);

        this.player = new Player(this.map.mapSize * vp.tileSize / 2, this.map.mapSize * vp.tileSize / 2, this.map.ctx);
        this.enemy = {
            x : 0,
            y : 0,

            draw : (vp) => {
                this.map.ctx.fillStyle = '#F52B00';
                this.map.ctx.fillRect(Math.round(this.x - vp.x), Math.round(this.y - vp.y), vp.tileSize, vp.tileSize);
            }
        }
    }

    drawObjs(viewport) {
        // Стираем все что есть на канвасе
        this.map.ctx.clearRect(0, 0, this.map.canvas.width, this.map.canvas.height);

        // Отрисовываем карту
        this.map.drawMap(viewport);

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