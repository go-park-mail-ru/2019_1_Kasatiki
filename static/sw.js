// const CACHE_NAME = 'advhater-cache';

// const cacheUrls = [
//     '/',
//     'api/isauth',
// ];

// this.addEventListener('install', (event) => {
//     event.waitUntil(
//         caches
//             .open(CACHE_NAME)
//             .then((cache) => {
//                 return cache.addAll(cacheUrls);
//             })
//     );
// });

// this.addEventListener('fetch', (event) => {
//     event.respondWith(
//         caches
//             .match(event.request)
//             .then((cachedResponse) => {
//                 if (!navigator.onLine && event.request === 'api/isauth') {
//                     return request.status;
//                 }
//                 if (!navigator.onLine && cachedResponse) {
//                     return cachedResponse;
//                 }

//                 return fetch(event.request)
//                     .then(response => caches
//                         .open(CACHE_NAME)
//                         .then((cache) => {
//                             if (event.request.method === 'GET') {
//                                 cache.put(event.request, response.clone());
//                             }
//                             return response;
//                         }));
//             })
//     );
// });