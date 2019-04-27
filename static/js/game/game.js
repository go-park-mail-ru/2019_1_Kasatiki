import Screen from './functions/Screen.js'

import CollisionHandler from './functions/collisionHandler.js';
import Player from './dynamic/Player.js';
import Buff from './static/Buff.js';
import Barrier from './static/Barrier.js'
import Handler from './functions/Handler.js'
import Bullet from './dynamic/bullet.js'
import Adv from './dynamic/Adv.js'
import Shop from './static/Shop.js';

import buffConfigs from '../game/configs/buffConfigs.js';

export default class Game {
    constructor( 
        root = document.body,
        router
    ) {
        this.router = router;

        // Родительский узел DOM 
        this._root = root;
        this._root.innerHTML = '';

        // Игровой экран
        this._screen = new Screen(this.root);

        // Вспомогательные функции
        this.CollisionHandler = new CollisionHandler();
        this.handler = new Handler(this._screen._canvas);

        // Массив объектов
        this.eventsMap = {};

        this.objects = {
        };

        // Генерация карты
        this.borderW = 20;
        this.sectionsCount = 10;
        this.prm;

        this.advsPos = {
            0: [this.borderW + 10, Math.floor(this._screen.height/2), 25, 25],
            1 : [Math.floor(this._screen.width/2), this.borderW + 10, 25, 25],
            2 : [Math.floor(this._screen.width/2), this._screen.height - this.borderW - 35, 25, 25],
        }

        // Выстрелы
        this.lastFire = Date.now();

        // Логика волн
        this.waveTrigger = true;
        this.wavePause = false;
        this.waveCount = 0;
        this.pauseTimer = 0;
        this.totalAdvSpawn = 0;
        this.currentAdvCount = 0;

        // Инициализация объектов
        this.objects['players'] = [];
        this.objects['buffers'] = [];
        this.objects['bullets'] = [];
        this.objects['barriers'] = [];
        this.objects['advs'] = [];
        this.objects['shops'] = [];

        this._player = new Player(
            Math.floor(this._screen.width/2), Math.floor(this._screen.height/2),
            20, 20,
            "none",
            5
        );

        // this._buff = new Buff(
        //     100, 10,
        //     20, 20,
        //     "none"
        // )

        this.objects['players'].push(this._player)
        // this.objects['buffers'].push(this._buff);

        // Игровые параметры
        this.objects['players'].forEach(player =>{
            player.score = 0;
        }); 
        this.currentTime = 0;
    }

    // Спасвним соперников
    _spawnAdvs(count) {
        for (let i = 0; i < count; i++) {
            let vel = 0.5 + 2 * Math.random();
            let pos = Math.floor(3 * Math.random())
            let adv = new Adv(...(this.advsPos[pos]), 'none', vel);

            this.objects['advs'].push(adv);
        }
    }

    // Строим границы
    _createBoards() {
        let barrierTop = new Barrier(0, 0, this._screen.width, this.borderW);
        let barrierLeft = new Barrier(0, 0, this.borderW, this._screen.height);
        let barrierRight = new Barrier(this._screen.width - this.borderW, this.borderW, this._screen.width, this._screen.height);
        let barrierBottom = new Barrier(this.borderW, this._screen.height - this.borderW, this._screen.width - this.borderW, this._screen.height);
        this.objects['barriers'].push(barrierTop, barrierLeft, barrierRight, barrierBottom);
    }

    // Вычисляем параметры сетки
    _calculateParams() {
        // Ширина и высота каждой секции 
        let xStep = (this._screen.width - 2 * this.borderW) / this.sectionsCount; 
        let yStep = (this._screen.height - 2 * this.borderW) / this.sectionsCount;

        // Ширина и высота каждого блока секции
        let xBlockSize = xStep / 2;
        let yBlockSize = yStep / 2;

        // Количество блоков 
        let blocksCount = this.sectionsCount * 2;

        const mapsParams = {
            'xStep' : xStep,
            'yStep' : yStep,
            'xBlockSize' : xBlockSize,
            'yBlockSize' : yBlockSize,

            'blocksCount' : blocksCount,
        }

        return mapsParams;
    }

