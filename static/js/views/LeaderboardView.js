import BaseView from './View.js';

import LeaderboardComponent from '../components/LeaderboardComponent/LeaderboardComponent.js';

const { NetworkHandler } = window;

/**
 * Класс с отрисовкой формы логина.
 */
export default class LeaderboardView extends BaseView {
    constructor() {
        super(...arguments);
        this.LeaderboardComponent = new LeaderboardComponent();
        this.initSpecialRoutes();
    }

    show() {
        const that = this;
        NetworkHandler.doGet({
			callback(data) {
				that.LeaderboardComponent._usersArr = data;
				that.LeaderboardComponent.render();
			},
			path: '/api/leaderboard?offset=1',
		});
    }

    initSpecialRoutes() {

    }
}