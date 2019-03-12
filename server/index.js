'use strict';

const express = require('express');
const body = require('body-parser');
const cookie = require('cookie-parser');
const morgan = require('morgan');
const uuid = require('uuid/v4');
const path = require('path');
const app = express();

app.use(morgan('dev'));
app.use(express.static(path.resolve(__dirname, '..', 'static')));
app.use(body.json());
app.use(cookie());

const users = [
	{
		"ID":"2",
		"nickname":"tony",
		"email":"trendpusher@hydra.com",
		"password":"qwerty",
		"Points":-100,
		"Age":22,
		"ImgUrl":"test",
		"Region":"Moscow",
		"About":"В правой алейкум",
	},
	{
		"ID":"1",
		"nickname":"evv",
		"email":"onetaker@gmail.com",
		"password":"evv",
		"Points":100,
		"Age":23,"ImgUrl":"test",
		"Region":"Voronezh",
		"About":"В левой руке салам",
	},
];

// const users = [
// 		{
// 			email: 'email1',
// 			age: 21,
// 			score: 228,
// 		},
// 		{
// 			email: 'email2',
// 			age: 22,
// 			score: 228,
// 		},
// 		{
// 			email: 'email3',
// 			age: 23,
// 			score: 228,
// 		}
// 	]
const ids = {};

app.post('/signup', function (req, res) {
	console.log('post: signup');
	const password = req.body.password;
	const email = req.body.email;
	const age = req.body.age;

	if (
		!password || !email || !age ||
		!password.match(/^\S{4,}$/) ||
		!email.match(/@/) ||
		!(typeof age === 'number' && age > 10 && age < 100)
	) {
		return res.status(400).json({error: 'Не валидные данные пользователя'});
	}
	if (users[email]) {
		return res.status(400).json({error: 'Пользователь уже существует'});
	}

	const id = uuid();
	const user = {password, email, age, score: 0};
	ids[id] = email;
	users[email] = user;

	res.cookie('sessionid', id, {expires: new Date(Date.now() + 1000 * 60 * 10)});
	res.status(201).json({id});
});

app.post('/login', function (req, res) {
	console.log('post: /login');
	const email = req.body.email;
	const password = req.body.password;

	console.log(email);
	console.log(password);

	if (!password || !email) {
		return res.status(400).json({error: 'Не указан E-Mail или пароль'});
	}
	if (!users[email] || users[email].password !== password) {
		return res.status(400).json({error: 'Не верный E-Mail и/или пароль'});
	}


	const id = uuid();
	ids[id] = email;

	res.cookie('sessionid', id, {expires: new Date(Date.now() + 1000 * 60 * 10)});
	res.status(200).json({id});
});

app.get('/me', function (req, res) {
	console.log('get: /me');
	const id = req.cookies['sessionid'];
	const email = ids[id];
	if (!email || !users[email]) {
		return res.status(401).end();
	}

	users[email].score += 1;

	res.json(users[email]);
});

app.get('/users', function (req, res) {	
	console.log('get: /users');
	const scorelist = Object.values(users)
	.sort((l, r) => r.Points - l.Points)
	.map(user => {
		return user;
	});

	res.json(scorelist);
});

const port = process.env.PORT || 3000;

app.listen(port, function () {
	console.log(`Server listening port ${port}`);
});