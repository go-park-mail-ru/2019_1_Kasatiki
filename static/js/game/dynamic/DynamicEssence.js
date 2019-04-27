import KeyboardControl from '../functions/KeyboardControl.js'
import MyMath from '../functions/myMath.js';

export default class DynamicEssence {
    
    constructor(
        xPos,
        yPos,
        xSize = 40,
        ySize = 40,
        URL = "/default_texture",
        velocity = 7,
    ) {
        this.keyHandler = new KeyboardControl();
        this.myMath = new MyMath();

        // Основные параметры
        this.hp = 100; // %
        this.hpCapacity = 100; // у.е
        this.velocity = velocity; // у.е
        this.vy = velocity;
        this.vx = velocity;

        // Координаты
        this.xPos = xPos;
        this.yPos = yPos;

        // Позиция прицела - только у плеера
        this.xAim = 0;
        this.yAim = 0;

        // Его размеры 
        this.xSize = xSize; // vh
        this.ySize = ySize; // vh

        this.teta = 0;

        // Тип оружия - только для usera 
        // this.melle = true;
        // this.gunId = 0; // 0 - knife

        // Бафы - только для usera 
        // this.bufs = {} // "key" : {}

        // Шмот
        // this.skinId = 0; // для игрока
        this.texture = URL; // URL 

        // this.immortal = false; // для рекламы
    }

    // Логика перемещения только для рекламы 
    logic() {
    }

    interact() {

    }

    // Геттеры

    get top() {
        return this.yPos;
    }

    get right() {
        return this.xPos + this.xSize;
    }

    get bottom() {
        return this.yPos + this.ySize;
    }

    get left() {
        return this.xPos;
    }

    get center() {
        return {x: this.xPos - this.xSize/2 ,
                y: this.yPos - this.ySize/2 };
    }

    // Сетеры

    set top(value) {
        this.yPos = value;
    }

    set right(value) {
        this.xPos = value - this.xSize;
    }

    set bottom(value) {
        this.yPos = value - this.ySize;
    }

    set left(value) {
        this.xPos = value;
    }
}
