import BaseView from './View.js';

import CountdownComponent from '../components/CountdownComponent/CountdownComponent.js';
/**
 * Класс с отрисовкой формы логина.
 */
export default class CountdownView extends BaseView {
    constructor() {
        super(...arguments);
        this.CountdownComponent = new CountdownComponent();
    }

    show() {
        this.root.innerHTML = this.CountdownComponent.render();
        this.CountdownComponent.canvasFullScreen();
        this.run();
    }

    /**
     * Рисует циферку в определенное по таймингу время.
     * @param {HTMLelement} canvas 
     * @param {number} deltaTime - Время с начала анимиции отсчета в мс.
     */
    draw(canvas, deltaTime) {
        let canvasContext = canvas.getContext('2d');
        canvasContext.clearRect(0, 0, canvas.width, canvas.height);

        canvasContext.font = '48px Arial, Helvetica, sans-serif';
        canvasContext.fillStyle = 'red';

        let digit = 3 - parseInt(deltaTime/1000);
        let x = canvas.width / 2;
        let y = canvas.height / 2;
        // Пол секунды движется, пол секунды стоит.
        if (deltaTime - 1000 * parseInt(deltaTime / 1000) < 500) {
            x = (deltaTime - 1000 * parseInt(deltaTime / 1000)) / 1000 * canvas.width;
        }
        canvasContext.fillText(digit, x, y);
    }

    run() {
        let canvas = document.getElementById('countdown__canvas');
        let start = Date.now();


        let frameId = () => {
            let deltaTime = Date.now() - start;
            if (deltaTime < 3000) {
                // this.draw(canvas, deltaTime);=
                window.requestAnimationFrame(frameId);
            } else {
                this.router.go('/');
                return;
            }
        }
        frameId();
    }
}