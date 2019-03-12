'use strict'

import {loginComponent} from './components/login/login.js';
import {SignUpComponent} from './components/signup/signup.js';
import {paginationComponent} from './components/pagination/pagination.js';
import {boardComponent} from './components/board/board.js';
import {helperComponent} from './components/helper/helper.js';
import {profileComponent} from './components/profile/profile.js';
import {ValidModule} from './modules/loginValidator.js';

const {AjaxModule} = window;
var authValid = new ValidModule();

authValid.status = false;

console.log("client started");

// Разименовывем необходимые элементы DOM'a

const app = document.getElementById("application");
const helper = new helperComponent();

const main = document.createElement('div');
main.className = "main";
app.appendChild(main);

// Валидация данных в форме
function validateNickname(nickname = "") {
	return /^[a-zA-z]{1}[a-zA-Z1-9]{2,20}$/.test(nickname);
}

function validateEmail(email = "") {
	return /^[0-9a-z-\.]+\@[0-9a-z-]{2,}\.[a-z]{2,}$/.test(email);
}

function validateLogin(login = "") {
	if (login.search(/@/i) !== -1) {
		return validateEmail(login);
	}
	return validateNickname(login);
}

function validatePassword(password = "") {
	if (password.length < 3 || password.length > 20) {
		return false;
	}
	return /^[a-zA-Z](.[a-zA-Z0-9_-]*)$/.test(password);
}

// Основные функции, отвечают за логику работы фронта/генерирование контента и тд
function createMenu() {
	
	console.log('auth status: ', authValid.status);
	
	// Разное меню для залогиненых/разлогиненых пользователей
	if (authValid.status) {
		var menuItems = {
			game: 'Играть',
			leaderboard: 'Таблица лидеров',
			// about: 'О приложении',
			profile: 'Профайл',
			logout: 'Выйти',
		};
	} else {
		var menuItems = {
			signup: 'Регистрация',
			login: 'Логин',
			game: 'Играть',
			leaderboard: 'Таблица лидеров',
			// about: 'О приложении',
			profile: 'Профайл',
		};
	}

	main.innerHTML = '';

	const menu = document.createElement('div');
	menu.className = 'menu';
	main.appendChild(menu);

	// Добаляем название блока
	helper.createTitle(menu, 'Main menu');

	// Создаем кнопки
	Object.keys(menuItems).forEach( (key) => {
		helper.createButton(menu, key, 'btn', menuItems[key]);
	});
}

// ToDo: Tmrln: нужно поправить
function createLogin() {
	// ToDo, Tmrln: написать (если нужно) пояснения к блоку
	main.innerHTML = '';

	const signInSection = document.createElement('div');
	signInSection.className = 'menu';
	signInSection.dataset.section = 'login';

	main.appendChild(signInSection);

	const profile = new profileComponent(signInSection);
	const login = new loginComponent({el: signInSection});

	login.render(authValid.status);

	// ToDo, Tmrln: зачем алерты? оч мешает
	let id = -1;
	const signInChildNodes = signInSection.childNodes;
	console.log(signInSection);
	for (let i = 0; i < signInChildNodes.length; i++) {
		if ("login-form".localeCompare(signInChildNodes[i].id) === 0) {
			id = i;
		}
	}
	if (id === -1) {
		console.log("form id changed!!!");
		return;
	}
	const form = signInChildNodes[id];

	// ToDo, Tmrln: у меня логика лежит в компонентах, не знаю, на сколько это правильно
	// 				(скорее всего неправильно :^( )
	form.addEventListener('submit', function (event) {
		// ToDo, Tmrln: Вот это вообще непонятно, ты пользуешься сабмитом с формы, при этом удаляешь все ее 
		//				первичные обработчики, при этом ты не можешь сделать из-за этого редирект на /me, 
		// 				нужно 100% менять (можно взять мой старый коммит)
		event.preventDefault();

		const fieldsName = [
			'nickname',
			'password',
		]
		
		// ToDo, Tmrln: обработчики не работают: при регистрации пароль менее 3 символов не пускает,
		// 				однако, при логине пароли в 1 символ проходят 
		const validationFunction = {
			nickname: validateLogin,
			password: validatePassword,
		}

		let validationError = false;
		authValid.status = true;
		fieldsName.forEach(function(fieldName) {
			const field = form.elements[fieldName];
			field.className = "login_input";

			if (!validationFunction[ fieldName ]( field.value )) {
				field.className = "login_input login_bad_input";
				validationError = true;
			} else {
				field.className = "login_input login_good_input";
			}
		});

		if (validationError) {
			return;
		}


		console.log('changed:',authValid.status);

		AjaxModule.doPost({
			callback(xhr) {
				const answer = JSON.parse(xhr.responseText);
				if (typeof(answer['Error']) === "undefined") {
					main.innerHTML = '';
					console.log(answer);
					console.log("OK");
					console.log(form.elements[ 'nickname' ].value);

					// вызываю сетер
					authValid.status = true;


					// Мы отказались от никнейма втроым аргументом и теперь делаем запрос на me
					profile.createProfile(authValid.status);
				} else {
					alert(answer['Error']);
					console.log("WTF?");
				}
			},
			path: '/login',
			body: {
				nickname: form.elements[ 'nickname' ].value,
				password: form.elements[ 'password' ].value,
			},
		});
	});

}

