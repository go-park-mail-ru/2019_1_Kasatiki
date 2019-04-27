export default class BaseView {
    /**
     * Конструктор
     * @param {HTMLelement} root - экран, где будет весь экш.
     * @param {Object} router - ссылка на инстанс роутера.
     */
    constructor(
        root = document.body,
        router = Object,
    ) {
        this.root = root;
        this.router = router;
        /**
         * Мапа специальных роутеров: Map {'string': 'function'}
         * Используется, например, в LoginView для маршрута /authorizeuser,
         * При нашатии на <a href="/autho...></a> роутер не переходит по ссылке,
         * а вызывает метод из Map {'string': 'function'}.
         */
        this.specialRoutes = {};
    }

    /**
     * метод, который будет печатать что-нибудь экран.
     */
    show() {
        this.root.innerHTML = '';
    }
}