    // Строим границы по сетке
    _spawnBarriers() {
        const that = this;

        for (let i = 1; i < this.sectionsCount - 1; i++) {

            for (let j = 1; j < this.sectionsCount - 1; j++) {

                if (!(i == Math.floor(this.sectionsCount/2) && j == Math.floor(this.sectionsCount/2))) {

                    let xSection = this.borderW + j * that.prm['xStep'];
                    let ySection = this.borderW + i * that.prm['yStep'];        

                    let barrierCout = Math.floor(Math.random() * 2);
                    let idxs = [];
                    

                    for (let k = 0; k < barrierCout; k++) {
                        let idx;
                        let check = false;
                        idx = Math.floor(Math.random() * 4);
                        if (idxs.length != 0) {
                            while (!check) {
                                idx = Math.floor(Math.random() * 4);
                                console.log('idx',idx);
                                for (let a = 0; a < idxs.length; a++) {
                                    if (idxs[a] == idx) {
                                        check = false;
                                        break;
                                    } 
                                    check = true;
                                }
                            }
                        }
                        idxs.push(idx);
                    }

                    for (let p = 0; p < idxs.length; p++) {
                        let barrier;
                        switch (idxs[p]) {
                            case 0: {
                                barrier = new Barrier(xSection, ySection, that.prm['xBlockSize'], that.prm['yBlockSize']);
                                that.objects['barriers'].push(barrier);
                                console.log(barrier.left);
                                break;
                            }
                            case 1: {
                                barrier = new Barrier(xSection + that.prm['xBlockSize'], ySection, that.prm['xBlockSize'], that.prm['yBlockSize']);
                                that.objects['barriers'].push(barrier);
                                console.log(barrier.left);
                                break;
                            }
                            case 2: {
                                barrier = new Barrier(xSection, ySection + that.prm['yBlockSize'], that.prm['xBlockSize'], that.prm['yBlockSize']);
                                that.objects['barriers'].push(barrier);
                                break;
                            }
                            case 3: {
                                barrier = new Barrier(xSection + that.prm['xBlockSize'], ySection + that.prm['yBlockSize'], that.prm['xBlockSize'], that.prm['yBlockSize']);
                                that.objects['barriers'].push(barrier);
                                break;
                            }

                        }
                    }
            
                }
            }

        }
    }

    // Генерируем карту
    _generateMap() {
        this._createBoards();
        this.prm = this._calculateParams();
        this._spawnBarriers();
    }

    _spawnBuffs(count) {
        for (let i = 0; i < count; i++) {
            let idx = Math.floor(Math.random() * 1.999);
            let buff = new Buff(
                Math.floor(Math.random() * this._screen.height),
                Math.floor(Math.random() * this._screen.width),
                20, 20,
                buffConfigs[idx],
            );
            console.log(buff.cfg);
            this.objects['buffers'].push(buff);
        }
    }

    // Вспомогательная сетка
    // _drawGrid() {
    //     const that = this;

    //     for (let i = 0; i < this.sectionsCount*4; i++) {

    //         for (let j = 0; j < this.sectionsCount*4; j++) { 
    //             that._screen.ctx.strokeStyle = 'red'; 
    //             that._screen.ctx.moveTo(this.borderW + j*this.prm['xBlockSize'], this.borderW + i*this.prm['yBlockSize']);
    //             that._screen.ctx.lineTo(this.borderW + j*this.prm['xBlockSize'], this._screen.height - this.borderW);
    //             that._screen.ctx.stroke();
    //         }
    //         that._screen.ctx.moveTo(this.borderW,this.borderW + i*this.prm['yBlockSize']);
    //         that._screen.ctx.lineTo(this._screen.width - this.borderW, this.borderW + i*this.prm['yBlockSize']);
    //         that._screen.ctx.stroke();
    //     }
    // }

    isEmpty(obj) {
        for (var key in obj) {
          return false;
        }
        return true;
    }

