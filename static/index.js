import Router from './js/modules/Router.js';
import './js/modules/NetworkHandler.js';
// import './js/modules/ws.js';

import MenuView from './js/views/MenuView.js';
import LoginView from './js/views/LoginView.js';
import LogoutView from './js/views/LogoutView.js';
import SignupView from './js/views/SignupView.js';
import CountdownView from './js/views/CountdownView.js';
import LeaderboardView from './js/views/LeaderboardView.js';
import GameView from './js/views/GameView.js';
import PaymentsView from './js/views/PaymentsView.js';

if ('serviceWorker' in navigator) {
    navigator.serviceWorker.register('sw.js')
    .then((reg) => {
        console.log('sw reg success:', reg);
    })
    .catch((err) => {
        console.error('sw reg err:', err);
    });
}

const router = new Router(document.body);
router.add('/', MenuView);
router.add('/countdown', CountdownView);
router.add('/login', LoginView);
router.add('/logout', LogoutView);
router.add('/signup', SignupView);
router.add('/leaderboard', LeaderboardView);
router.add('/play', GameView);
router.add('/payment', PaymentsView);
router.run();