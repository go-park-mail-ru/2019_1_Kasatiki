import BaseView from './View.js';

const { NetworkHandler } = window;

/**
 * Класс с отрисовкой формы логина.
 */
export default class LogoutView extends BaseView {
    constructor() {
        super(...arguments);
    }

    show() {
        let that = this;
        NetworkHandler.doGet({
            // eslint-disable-next-line no-unused-vars
            callback() {
                that.router.go('/');
            },
            path : '/api/logout',
        });
    }
}