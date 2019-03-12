export class boardComponent {
    constructor ({
        parentElement = document.body, 
    } = {}) {
        this._parentElement = parentElement;
    }

    get data() {
        return this._usersArr;
    }

    set data(usersArr = []) {
        this._usersArr = usersArr;
    }

    render(users) {
        console.log(users);
        // Итерируясь по юзерам, выводим строки таблицы
        const templateScript = `
            <table class="leadersTable" border="1" cellpadding="0" cellspacing="0">
                <thead>
                    <tr>
                        <th>Nickname</th>
                        <th>Region</th>
                        <th>Score</th>
                    </tr>
                </thead>
                <tbody>
                    {{#each .}}                  
                        <tr class="tr">
                            <td>{{nickname}}</td>
                            <td>{{Region}}</td>
                            <td>{{Points}}</td>
                        </tr>
                    {{/each}}
                </tbody>
            </table>
        `;

        // console.log(templateScript);
        const template = Handlebars.compile(templateScript);
        // console.log(template);
        this._parentElement.innerHTML += template(users); 
    }
}