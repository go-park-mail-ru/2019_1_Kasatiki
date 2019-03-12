const {AjaxModule} = window;
import {boardComponent} from '../board/board.js';

// Handlebars.registerHelper('iff', function(a, operator, b, opts) {
//     let bool = false;
//     switch(operator) {
//         case '!=':
//             bool = a != b;
//             break;
//         case '==':
//             bool = a == b;
//             break;
//         case '>':
//             bool = a > b;
//             break;
//         case '<':
//             bool = a < b;
//             break;
//         default:
//             throw "Unknown operator " + operator;
//     }
 
//     if (bool) {
//         return opts.fn(this);
//     } else {
//         return opts.inverse(this);
//     }
// });

export class paginationComponent {
    constructor({
        parentElement = document.createElement('div.main'),
        usersPerPage = 1,
        totalPages = 1,
    } = {}) {
        this._parentElement = parentElement;

        this._usersPerPage = usersPerPage;

        this._pagesDict = {
            _currentPage: 1,
            _totalPages: totalPages,
        };

    }

    _getPrevPage() {
        if (this._pagesDict._currentPage > 1)
            this._pagesDict._currentPage --;

        const offset = this._pagesDict._currentPage*this._usersPerPage;

        const board = new boardComponent({parentElement : this._parentElement});

        const that = this;

        AjaxModule.doGet({	
            callback(xhr) {
                let res = JSON.parse(xhr.responseText); 
                that._parentElement.innerHTML='';
                board.render(res);
                that.renderPaginator();
            },
            path : '/leaderboard?offset=' + offset,
        });
    }

    _getNextPage() {
        if (this._pagesDict._currentPage < this._pagesDict._totalPages)
            this._pagesDict._currentPage ++;

        const offset = this._pagesDict._currentPage*this._usersPerPage;
        console.log(offset);

        const board = new boardComponent({parentElement : this._parentElement});

        const that = this;

        AjaxModule.doGet({	
            callback(xhr) {
                let res = JSON.parse(xhr.responseText); 
                that._parentElement.innerHTML='';
                board.render(res);
                that.renderPaginator();
            },
            path : '/leaderboard?offset=' + offset,
        });
    }


    renderPaginator() {
        const templateScript = `
            <div class="paginatorBox">
                <p class="prev"><</p>
                <p class="next">></p>
            </div>

            <button data-section="menu" class="btn">Назад</button>
        `;


        // console.log(templateScript);
        let template = Handlebars.compile(templateScript);
        console.log(this._pagesDict);
        this._parentElement.innerHTML += template(this._pagesDict); 

        let pageBox = this._parentElement.querySelector(".paginatorBox");
        // console.log(pageBox);
        let prevButton = this._parentElement.querySelector(".prev");
        let nextButton = this._parentElement.querySelector(".next");

        nextButton.addEventListener('click', () => {
            pageBox.innerHTML = '';
            this._getNextPage();
            this.renderPaginator();
        });


        prevButton.addEventListener('click', () => {
            pageBox.innerHTML = '';
            this._getPrevPage();
            this.renderPaginator();
        });

    }
}


