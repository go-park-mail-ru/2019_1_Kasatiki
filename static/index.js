const canvas = document.querySelector('.game');
const ctx = canvas.getContext('2d');

let mapChange = true;

let map = [];
let barriers = [];
let mapSize = 100;

let tileSize = 20;
let gameScreenSize = 10;

canvas.height = mapSize * tileSize;
canvas.width = mapSize * tileSize;


class Adv {
    constructor(
        x, y
    ) {
        this.x = x;
        this.y = y;

        this.v = 5;
    }

    draw() {
        ctx.beginPath();
        ctx.fillStyle = '#ffffff';
        ctx.fillRect(this.x, this.y, tileSize, tileSize);
        ctx.closePath();
    }
}

advs = [];
for (let i = 0; i < 10; i++) {
    advs[i] = new Adv(0, 0);
}

class Player {
    constructor(
        x, y
    ) {
        this.x = x;
        this.y = y;

        this.v = 5;
    }

    draw() {
        ctx.beginPath();
        ctx.fillStyle = '#45FF70';
        ctx.fillRect(this.x, this.y, tileSize, tileSize);
        ctx.fillRect(5, 5, tileSize, tileSize);
        ctx.closePath();
    }

    move(keyMap) {
        // if (keyMap.up) {
        //     this.y -= this.v;
        // }
        // if (keyMap.down) {
        //     this.y += this.v;
        // }
        // if (keyMap.left) {
        //     this.x -= this.v;
        // }
        // if (keyMap.right) {
        //     this.x += this.v;
        // }
    }
}

var keyMap = {
    up : false,
    left : false,
    down : false,
    right : false,
    zoom : false,
}


document.addEventListener('keydown', keyDown, false);
document.addEventListener('keyup', keyUp, false);

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

function checkMapSize() {
    if (keyMap.zoom) {
        if (gameScreenSize < 30) {
            gameScreenSize ++;
            tileSize--;
        }
    } else if (!keyMap.zoom) {
        if (gameScreenSize > 20) {
            gameScreenSize --;
            tileSize++;
        }
    }
}

let enemy = {
    x : 48,
    y : 48,
}

const player = new Player(6 * tileSize, 6 * tileSize);

let viewport = {
    screen : [],
    offset : [],
    startTile : [],
    endTile : [],
    update	: function(px, py) {
        this.offset[0] = Math.floor((this.screen[0]/2) - px);
        this.offset[1] = Math.floor((this.screen[1]/2) - py);

        var tile = [px, py]

        this.startTile[0] = tile[0] - 1 - Math.ceil((this.screen[0]/2) / tileSize);
        this.startTile[1] = tile[1] - 1 - Math.ceil((this.screen[1]/2) / tileSize);

        if(this.startTile[0] < 0) { this.startTile[0] = 0; }
        if(this.startTile[1] < 0) { this.startTile[1] = 0; }

        this.endTile[0] = tile[0] + 1 + Math.ceil((this.screen[0]/2) / tileSize);
        this.endTile[1] = tile[1] + 1 + Math.ceil((this.screen[1]/2) / tileSize);

        if(this.endTile[0] >= mapSize) { this.endTile[0] = mapSize-1; }
        if(this.endTile[1] >= mapSize) { this.endTile[1] = mapSize-1; }
    }
}

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

function drawTile(tile, x, y) {
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
    ctx.fillRect(y, x, tileSize, tileSize);
}

function drawBarier() {
    for (let i = 0; i < barriers.length; i++) {
        ctx.strokeStyle = '#000000';
        ctx.strokeRect(barriers[i].object.x, barriers[i].object.y, barriers[i].object.xsize, barriers[i].object.ysize);

        // console.log(barriers[i].object.x, barriers[i].object.y, barriers[i].object.xsize, barriers[i].object.ysize);
    }
}

function drawMap(x, y) {
    // console.log('x: ', x, 'y: ', y);
    // viewport.update(y, x);

    // console.log(viewport.startTile, viewport.endTile);

    // for(var i = viewport.startTile[1]; i <= viewport.endTile[1]; ++i)
    // {
    //     for(var j = viewport.startTile[0]; j <= viewport.endTile[0]; ++j)
    //     {
    //         drawTile(map[i][j], i * tileSize, j * tileSize);
    //     }
    // }

    for (i = 0; i < mapSize; i++) {
        for (j = 0; j < mapSize; j++) {
            drawTile(map[i][j], i * tileSize, j * tileSize);
        }
    }
}

let cord = {
    x : 0,
    y : 0,
}

const bounds = canvas.getBoundingClientRect();
document.body.onmousemove = function(evt) {
    cord.x = evt.clientX - bounds.left;
    cord.y = evt.clientY - bounds.top;
}

function renderEnemy() {
    ctx.beginPath();
    ctx.fillStyle = '#F52B00';
    ctx.fillRect(enemy.x, enemy.y, tileSize, tileSize);
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
        map = data["map"].field
        mapSize = data["map"].sizex
        tileSize = data["map"].tailsize

        barriers = data["barriers"]

        // console.log(barriers[0]);

        console.log(barriers[0].object.x, barriers[0].object.y, barriers[0].object);

        // canvas.height = mapSize * tileSize;
        // canvas.width = mapSize * tileSize;
        mapChange = false
        // console.log(data)
    }

    if (data["players"][0].id == 1) {
        player.x = data["players"][0].object.x
        player.y = data["players"][0].object.y

        enemy.x = data["players"][1].object.x
        enemy.y = data["players"][1].object.y
    } else {
        player.x = data["players"][1].object.x
        player.y = data["players"][1].object.y

        enemy.x = data["players"][0].object.x
        enemy.y = data["players"][0].object.y
    }
    
    for (let i = 0; i < advs.length; i++) {
        advs[i].x = data['advs'][i].object.x;
        advs[i].y = data['advs'][i].object.y;
    }


});

socket.addEventListener("error", (error) => {
    console.error(error)
});


function loop() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    // player.move(keyMap);
    // checkMapSize();
    let y = Math.floor(player.y/tileSize);
    let x = Math.floor(player.x/tileSize);

    drawMap( x, y );
    drawBarier();


    player.draw();
    renderEnemy();
    for (let i = 0; i < advs.length; i++) {
        advs[i].draw();
    }

    let json = JSON.stringify(keyMap);

    if (socketOpen) {
        socket.send(json);
    }
    requestAnimationFrame(loop)
}

function gameStart() {
    createMap();

    // viewport.screen = [gameScreenSize * tileSize, gameScreenSize * tileSize]

    requestAnimationFrame(loop);
}

gameStart();