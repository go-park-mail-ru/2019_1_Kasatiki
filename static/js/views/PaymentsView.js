import BaseView from './View.js';

import PaymentsComponent from '../components/PaymentsComponent/PaymentsComponents.js'

const { NetworkHandler } = window;

export default class PaymentsView extends BaseView {
    constructor() {
        super(...arguments);
        this.PaymentsComponent = new PaymentsComponent();
    }

    show() {
        let that = this;
        NetworkHandler.doGet({
            callback(data) {
                console.log('data',data);
                if (typeof(data) === 'object') {
                    console.log("IM IN",data.status);
                    that.root.innerHTML = that.PaymentsComponent.render(true);
                } else {
                    that.root.innerHTML = that.PaymentsComponent.render(true);
                }
                that.initSpecialRoutes();
            },
            path: '/api/isauth',
        });
    }

    initSpecialRoutes() {
        this.specialRoutes['/payout'] = this.send;
    }

    send() {
        const form = document.querySelector('.payments__input-section');

        let payload = {
            phone : form.phone.value,
            amount : form.amount.value,
        }

        NetworkHandler.doPost({
            callback(data) {
                console.log('Success:',data);
                if (data === 201) {
                    that.router.go('/');
                }
            },
            path: '/api/payments',
            body: JSON.stringify( payload ),
        });
    }
}