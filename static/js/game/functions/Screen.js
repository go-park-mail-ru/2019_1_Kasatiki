import Escape from '../static/escape.js'

export default class Screen {
    constructor(
        root = document.body
    ) {
        this._root = root;

        // Параметры canvas
        this._canvas = document.createElement('canvas');
        // this._canvas.style.zIndex = '0';
        this.ctx;
        this._canvas.className = "gameScreen";

        // Размеры карты (видимая область)
        this.width = window.innerWidth;
        this.height = window.innerHeight;

        // Параметры отображения текста
        this.fontCfg = '25px Arial';
        this.textPosY = 45;
    }

    set canvas(ctx) {
        this.ctx = ctx;
    } 

    _createMap() {

    }

    createCanvas() {
        this._root.innerHTML = '';
        this._canvas.width = this.width;
        this._canvas.height = this.height;
        this._root.appendChild(this._canvas);
        this.ctx = this._canvas.getContext('2d');
        return this._canvas;
    }

    render(objects = []) {
        this.ctx.clearRect(0, 0, this.width, this.height);

        this.width = window.innerWidth;
        this.height = window.innerHeight;
        
        const that = this;
        objects['players'].forEach(obj => {
            obj.render(that.ctx);
            // that._renderEssence('players', obj);
        });
        objects['bullets'].forEach(obj => {
            obj.render(that.ctx);
        });
        objects['barriers'].forEach(obj => {
            obj.render(that.ctx);
        });
        objects['advs'].forEach(obj => {
            obj.render(that.ctx);
        });
        objects['shops'].forEach(obj => {
            obj.render(that.ctx);
        });
        objects['buffers'].forEach(obj => {
            obj.render(that.ctx);
        });
    }

    showInfo(score, health) {
        this.ctx.fillStyle = "#000";
        this.ctx.font = this.fontCfg;
        this.ctx.fillText('score: ' + score,this.width/2 - 250, this.textPosY);
        this.ctx.fillStyle = "red";
        this.ctx.font = this.fontCfg;
        this.ctx.fillText('hp: ' + health,this.width/2 , this.textPosY);
    }

    showPauseTime(time) {
        this.ctx.fillStyle = "#000";
        this.ctx.font = this.fontCfg;
        this.ctx.fillText('pause: ' + time,this.width - 300, this.textPosY);
    }

    showWaveNumber(number) {
        this.ctx.fillStyle = "#000";
        this.ctx.font = this.fontCfg;
        this.ctx.fillText('Wave: ' + number,100, this.textPosY);
    }
}