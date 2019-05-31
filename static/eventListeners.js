export default function startListen(keyMap, canvas) {
    document.addEventListener('keydown', keyDown, false);
    document.addEventListener('keyup', keyUp, false);
    document.addEventListener('mouseup', mouseUp, false);
    document.addEventListener('mousedown', mouseDown, false);

    // Позиция курсора относительно вьюпорта (ширины/высоты канваса)
    let bounds = canvas.getBoundingClientRect();

    // Время последнего выстрела
    let lastFire = Date.now();

    function mouseDown(evt) {
        if (Date.now() - lastFire > 100) {
            keyMap.shot = true;
            let x = evt.clientX - bounds.left - 400;
            let y = evt.clientY - bounds.top - 400;
        
            keyMap.angle = Math.atan2(y, x);

            lastFire = Date.now();
        } else {
            keyMap.shot = false;
        }
    }

    function mouseUp() {
        keyMap.shot = false;
    }

    function keyDown(e) {
        if (e.keyCode === 32) {
            keyMap.zoom = true;
        }

        if (e.keyCode === 65) {
            keyMap.left = true;
        }
        if (e.keyCode === 68) {
            keyMap.right = true;
        }
        if (e.keyCode === 83) {
            keyMap.down = true;
        }
        if (e.keyCode === 87) {
            keyMap.up = true;
        }
    }

    function keyUp(e) {
        if (e.keyCode === 32) {
            keyMap.zoom = false;
        }

        if (e.keyCode === 65) {
            keyMap.left = false;
        }
        if (e.keyCode === 68) {
            keyMap.right = false;
        }
        if (e.keyCode === 83) {
            keyMap.down = false;
        }
        if (e.keyCode === 87) {
            keyMap.up = false;
        }
    }
};