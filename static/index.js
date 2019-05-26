const canvas = document.querySelector('.game');
const ctx = canvas.getContext('2d');

console.log(canvas);

let mapChange = true;

let map = [];
let barriers = [];
let bullets = [];
let mapSize = 100;

let bounds = canvas.getBoundingClientRect();

let tileSize = 50;
let gameScreenW = 16;
let gameScreenH = 16;

// Время последнего выстрела
let lastFire = Date.now();

canvas.height = tileSize * gameScreenH;
canvas.width = tileSize * gameScreenW;

let keyMap = {
    up : false,
    left : false,
    down : false,
    right : false,

    shot: false,
    angle: 0,
}

class Player {
    constructor(
        x, y
    ) {
        this.x = x;
        this.y = y;

        this.id = 1;

        this.v = 5;

        this.c = 0.2;
    }

    draw() {
        ctx.beginPath();
        ctx.fillStyle = '#45FF70';
        // if (this.x <= canvas.width / 2 && this.y <= canvas.height / 2) {
        //     ctx.fillRect(this.x, this.y, tileSize, tileSize);
        // } else if (this.x > canvas.width / 2 && this.y > canvas.height / 2) {
        //     ctx.fillRect(canvas.width / 2, canvas.height / 2,  tileSize, tileSize);
        // } else if (this.x <= canvas.width / 2 && this.y > canvas.height / 2) {
        //     ctx.fillRect(this.x, canvas.height / 2,  tileSize, tileSize);
        // } else if (this.x > canvas.width / 2 && this.y <= canvas.height / 2) {
        //     ctx.fillRect(canvas.width / 2, this.y,  tileSize, tileSize);
        // }
        // ctx.fillRect(canvas.width / 2, canvas.height / 2,  tileSize, tileSize);
        ctx.fillRect(Math.round(this.x - viewport.x), Math.round(this.y - viewport.y), tileSize, tileSize);
        ctx.closePath();
    }

    move(keyMap) {
        if (keyMap.up) {
            this.y -= this.v;
        }
        if (keyMap.down) {
            this.y += this.v;
        }
        if (keyMap.left) {
            this.x -= this.v;
        }
        if (keyMap.right) {
            this.x += this.v;
        }
    }
}

class Viewport {
    constructor(x, y, w, h) {
        this.x = x;
        this.y = y;
        this.w = w;
        this.h = h;
    }

    update(x ,y) {

        if (zoom) {
            if (this.w < gameScreenH * 2 * 50) {
                tileSize -= 0.5;
                this.w = this.w + gameScreenW;
                this.h = this.h + gameScreenH;
            }
        } else if (!zoom) {
            if (this.w > gameScreenH * 50) {
                tileSize += 0.5;
                this.w = this.w - gameScreenW;
                this.h = this.h - gameScreenH;
            }
        }

        // this.x += (x - this.x - this.w * 0.5) * 0.05;
        // this.y += (y - this.y - this.h * 0.5) * 0.05;

        this.x = x - this.w / 2;
        this.y = y - this.h / 2;

        // console.log('port', this.x, this.y, this.w, this.h);
    }
}

const viewport = new Viewport(0, 0, gameScreenW * tileSize, gameScreenH * tileSize);

let zoom = false;

document.addEventListener('keydown', keyDown, false);
document.addEventListener('keyup', keyUp, false);
document.addEventListener('mouseup', mouseUp, false);
document.addEventListener('mousedown', mouseDown, false);

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
        zoom = true;
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
        zoom = false;
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

let enemy = {
    x : 48,
    y : 48,
}

const player = new Player(mapSize * tileSize / 2, mapSize * tileSize / 2);

function createMap() {
    for (let i = 0; i < mapSize; i++) {
        for (let j = 0 ; j < mapSize; j++) {
            if (i == 0 || i == mapSize-1 || j == 0 || j == mapSize -1) {
                map[mapSize*i + j] = 1;
            } else {
                map[mapSize*i + j] = Math.floor( Math.random() * 1.999999 );
            }
        }
    }
}

function drawTile(tile, x, y, i, j) {
    let color;
    switch (tile) {
        case 0:
            color = '#979797';
            break;
        case 1:
            color = '#5B5B5B';
            break;
    }

    ctx.fillStyle = color;
    ctx.fillRect(x, y, tileSize, tileSize);
    // ctx.fillStyle = '#000000'
    // ctx.font = tileSize / 4 + 'px serif';
    // let text = i + ', ' + j;
    // ctx.textAlign = 'center';
    // ctx.fillText(text, y * tileSize, x * tileSize);
}

function drawBarriers() {
    for (let i = 0; i < barriers.length; i++) {
        ctx.strokeStyle = '#000000';

        // if (barriers[i].object.x > viewport.x && barriers[i].object.y > viewport.y && barriers[i].object.x < (viewport.x + viewport.w) && barriers[i].object.y < (viewport.y + viewport.h)) {
        ctx.strokeRect(barriers[i].object.x - viewport.x, barriers[i].object.y - viewport.y, barriers[i].object.xsize, barriers[i].object.ysize);
        // }

        // console.log(barriers[i].object.x, barriers[i].object.y, barriers[i].object.xsize, barriers[i].object.ysize);
    }
}

