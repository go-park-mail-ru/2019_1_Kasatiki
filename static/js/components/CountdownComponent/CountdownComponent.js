export default class CountdownComponent {
    render() {
        const templateScript = `
            <canvas id="countdown__canvas" class="canvas"></canvas>
        `;
        const template = Handlebars.compile(templateScript);
        return template();
    }

    canvasFullScreen() {
        let canvas = document.getElementById("countdown__canvas");
        canvas.height = document.body.clientHeight;
        canvas.width = document.body.clientWidth;
    }
}