    frame() {
        this.eventsMap = this.handler.sendEventMap();

        if (this.waveTrigger) {
            this.totalAdvSpawn += 5;
            this.currentAdvCount = this.totalAdvSpawn;
            this.waveCount++;
            // console.log(this.totalAdvSpawn)
            this._spawnAdvs(this.totalAdvSpawn);
            // this._spawnBuffs(5);
            // console.log(this.objects['buffers']);
        }

        this.waveTrigger = false;

        // Стрельба
        if (this.eventsMap['mouseClick'] && !this.objects['players'][0].inShop) {
            if (Date.now() - this.lastFire > this.objects['players'][0].weapon.fireRate) {
                if (this.objects['players'][0].weapon.name == 'Shotgun') {
                    for (let i = 0; i < 10; i++) {
                        let bullet = new Bullet(
                            this.objects['players'][0].centerX,
                            this.objects['players'][0].centerY,
                            this.objects['players'][0].weapon.bulletSize,
                            this.objects['players'][0].weapon.bulletSize,
                            this.objects['players'][0].weapon.bulletColor,
                            this.objects['players'][0].weapon.velocity,
                            this.objects['players'][0].weapon.damage,
                            this.eventsMap['mouseX'] + 20 * Math.random()*Math.pow(-1, i), this.eventsMap['mouseY'] + 20 * Math.random()*Math.pow(-1, i)
                        );
                        this.objects['bullets'].push(bullet);
                    }
                } else {
                    let bullet = new Bullet(
                        this.objects['players'][0].centerX - this.objects['players'][0].weapon.bulletSize / 2,
                        this.objects['players'][0].centerY - this.objects['players'][0].weapon.bulletSize / 2,
                        this.objects['players'][0].weapon.bulletSize,
                        this.objects['players'][0].weapon.bulletSize,
                        this.objects['players'][0].weapon.bulletColor,
                        this.objects['players'][0].weapon.velocity,
                        this.objects['players'][0].weapon.damage,
                        this.eventsMap['mouseX'], this.eventsMap['mouseY']
                    );
                    this.objects['bullets'].push(bullet);
                }
                this.lastFire = Date.now();
            }
        }



        // обработка логики объектов
        this.objects['players'][0].logic(this.eventsMap, this.width, this.height);
        this.objects['advs'].forEach(adv => {
            adv.logic(this.objects['players'][0].xPos, this.objects['players'][0].yPos);
        });
        // движение пуль
        this.objects['bullets'].forEach(element => {
            element.go();
        });

        this._screen.render(this.objects);

        this.CollisionHandler.handleCollisions(this.objects, this.scoreObj);

        this.currentTime++;

        if (this.currentTime >= 60) {
            this.currentTime = 0;
            this.objects['players'].forEach(player =>{
                player.score++;
            }); 
            this.objects['players'][0].hp -= 0.5;
        }

        if (this.objects['advs'].length == 0 && !this.wavePause) {
            this.objects['players'][0].hp = 100;
            this.currentTime = 0;
            this.pauseTimer = 25 * 60;
            this.wavePause = true;
            let shop = new Shop(this._screen.width - 120 - this.borderW, this._screen.height/2, 100, 100);
            this.objects['shops'].push(shop);
        }

        if (this.wavePause) {
            this.currentTime = 0;
            this.pauseTimer -= 1;
            this._screen.showPauseTime(Math.floor(this.pauseTimer/60));
            if (this.pauseTimer == 0) {
                this.objects['shops'].splice(0,1);
                this.wavePause = false;
                this.waveTrigger = true;
            }
        }

        this._screen.showWaveNumber(this.waveCount); 
        this._screen.showInfo(this.objects['players'][0].score, this.objects['players'][0].hp);

        this._checkDeath();

        // console.log(this.objects['players'][0].xPos, this.objects['players'][0].yPos);

        // console.log(this.objects['players'][0].left, this.objects['players'][0].top);
        // let a = this.objects['barriers'][5].left; let b = this.objects['barriers'][5].top;
        // console.log(a, b);


        requestAnimationFrame(time => this.frame());
    }

    _checkDeath() {
        if (this.objects.players[0].hp <= 0) {
            this.defeat();
        }
    }

    // Поражение
    defeat() {
        this.router.go('/');
        throw new Error('Ok');
    }

    // Победа 
    victory() {
        this.router.go('/win');
    }

    // Инит метод : цикл -> отрисовка 
    run() {
        this.canvas = this._screen.createCanvas();
        this._generateMap();
        requestAnimationFrame(time => this.frame());
    }


}  