/* eslint-disable import/prefer-default-export */

/**
 * Прототип класса с модулем проверки авторизации пользователя.
 */
const noop = () => null;

/**
 * Класс с методами отправки AJAX-запросов на сервер
 */
export default class NetworkHandler {
	_send({
		callback = noop,
		path = '/',
		method = 'GET',
		body,
	} = {}) {
		const options = {
			method,
			// Настройка CORS
			headers : {
				// Запрещаем открытие iframe на сайте
				'X-Frame-Options' : 'DENY',
				'Content-Type' : 'application/json',
				'Accept':  'application/json',
				// // Мы разворачиваемся на этом домене
				// 'Access-Control-Allow-Origin' : 'http://advhater.ru/',
				'Access-Control-Allow-Credentials' : true,
				// Допускаем только GET, POST, DELETE, HEAD запросы
				'Access-Control-Request-Method' : 'POST, GET, PUT, DELETE, HEAD,',
				// Для "непростых запросов"
				// 'Origin' : '',
			},
			// credentials: "same-origin",
			credentials : "include",
			// mode : 'cors',
			cache : 'default',
			body,
		}
		
		fetch(path, options)
			.then(function (response) {
				console.log(path, response);
				console.log('headers', response.headers);
				if (response.status === 200) {
					console.log('network 200 success', response);
					return response.json();
				} else if (response.status === 201) {
					console.log ('network 201 success', response);
					console.log (response.headers.get('Set-Cookie'));
					console.log(JSON.stringify(response.headers));
					// document.cookie = response.headers['Set-Cookie'];
					return 201;
				} else if (response === undefined) {
					return 404;
					// throw new Error('Wrong network response');
				}
				return response.status;
				
			})
			.then(function (data) {
				callback(data);
			})
			// .catch((error) => {
			// 	// console.log('status', response.status);
			// 	callback('err');
			// })
	}

	doGet({
		callback = noop,
		path = '/',
	} = {}) {
		this._send({
			callback,
			path,
			method : 'GET',
		}) 
	}

	doHead({
		callback = noop,
		path = '/',
	} = {}) {
		this._send({
			callback,
			path,
			method : 'HEAD',
		}) 
	}

	doPost({
		callback = noop,
		path = '/',
		body = {},
	} = {}) {
		this._send({
			callback,
			path,
			method : 'POST',
			body,
		}) 
	}

	doPut({
		callback = noop,
		path = '/',
		body = {},
	} = {}) {
		this._send({
			callback,
			path,
			method : 'PUT',
			body,
		}) 
	}

	doDelete({
		callback = noop,
		path = '/',
		body = {},
	} = {}) {
		this._send({
			callback,
			path,
			method : 'DELETE',
			body,
		}) 
	}
}

window.NetworkHandler = new NetworkHandler;