function drawBullets() {
    if ( bullets.length != 0) {
        for (let i = 0; i < bullets.length; i++) {
            ctx.strokeStyle = '#000000';
            ctx.strokeRect(bullets[i].object.x - viewport.x, bullets[i].object.y - viewport.y, 5, 5);
        }
    }
}

function drawMap(x, y) {
    var x_min = Math.floor((viewport.x) / tileSize);
    var y_min = Math.floor((viewport.y) / tileSize);
    var x_max = Math.ceil((viewport.x + viewport.w) / tileSize);
    var y_max = Math.ceil((viewport.y + viewport.h) / tileSize);

    if (x_min < 0) {
        x_min = 0;
        // viewport.x = 0;
        // x_max = viewport.w / tileSize;
    }
    if (y_min < 0) {
        y_min = 0;
        // viewport.y = 0;
        // y_max = viewport.w / tileSize;
    }
    if (x_max > mapSize) {
        x_max = mapSize;
    }
    if (y_max > mapSize) {
        y_max = mapSize;
    }

    // console.log(x_min, y_min, x_max, y_max,
    //     'vp:', viewport.x, viewport.y, viewport.w, viewport.h,
    //     'ts', tileSize);

    for (let i = x_min; i < x_max; i++) {
        for (let j = y_min; j < y_max; j++) {
            let tile_x = Math.floor(i * tileSize - viewport.x);
            let tile_y = Math.floor(j * tileSize - viewport.y);

            drawTile(map[j][i], tile_x, tile_y, i ,j);
        }
    }
}

let cord = {
    x : 0,
    y : 0,
}

// document.body.onmousemove = function(evt) {
//     cord.x = evt.clientX - bounds.left;
//     cord.y = evt.clientY - bounds.top;
// }

function renderEnemy() {
    ctx.beginPath();
    ctx.fillStyle = '#F52B00';
    ctx.fillRect(Math.round(enemy.x - viewport.x), Math.round(enemy.y - viewport.y), tileSize, tileSize);
    ctx.closePath();
}

var socket = new WebSocket("ws://"+location.host+"/game/start");
var socketOpen = false;

let d = new Date();
d.setDate(d.getDate()+1);
document.cookie =
    "session_id="+Math.round(Math.random()*2**32).toString()+"; "+
    "path=/; "+
    "expires="+d.toUTCString()+";";

socket.addEventListener("open", (event) => {
    socketOpen = true;
});

socket.addEventListener("close", () => {
});

socket.addEventListener("message", (event) => {
    let data = JSON.parse(event.data)
    if (mapChange) {
        map = data.map.field;
        // mapSize = data.map.sizex;
        // tileSize = data.map.tailsize;

        barriers = data.barriers;

        player.id = data.id;

        mapChange = false;
        // console.log(data);

    }


    // if (data["players"][0].id == 1) {
    //     player.x = data["players"][0].object.x;
    //     player.y = data["players"][0].object.y;

    //     // enemy.x = data["players"][1].object.x
    //     // enemy.y = data["players"][1].object.y
    // } else {
    //     player.x = data["players"][1].object.x;
    //     player.y = data["players"][1].object.y;
    // }


    // console.log(data["bullets"], bullets)

    if (data["bullets"] != null) {
        bullets = data["bullets"]
    }

    if (data["players"] != null) {
        if (data["players"][0].id == player.id) {
            player.x += (data["players"][0].object.x - player.x) * player.c;
            player.y += (data["players"][0].object.y - player.y) * player.c;
    
            enemy.x += (data["players"][1].object.x - enemy.x) * player.c;
            enemy.y += (data["players"][1].object.y - enemy.y) * player.c;
            // console.log("player: ", data["players"][0].object.x, data["players"][0].object.y);
            // console.log("enemy: ", data["players"][1].object.x, data["players"][1].object.y);
            // }ss
        } else {
            player.x += (data["players"][1].object.x - player.x) * player.c;
            player.y += (data["players"][1].object.y - player.y) * player.c;
    
            enemy.x += (data["players"][0].object.x - enemy.x) * player.c;
            enemy.y += (data["players"][0].object.y - enemy.y) * player.c;
    
            // console.log("player: ", data["players"][1].object.x, data["players"][1].object.y);
            // console.log("enemy: ", data["players"][0].object.x, data["players"][0].object.y);
        }
    }



});

socket.addEventListener("error", (error) => {
    console.error(error)
});


function loop() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    drawMap(player.x, player.y);
    viewport.update(player.x, player.y);

    player.draw();
    renderEnemy();
    // drawBarriers();
    drawBullets();

    let json = JSON.stringify(keyMap);

    console.log(keyMap.shot);

    if (socketOpen) {
        socket.send(json);
    }
    requestAnimationFrame(loop)
}

function gameStart() {
    createMap();
    requestAnimationFrame(loop);
}

gameStart();