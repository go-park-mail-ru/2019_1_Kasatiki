package main

import "models"

func (instance *App) CreateTables() (err error) {
	sql := `\
	CREATE EXTENSION IF NOT EXISTS CITEXT;
	DROP TABLE IF EXISTS users CASCADE;

	CREATE TABLE IF NOT EXISTS users (
	id 				CITEXT							NOT NULL	PRIMARY KEY,
	nickname		CITEXT			UNIQUE			NOT NULL,
	email			CITEXT			UNIQUE,
	password		TEXT							NOT NULL,
	points			BIGSERIAL,
	age				SMALLINT,
	imgurl			TEXT,
	region			TEXT,
	about			TEXT
);`
	_, err = instance.Connection.Exec(sql)
	return err
}

func (instance *App) InsertUser(user models.User) (err error) {
	sql := `
		INSERT INTO users (id, nickname, email, password, points, age, imgurl, region, about)
			VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9);
`
	_, err = instance.Connection.Query(sql, user.ID, user.Nickname, user.Email, user.Password, user.Points, user.Age, user.Age, user.ImgUrl, user.Region, user.About)
	return err
}

func (instance *App) UpdateUser(user models.User) (err error) {
	sql := `
		UPDATE users SET (nickname = $2, email = $3, password = $4, points = $5, age = $6, imgurl = $7, region = $8, about = $9) 
			WHERE id = $1;
`
	_, err = instance.Connection.Query(sql, user.ID, user.Nickname, user.Email, user.Password, user.Points, user.Age, user.Age, user.ImgUrl, user.Region, user.About)
	return err
}

func (instance *App) GetUser(id string) (user models.User, err error) {
	sql := `
		SELECT id, nickname, email, password, points, age, imgurl, region, about FROM users 
			WHERE id = $1;
`
	err = instance.Connection.QueryRow(sql, id).Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.Points, &user.Age, &user.Age, &user.ImgUrl, &user.Region, &user.About)
	return user, err
}
