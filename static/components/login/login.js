export class loginComponent {
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

    render(authStatus) {
        if (authStatus) {
            var templateScript = `
                <div>
                    <h1>Вы уже авторизованы</h1>
                    <button data-section="menu" class="btn">Назад</button>
                </div>
            `;
        } else {
            // ToDo, Tmrln: Тут надо что-то делать: либо у <form> ставить атрибуты action="/login", method="post" (потом будут проблемы с переадресацией),
            //              либо делать <div class="menu"> (обрати внимание на адресную строку, туда печатается гет запрос с формы - это вообще не круто)
            var templateScript = `
                <form id="login-form">
                    <h1>Login</h1>
                    <input
                        name="nickname"
                        type="text"
                        placeholder="Nickname"
                        class="login_input"
                        title='Никнейм должен состоять только из 3-20 латинских букв, цифр и символов "_" и "-", а также должен начинаться с буквы. Можно ввести email.'>
                    <input
                        name="password"
                        type="password"
                        placeholder="Password"
                        class="login_input"
                        title="Пароль должен состоять только из 3-20 латинских букв, цифр и символов пунктуации.">
                    <input name="submit" type="submit" class="login_btn">
                    <button data-section="menu" class="btn">Назад</button>
                </form>
            `;
        }
        // ToDo, Tmrln: У меня тут была логигка редиректа на /me (профиль), я делал все через станартные кнопки, на которые навешивал свои обрабтчики
        const template = Handlebars.compile(templateScript);
        this._el.innerHTML = template();
    }
}