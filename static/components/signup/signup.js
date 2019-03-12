export class SignUpComponent {
    constructor({
        el = document.body,
    } = {}) {
        this._el = el;
    }

    get data() {
        return this._data;
    }

    set data(d = []) {
        this._data = d;
    }

    render() {
        var templateScript = `
            <h1>Sign Up</h1>

            <form id="signup-form">
                <input
                    name="nickname"
                    type="text"
                    placeholder="Username"
                    class="signup_input">
                <input name="email" type="text" placeholder="Email" class="signup_input">
                <input
                    name="password"
                    type="password"
                    placeholder="Password"
                    class="signup_input"
                    title="Пароль должен состоять только из 3-20 латинских букв, цифр и символов пунктуации.">
                <input
                    name="password_repeat"
                    type="password"
                    placeholder="Repeat Password"
                    class="signup_input"
                    title="Пароли должны совпадать">
                <input
                    name="submit"
                    type="submit"
                    placeholder="undefined"
                    class="signup_btn">
                    
                <button data-section="menu" class="signup_btn">Назад</button>
            </form>
        `;

        const template = Handlebars.compile(templateScript);
        this._el.innerHTML = template();
    }
}


// board.data
// board.data = 2