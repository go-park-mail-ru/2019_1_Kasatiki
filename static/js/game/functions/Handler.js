export default class Handler {

    constructor(
        canvas
    ) {
        this.canvas = canvas;
        this.eventsMap = {
            // Нажатие клавиш
            'up': false,
            'left': false,
            'down': false,
            'right': false,
            'interact' : false,

            // Нажатие мышки
            'mouseClick' : false,

            // Координаты мышки
            'mouseX' : 0,
            'mouseY' : 0
        }

        this.bounds = this.canvas.getBoundingClientRect();
        this._listenEvents();
    }

    _listenEvents() {
        const that = this;

        document.addEventListener('keydown', keyDownHandler, false);
        document.addEventListener('keyup', keyUpHandler, false);
        document.addEventListener('mousedown', mouseClickDown, false)
        document.addEventListener('mouseup', mouseClickUp, false)
        document.addEventListener('mousedown', keyDownHandler, false);

        document.body.onmousemove = function(evt) {
            that.eventsMap.mouseX = evt.pageX - that.bounds.left;
            that.eventsMap.mouseY = evt.pageY - that.bounds.top;
        }

        function mouseClickDown() {
            that.eventsMap.mouseClick = true;
        } 

        function mouseClickUp() {
            that.eventsMap.mouseClick = false;
        }


        function keyDownHandler(e) {
            if(e.code == "KeyD" || e.code == "ArrowRight") {
                that.eventsMap['right'] = true;
            }
            else if(e.code == "KeyA" || e.code == "ArrowLeft") {
                that.eventsMap['left'] = true;
            }
            else if(e.code == "KeyW" || e.code == "ArrowUp") {
                that.eventsMap['up'] = true;
            }
            else if(e.code == "KeyS" || e.code == "ArrowDown") {
                that.eventsMap['down'] = true;
            }
            
            if(e.code == 'KeyE') {
                if (that.eventsMap['interact']) {
                    that.eventsMap['interact'] = false ;
                } else {   
                    that.eventsMap['interact'] = true;
                }
            }
        }
        
        function keyUpHandler(e) {
            if(e.code == "KeyD" || e.code == "ArrowRight") {
                that.eventsMap['right'] = false;
            }
            else if(e.code == "KeyA" || e.code == "ArrowLeft") {
                that.eventsMap['left'] = false;
            }
            else if(e.code == "KeyW" || e.code == "ArrowUp") {
                that.eventsMap['up'] = false;
            }
            else if(e.code == "KeyS" || e.code == "ArrowDown") {
                that.eventsMap['down'] = false;
            } 
        }

    }

    addEventListener(eventType, callback) {
        document.addEventListener(eventType, callback, false);
    }


    addObject(name, obj) {
        this.objects[name]  = obj;
    }

    sendEventMap() {
        return this.eventsMap;
    }

}

