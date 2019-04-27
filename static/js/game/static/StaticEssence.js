export default class StaticEssence {
    
    constructor(
        xPos,
        yPos,

        xSize = 50,
        ySize = 50,
        URL = "/default_texture",
    ) {
        // Основные параметры
        this.hpCapacity = 100; // у.е
        // this.velocity = velocity; // у.е

        // Координаты
        this.xPos = xPos;
        this.yPos = yPos;

        // Его размеры 
        this.xSize = xSize; // vh
        this.ySize = ySize; // vh


        this.texture = URL; // URL 

        this.immortal = true; // для рекламы
    }

    _render() {
        
    }

    interact() {

    }

}