export default class MenuComponent {
    render(data) {
        console.log('profile Data:', data)
        let templateScript = ``;
        if (typeof(data) == 'object') {
            templateScript = `
            <div class="menu">
                <h1 class="menu__title">ADV<a class="hate">HATER</a></h1>
                <div class="main-bar">
                    <button class="main-bar__item menu_btn" href="/leaderboard" data-title="LEADERBOARD" data-section="leaderboard"><i class="fas fa-list-ul"></i></button>
                    <button class="main-bar__item menu_btn" href="/" data-title="PAYOUT" data-section="payout"><i class="fas fa-wallet"></i></button>
                    <button class="main-bar__item menu_btn" href="/shop" data-title="SHOP" data-section="shop"><i class="fas fa-shopping-cart"></i></button>
                </div>
                <div class="menu__profile">
                    <img src="{{imgurl}}" class="profile-avatar"></img>
                    <div class="profile-info">
                        <h4 class="profile-nickname">{{nickname}}</h4>
                        <h4 class="profile-score">{{points}} $</h4>
                    </div>
                    <div class="menu__profile-buttons-section" id="menu__profile-buttons-section">
                        <button class="profile-button">Edit</button>
                        <button href="/logout" class="profile-button">Logout</button>
                    </div>
                </div>
                <button href="/play" class="menu__play-btn data-section="play"><i class="fas fa-play"></i></button>
                <div class="options">
                    <a href="#" class="options__link">Report bug</a>
                    <a href="#" class="options__link">Help</a>
                    <a href="#" class="options__link">Hate with us!</a>
                </div>  
            </div>
            <div class="chat">
                <div class="chat__chatbox"></div>
                <form class="chat__form">
                    <input type="text" class="chat__input">
                    <input type="submit" class="chat__submit" value="Allo">
                </form>
            </div>
        `;
        } else {
            templateScript = `
                <div class="chose">
                    <h1 class="chose__title">ADVHATER</h1>
                    <div class="main-bar">
                        <button href="/signup" class="main-bar__item chose_btn" data-title="SIGNUP" data-section="signup"><i href="/signup" class="fas fa-user-plus"></i></button>
                        <button href="#" class="main-bar__item chose_btn" data-title="QUICK PO" data-section="chose"><i href="#" class="fas fa-wallet"></i></button>
                        <button href="/login" class="main-bar__item chose_btn" data-title="LOGIN" data-section="login"><i href="/login" class="fas fa-sign-in-alt"></i></button>
                    </div> 
                </div>
            `;
        }

        const template = Handlebars.compile(templateScript);
		return template(data);
    }

    _doSmt() {
        console.log("i am not private");
    }
}