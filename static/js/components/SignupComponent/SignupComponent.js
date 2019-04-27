export default class SignupComponent {
    /**
     * Функция, возвращающая строку html с формой логина.
     * @param {boolean} isAuth - Статус авторизации пользователя.
     */
    render(isAuth) {
        let templateScript = '';
		if (isAuth) {
			templateScript = `
                <h1 class="title">Вы уже авторизованы</h1>
				<a href="/" class="btn">Назад</a>
            `;
		} else {
			templateScript = `
                <div class="signup">
                    <h1 class="title">Sign Up</h1>
                    <form id="signup-form">
                        <input
                            name="nickname"
                            type="text"
                            placeholder="Username"
                            class="signup__input"
                            autocomplete="on"
                            >
                        <input
                            name="email"
                            type="text"
                            placeholder="Email"
                            class="signup__input"
                            autocomplete="on"
                            >
                        <input
                            name="password"
                            type="password"
                            placeholder="Password"
                            class="signup__input"
                            autocomplete="on"
                            >
                        <input
                            name="password_repeat"
                            type="password"
                            placeholder="Repeat Password"
                            class="signup__input"
                            >
                        <div id="signup-form__error-text-field" class="signup_error_text"></div>  
                        <div class="signup__btn-section">
                            <button href="/" class="signup-btn"><i href="/" class="fas fa-undo-alt"></i></button>
                            <button href="/signupuser" class="signup-btn"><i href="/signupuser" class="fas fa-angle-double-right"></i></button>
                        </div>
                    </form>
                </div>
            `;
		}
        const template = Handlebars.compile(templateScript);
        return template();
    }
    
    setOnChangeListener(input) {
		input.addEventListener("input", () => {
			this.goodField(input);
		});
	}
	
	get form() {
        return document.querySelector('#signup-form');
    }

	setErrorText(text) {
		if (typeof text !== 'string') {
			console.log('NOTE: text is not string');
			return;
		}
		let errorField = document.getElementById('signup-form__error-text-field');
		errorField.textContent = text;
	}

	goodField(input) {
		input.setCustomValidity("");
	}

	errorField(input) {
		input.setCustomValidity("-_-");
	}
}