function createSignup() {
	main.innerHTML = '';

	const signUpSection = document.createElement('section');
	signUpSection.className = 'menu';	
	signUpSection.dataset.section = 'sign_up';

	const profile = new profileComponent(signUpSection);
	const signUp = new SignUpComponent({
		el: signUpSection,
	});

	signUp.render();

	let id = -1;
	const signUpChildNodes = signUpSection.childNodes;
	for (let i = 0; i < signUpChildNodes.length; i++) {
		if ("signup-form".localeCompare(signUpChildNodes[i].id) === 0) {
			id = i;
		}
	}
	if (id === -1) {
		alert("form id changed!!!");
		return;
	}

	const form = signUpChildNodes[id];

	form.addEventListener('submit', function (event) {
		event.preventDefault();

		const email = form.elements[ 'email' ].value;
		const nickname = form.elements[ 'nickname' ].value;
		const password = form.elements[ 'password' ].value;

		const fieldsName = [
			'email',
			'nickname',
			'password',
		]

		
		// ToDo, Tmrln: обработчики не работают: при регистрации пароль менее 3 символов не пускает,
		// 				однако, при логине пароли в 1 символ проходят 
		const validationFunction = {
			email: validateEmail,
			nickname: validateLogin,
			password: validatePassword,
		}

		let validationError = false;
		fieldsName.forEach(function(fieldName) {
			const field = form.elements[fieldName];
			field.className = "signup_input";

			if (!validationFunction[ fieldName ]( field.value )) {
				field.className = "signup_input signup_bad_input";
				validationError = true;
			} else {
				field.className = "signup_input signup_good_input";
			}
		});

		if (password !== form.elements[ 'password_repeat' ].value) {
			form.elements[ 'password_repeat' ].className = "signup_input signup_bad_input";
			validationError = true;
		} else {
			form.elements[ 'password_repeat' ].className = "signup_input signup_good_input";
		}

		if (validationError) {
			return;
		}

		AjaxModule.doPost({
			callback(xhr) {
				const responseAnswer = xhr.responseText;
				const answer = JSON.stringify(responseAnswer);

				if (typeof(answer['Error']) === "undefined") {
					main.innerHTML = '';
					createMenu();					
					console.log("OK");
				} else {
					alert(answer['Error']);
				}
			},
			path: '/signup',
			body: {
				nickname: nickname,
				email: email,
				password: password
			}
		});
	});

	main.appendChild(signUpSection);
}

function createProfile() {
	// Создаем родительский элемент
	const profileSection = document.createElement('div');
	profileSection.className = 'menu';
	profileSection.dataset.section = 'profile';

	// Добавляем его к main'у
	main.appendChild(profileSection);

	// Создаем объект класса profile...
	const profile = new profileComponent(profileSection);
	
	// Запрос на me
	AjaxModule.doGet({	
		callback(xhr) {
			const user = JSON.parse(xhr.responseText);
			profile.render(authValid.status ,user);
			console.log(user);
		},
		path : '/me',
	});
}

// Примитивная реализация геймплея, при выходе нужно обновить страницу
// пока не трогаем 
function createGame() {
	main.innerHTML = '';

	const menu = document.createElement('div');
	menu.className = 'gameBlock';
	main.appendChild(menu);

	const canvas = document.createElement('canvas');
	canvas.id = 'gameCanvas';
	canvas.width = 700;
	canvas.height = 700;
	menu.appendChild(canvas);

	const buf = document.getElementById('gamejs');
	if (buf !== null) {
		alert('got!');
	}

	const gameLogic = document.createElement('script');
	gameLogic.id = "gamejs";
	gameLogic.src = '/game.js'
	main.appendChild(gameLogic);

	createButton(menu, 'menu', 'btn', 'Back');
}

function createLeaderboard(users) {	
	main.innerHTML = '';

	AjaxModule.token = 'dadaad';

	// Создаем родительский блок
	const leaderboard = document.createElement('div');
	leaderboard.dataset.section = 'leaderboard';
	leaderboard.className = 'menu';

	helper.createTitle(leaderboard, 'Leaderboard');

	if (users) {
		const paginator = new paginationComponent({parentElement : leaderboard,
													// Количество пользователей на странице
													usersPerPage : 1,
													// Можно прикрутить максимально количество страниц, но работает и без этого
													totalPages : 3,
													});
		const board = new boardComponent({parentElement : leaderboard});

		// Сетим дату
		board.data = JSON.parse(JSON.stringify(users));

		// Отрисовываем борду, затем пагинатор
		board.render(users); 
		paginator.renderPaginator(3);
	// Если юзеры не пришли, еще раз стучимся
	} else {
		console.log('data loading, please wait');

		AjaxModule.doGet({	
			callback(xhr) {
				const user = JSON.parse(xhr.responseText);
				main.innerHTML = '';
				createLeaderboard(user);
			},
			path : '/leaderboard',
		});
	}

	main.appendChild(leaderboard);
}

function createLogout() {
	authValid.status = false;
	createMenu();
}

// Главный блок: создаем меню 
createMenu();

// Карта функций, сюда будем подгруать новые функции
const functions = {
	menu: createMenu,

	// Menu elements
	signup: createSignup,
	login: createLogin,
	game: createGame,
	leaderboard: createLeaderboard,
	profile: createProfile,
	// about: createAbout,
	logout: createLogout,


	// Other functions
	title: helper.createTitle,
};

// Обработчик всех евентов в DOM'е
app.addEventListener('click', (evt) => {

	const target = evt.target;

	// console.log('click on ', target, 'datasec: ', target.dataset.section);

	// Если target является кнопкой 
	if (target instanceof HTMLButtonElement) {
		if (target.dataset.section in functions) {
			// Убираем все стандартные обработчики	
			evt.preventDefault();

			const section = target.dataset.section;

			functions[section]();
		}
	}
});