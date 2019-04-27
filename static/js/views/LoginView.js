import BaseView from './View.js';

import LoginComponent from '../components/LoginComponent/LoginComponent.js';
import Validation from '../modules/Validation.js';

const { NetworkHandler } = window;

/**
 * Класс с отрисовкой формы логина.
 */
export default class LoginView extends BaseView {
    constructor() {
        super(...arguments);
        this.Validation = new Validation();
        this.LoginComponent = new LoginComponent();
        this.initSpecialRoutes();
    }

    show() {
        let that = this;
        NetworkHandler.doGet({
            // eslint-disable-next-line no-unused-vars
            callback(data) {
                let isAuth = data['is_auth'];
                that.root.innerHTML = that.LoginComponent.render(isAuth);
                that.LoginComponent.setOnChangeListener(that.LoginComponent.login);
                that.LoginComponent.setOnChangeListener(that.LoginComponent.password);
            },
            path: '/api/isauth',
        });
    }
    
    /**
     * Функция инициализирует "специальные" роутеры. Подробнее в файле View.js
     * или Router.js : go(path)
     */
    initSpecialRoutes() {
        this.specialRoutes['/authorizeuser'] = this.authorizeUser;
    }

    /**
     * Функция вызывается Роутером при нажатии на <a href=/authori...></a>
     * @param {Object} that - ссылка на инстанс LoginView
     */
    authorizeUser(that) {
        let login = that.LoginComponent.login;
        let password = that.LoginComponent.password;

        let isValid = that.validateValue(login, that.Validation.checkNickname, that.LoginComponent);
        if (!isValid) {
            return;
        }
        isValid = that.validateValue(password, that.Validation.checkPassword, that.LoginComponent);
        if (!isValid) {
            return;
        }

        const payload = {
            nickname : login.value,
            password : password.value,
        }

        NetworkHandler.doPost({
            callback(data) {
                console.log('data in login', data);
                if (data === 201) {
                    console.log('doc cookie: ',document.cookie);
                    // console.log('data in login:', data);
                    // that.router.handle('profile', data);
                    that.router.go('/');
                } else {
                    that.LoginComponent.setErrorText(data.Error);
                }
            },
            path: '/api/login',
            body: JSON.stringify( payload ),
        });
    }

    /**
     * Функциия, проверяющая на валидность поле input.
     * Если поле не валидное, то вызывает метод errorField, которое подсвечивает поле.
     * Если все ок - убирает подсветку.
     * @param {HTMLelement} input
     * @param {function} validationFunc 
     * @param {Object} LoginComponent 
     */
    validateValue(input, validationFunc, LoginComponent) {
        let validationMessage = validationFunc(input.value);
        if (validationMessage.localeCompare('OK') !== 0) {
            LoginComponent.setErrorText(validationMessage,);
            LoginComponent.errorField(input);
            return false;
        }
        LoginComponent.goodField(input);
        return true;
    }
}