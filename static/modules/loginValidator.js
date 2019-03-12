export class ValidModule {

    constructor() {
        this._authStatus = false;
    }

    set status(validUser) {
        this._authStatus = validUser;
    }

    validUser() {
        const that = this;
        // let promise = new Promise((resolve, reject) => {
        //     AjaxModule.doGet({	
        //         callback(xhr) {
        //             const res = JSON.parse(xhr.responseText);
        //             console.log('validUser:', res.is_auth);
        //             that._authStatus = res.is_auth;
        //             if (res.is_auth) {
        //                 resolve('authorized');
        //             } else {
        //                 reject('unauthorized');
        //             }
        //         },
        //         path : '/isauth',
        //     });
        // });

        

        async function getStatus() {
            try {
                const res = await fetch ('/isauth');
                return await res.json();
            } catch(err) {
                console.err(err);
            }
        }

        console.log(getStatus());
        
        // promise
        //     .then((resolveStatus) => {
        //         console.log(resolveStatus);
        //         that._authStatus = true;
        //     }, (rejectStatus) => {
        //         console.log(rejectStatus)
        //     })
            // .then(() => {

            // }, () => {

            // });
    }

    get status() {
        return this._authStatus;
    }

}
