export class helperComponent {

    createInput(object, name, type, placeholder, className) {
        const optionsDict = {
            name: name,
            type: type,
            placeholder: placeholder,
            className: className,
        }

        const templateScript = `
            <input name="{{name}}" type="{{type}}" placeholder="{{placeholder}}" class="{{className}}">
        `;

        // console.log(templateScript);
        const template = Handlebars.compile(templateScript);
        // console.log(template);
        object.innerHTML += template(optionsDict); 
    }

    createButton(object, sectionName, className, textContent) {
        const optionsDict = {
            sectionName: sectionName,
            textContent: textContent,
            className: className,
        }

        const templateScript = `
            <button data-section="{{sectionName}}" class="{{className}}">{{textContent}}</button>
        `;

        // console.log(templateScript);
        const template = Handlebars.compile(templateScript);
        // console.log(template);
        object.innerHTML += template(optionsDict); 
    }

    createTitle(parentObject, text) {
        const optionsDict = {
            text: text,
        }

        const templateScript = `
            <div class="title">
                <h1>{{text}}</h1>
            </div>
        `;

        const template = Handlebars.compile(templateScript);
        parentObject.innerHTML += template(optionsDict); 
    }

}