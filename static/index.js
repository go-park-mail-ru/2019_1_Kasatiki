const canvas = document.querySelector('.game');
const ctx = canvas.getContext('2d');

let mapChange = true;

let map = [];
let mapSize = 15;

let tileSize = 15;
let gameScreenSize = 25;

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
        ctx.fillRect(this.x, this.y, tileSize * 2, tileSize * 2);
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
    update : function(px, py) {
        this.offset[0] = Math.floor((this.screen[0]/2) - px * tileSize);
        this.offset[1] = Math.floor((this.screen[1]/2) - py * tileSize);

        let tile = [ px, py ];


        this.startTile[0] = tile[0] - 1 - Math.ceil((this.screen[0]/2) / tileSize);
        this.startTile[1] = tile[1] - 1 - Math.ceil((this.screen[1]/2) / tileSize);

        if(this.startTile[0] < 0) { this.startTile[0] = 0; }
        if(this.startTile[1] < 0) { this.startTile[1] = 0; }

        this.endTile[0] = tile[0] + 1 + Math.ceil((this.screen[0]/2) / tileSize);
        this.endTile[1] = tile[1] + 1 + Math.ceil((this.screen[1]/2) / tileSize);

        if(this.endTile[0] >= mapSize) { this.endTile[0] = mapSize; }
        if(this.endTile[1] >= mapSize) { this.endTile[1] = mapSize; }
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
    ctx.fillRect(x, y, tileSize, tileSize);
}

function drawMap(x, y) {
    // let iStart;
    // let iEnd;
    // let jStart;
    // let jEnd;

    // if (x - Math.floor(gameScreenSize) / 2 < 0) {
    //     iStart = 0;
    //     iEnd = gameScreenSize;
    // } else if (x + Math.ceil(gameScreenSize) / 2 > mapSize) {
    //     iStart = mapSize - gameScreenSize;
    //     iEnd = mapSize;
    // } else if (0 < x < mapSize) {
    //     iStart = x - Math.floor(gameScreenSize) / 2;
    //     iEnd = x + Math.ceil(gameScreenSize) / 2;
    // }

    // if (y - Math.floor(gameScreenSize) / 2 < 0) {
    //     jStart = 0;
    //     jEnd = gameScreenSize;
    // } else if (y + Math.ceil(gameScreenSize) / 2> mapSize) {
    //     jStart = mapSize - gameScreenSize;
    //     jEnd = mapSize;
    // } else if (0 < y < mapSize) {
    //     jStart = y - Math.floor(gameScreenSize) / 2;
    //     jEnd = y + Math.ceil(gameScreenSize) / 2;
    // }

    // let a = 0;
    // let b = 0;
    // for (i = iStart; i < iEnd; i++) {

    //     for (j = jStart; j < jEnd; j++) {
    //         drawTile(map[mapSize * i + j], a * tileSize, b * tileSize);
    //         b++;
    //     }
    //     b = 0;
    //     a++;
    // }

    for (i = 0; i < mapSize; i++) {
        for (j = 0; j < mapSize; j++) {
            drawTile(map[mapSize * i + j], i * tileSize, j * tileSize);
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

        canvas.height = mapSize * tileSize;
        canvas.width = mapSize * tileSize;
        mapChange = false
    }

    for (let i = 0; i < advs.length; i++) {
        advs[i].x = data['advs'][i].object.x;
        advs[i].y = data['advs'][i].object.y;
    }

    player.x = data["players"][0].object.x
    player.y = data["players"][0].object.y

    enemy.x = data["players"][1].object.x
    enemy.y = data["players"][1].object.y

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
    requestAnimationFrame(loop);
}

gameStart();