export default class ShopComponent {

    constructor(
        root,
        weapons) {
        
        this.weapons = weapons;
        this.root = root;
    }

    render() {
        const templateScript = `
            <div class="shop-background">
                <div class="shop-container">
                    <div class="shop__menu">
                        {{#each .}}
                        <div class="shop__menu-item" data-section="{{id}}"><img src="{{icon}}" alt="{{name}} data-section="{{id}}"></div>
                        {{/each}}
                    </div>
                    <div class="weapon__about">
                    </div>
                </div>
            </div>
        `
        const template = Handlebars.compile(templateScript);		
        this.root.innerHTML = template(this.weapons);
    }

    renderWeaponInfo(id) {
       const about = document.querySelector('.weapon__about');

       const templateScript = `
            <h1 class="weapon__about-name">{{name}}</h1>
            <div class="weapon__about-main">
                <img src="{{icon}}" alt="weapon" class="weapon__about-main-image">
                <button class="weapon__about-main-purchase" data-section="{{id}}">Buy</button>
            </div>
            <div class="weapon__about-info">
                <div class="weapon__about-info-property">
                    <div class="weapon__about-info-property-item">
                        <img src="../../../icons/property/cost.svg" alt="cost">
                        <h2>{{cost}}</h2>
                    </div>
                    <div class="weapon__about-info-property-item">
                        <img src="../../../icons/property/damage.svg" alt="damage">
                        <h2>{{damage}}</h2>
                    </div>
                    <div class="weapon__about-info-property-item">
                        <img src="../../../icons/property/firerate.svg" alt="firerate">
                        <h2>{{fireRate}}</h2>
                    </div>
                </div>
            </div>
            <div class="weapon__about-info-description">
                <h3>{{about}}</h3>	
            </div>
        `;

        const template = Handlebars.compile(templateScript);	
        about.innerHTML = '';	
        about.innerHTML = template(this.weapons[id]);
        
    }
    
}