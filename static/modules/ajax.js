(function() {
	const noop = () => null;

	class AjaxModule {
		_ajax({
			addHeader = {},
			callback = noop,
			method = 'GET',
			path = '/',
			body = {},
		} = {}) {
			const xhr = new XMLHttpRequest();
			xhr.open(method, path, true);
			xhr.withCredentials = true;

			// Настравиваем безопасность 
			// Запрет на открытие iframe (любых)
			xhr.setRequestHeader('X-Frame-Options', 'DENY');
			// Работа с токеном (CSRF)
			if (this._csrfToken !== undefined) {
				xhr.setRequestHeader('X-CSRF-Token', this._csrfToken);
			}

			if (body) {
				xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');
			}

			xhr.onreadystatechange = function () {
				if (xhr.readyState !== 4) {
					return;
				}

				callback(xhr);
			};

			if (body) {
				xhr.send(JSON.stringify(body));
			} else {
				xhr.send();
			}
		}

		set token(tokenValue) {
			this._csrfToken = tokenValue;
		}

		doGet({
			callback = noop,
			path = '/',
			body = {},
		} = {}) {
			this._ajax({
				callback,
				path,
				body,
				method: 'GET',
			});
		}

		doPost({
			callback = noop,
			path = '/',
			body = {},
		} = {}) {
			this._ajax({
				callback,
				path,
				body,
				method: 'POST',
			});
		}
		
	}
	window.AjaxModule = new AjaxModule();
})();