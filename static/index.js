import Viewport from './viewport.js';
import startListen from './eventListeners.js';
import Socket from './ws.js';
import Config from './config.js';
import Objects from './objects.js';

// Кнопки управления 
let keyMap = {
    up : false,
    left : false,
    down : false,
    right : false,

    shot: false,
    angle: 0,

    zoom: false,
}

const cfg      = new Config();
const viewport = new Viewport();
const ws       = new Socket(location.host);

let objs       = new Objects(cfg, viewport);

function loop() {
    objs.drawObjs(viewport);    
    viewport.update(objs.player.x, objs.player.y, keyMap.zoom, {val : objs.map.tileSize});

    objs.drawPlayers(viewport);

    let json = JSON.stringify(keyMap);

    if (ws.socketOpen) {
        ws.socket.send(json);
    }
    requestAnimationFrame(loop)
}

function gameStart() {
    objs.map._createMap();
    ws.startServe(objs);
    startListen(keyMap, objs.map.canvas);
    objs.map.setCanvasWH(viewport.gameScreenH, viewport.gameScreenW);
    requestAnimationFrame(loop);
}

gameStart();