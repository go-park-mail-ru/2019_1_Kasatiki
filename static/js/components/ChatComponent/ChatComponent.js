export default class ChatComponent {
    render(isPage) {
        let templateScript = ``

        if (isPage === false) {
            templateScript =`
            <div class="chat">
                <div class="chat__chatbox"></div>
                <form class="chat__form">
                    <input type="text" class="chat__input">
                    <input type="submit" class="chat__submit" value="Allo">
                </form>
            </div>
            `;
        } else {
            templateScript =`
            <a href="/" class="chat__menu-link">Menu</a>
            <div class="chat-section">
                <div class="chat">
                    <div class="chat__chatbox"></div>
                    <form class="chat__form">
                        <input type="text" class="chat__input">
                        <input type="submit" class="chat__submit" value="Allo">
                    </form>
                </div>
            </div>
            `;
        }

		const template = Handlebars.compile(templateScript);		
        return template();
	}
}