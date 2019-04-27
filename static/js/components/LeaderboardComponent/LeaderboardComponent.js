const { NetworkHandler } = window;

export default class LeaderboardComponent {
    	/**
	 * Конструктор
	 * @param {Object} parentElement - Элемент DOM дерева,
	 * куда будет отрисовываться leader board.
	 */
	constructor(
		parentElement = document.body,
		usersPerPage = 5,
		totalPages = 2,
	) {
		// Поля таблицы
		this._parentElement = parentElement;
		this._usersArr = {};

		// Поля пагинатора
		this._usersPerPage = usersPerPage;

		this._pagesDict = {
			_currentPage: 1,
			_totalPages: totalPages,
		};

		// Блок пагинатора
		this._paginatorSection = this._parentElement;
	}

	// Методы пагинатора

	// Методы пагинатора рендарят борду
	/**
	 * Метод отрисовки предыдущей страницы.
	 */
	_getPrevPage() {
		if (this._pagesDict._currentPage > 1) 
		{ 
			this._pagesDict._currentPage--;

			// Формирую оффсет, который передается на бэк
			const offset = this._pagesDict._currentPage * this._usersPerPage;

			const that = this;

			NetworkHandler.doGet({
				callback(data) {
					// Принимаю ответ с бэка с отсортирваными пользователями, 
					// записываю его в _usersArr, отрисовываю борду
					that._usersArr = data;
					that.render();
				},
				path: `/api/leaderboard?offset=${offset}`,
			});
		}
	}

	/**
	 * Метод отрисовки следующей страницы.
	 */
	_getNextPage() {
		if (this._pagesDict._currentPage < this._pagesDict._totalPages) 
		{
			this._pagesDict._currentPage++;

			const offset = this._pagesDict._currentPage * this._usersPerPage;
	
			const that = this;
	
			NetworkHandler.doGet({
				callback(data) {
					// Принимаю ответ с бэка с отсортирваными пользователями, 
					// записываю его в _usersArr, отрисовываю борду
					that._usersArr = data;
					that.render();
				},
				path: `/api/leaderboard?offset=${offset}`,
			});
		} 
	}

	_renderPaginator() {
		// Шаблон без div, так как div прописан в шаблоне борды
		const templateScript = `
			<button class="prev leaderboard_page-button"><i class="fas fa-arrow-left"></i></button>
			<h1 class="leaderboard_page-pageNumber">{{_currentPage}}</h1>
			<button class="next leaderboard_page-button"><i class="fas fa-arrow-right"></i></button>
		`;
		
		const template = Handlebars.compile(templateScript);
		this._paginatorSection.innerHTML += template(this._pagesDict);

		const prevButton = this._parentElement.querySelector('.prev');
		const nextButton = this._parentElement.querySelector('.next');

		// Отрисовываю сначалу борду, создавая тем самым
		nextButton.addEventListener('click', () => {
			this._paginatorSection.innerHTML = '';
			this._getNextPage();
			this._renderPaginator();
		});

		prevButton.addEventListener('click', () => {
			this._paginatorSection.innerHTML = '';
			this._getPrevPage();
			this._renderPaginator();
		});
	}

	// Методы борды

	/**
	 * Геттер данных о пользователях.
	 */
	get users() {
		return this._usersArr;
	}

	/**
	 * Сеттер данных о пользователях.
	*/
	set users(users = []) {
		this._usersArr = users;
	}

	/**
	 * Метод для отрисоки leader board.
	 * @param {array} users - массив данных о пользователях.
	 */
	render() {
		this._parentElement.innerHTML = '';
		// Итерируясь по юзерам, выводим строки таблицы
		// Зарание создал место для пагинатора: <div class="paginatorSection"></div>
		const templateScript = `
		<div class="leaderboard">
			<h1 class="leaderboard__title">Leaderboard</h1>
			<div class="board">
				{{#each .}} 
				<div class="board__player">
					<h3 class="board__player-place">{{@index}}</h3>
					<h3 class="board__player-nickname">{{nickname}}</h3>
					<h3 class="board__player-points">{{points}}</h3>
				</div>
				{{/each}} 
			</div>
			<div class="paginator-section"></div>
			<button class="leaderboard_back-button" href="/" data-section="menu"><i href="/" class="fas fa-undo-alt"></i>
			</button>
		</div>
		`;

		const template = Handlebars.compile(templateScript);
		this._parentElement.innerHTML += template(this._usersArr, this._pagesDict._currentPage);

		// Вытаскиваю из DOM'а <div class="paginatorSection"></div>, записываю его в 
		// _paginatorSection: 
		this._paginatorSection = this._parentElement.querySelector('.paginator-section');

		// Рендерю пагинатор в _paginatorSection
		this._renderPaginator();
	}
}