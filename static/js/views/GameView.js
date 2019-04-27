import BaseView from './View.js';

import Game from '../game/game.js';

export default class GameView extends BaseView {
    constructor() {
        super(...arguments);
        this.Game = new Game(document.body, this.router);
        // this.initSpecialRoutes();
    }

    show() {
        this.Game.run();
    }

}