package dbhandler

import (
	"2019_1_Kasatiki/pkg/models"
	"fmt"
	"github.com/jackc/pgx"
	"math/rand"
	"time"
)

func RandStr(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type DBHandler struct {
	Connection *pgx.Conn
}

func (instance *DBHandler) dataCreating(lim int) {
	var u models.User
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < lim; i++ {
		u.Nickname = RandStr(10)
		u.Password = RandStr(10)
		u.Email = RandStr(10) + "@ya.ru"
		u.Points = r1.Intn(1000)
		instance.InsertUser(u)
	}
}

func (instance *DBHandler) msgCreating(lim int) {
	var msg models.Message
	//s1 := rand.NewSource(time.Now().UnixNano())
	//r1 := rand.New(s1)
	for i := 0; i < lim; i++ {
		msg.Nickname = RandStr(10)
		msg.Body = RandStr(50)
		msg.Imgurl = "advhater.ru/" + RandStr(10) + ".jpeg"
		msg.Timestamp = fmt.Sprintf("%02d:%02d:%02d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
		instance.InsertMessage(msg)
	}
}

func (instance *DBHandler) CreateTables() (err error) {
	sql := `
	CREATE EXTENSION IF NOT EXISTS CITEXT;
	DROP TABLE IF EXISTS users CASCADE;

	CREATE TABLE IF NOT EXISTS users (
	id 				BIGSERIAL						NOT NULL	PRIMARY KEY,
	nickname		CITEXT			UNIQUE			NOT NULL,
	email			CITEXT			UNIQUE			NOT NULL,
	password		TEXT							NOT NULL,
	points			INT,
	age				SMALLINT,
	imgurl			TEXT,
	region			TEXT,
	about			TEXT) ;`
	_, err = instance.Connection.Exec(sql)
	// Mocked users
	instance.dataCreating(100)

	return err
}

func (instance *DBHandler) CreateMessageTable() (err error) {
	sql := `
	CREATE EXTENSION IF NOT EXISTS CITEXT;
	DROP TABLE IF EXISTS messages CASCADE;

	CREATE TABLE IF NOT EXISTS messages (
	id 				BIGSERIAL						NOT NULL	PRIMARY KEY,
	nickname		CITEXT							NOT NULL,
	message			TEXT							NOT NULL,
	imgurl			TEXT,
	time			TEXT) ;`
	_, err = instance.Connection.Exec(sql)
	// Mocked users
	instance.msgCreating(100)
	return err
}

func (instance *DBHandler) InsertMessage(msg models.Message) (err error) {
	sql := `
		INSERT INTO messages (nickname, message, imgurl, time)
			VALUES ( $1, $2, $3, $4);`
	_, err = instance.Connection.Exec(sql, msg.Nickname, msg.Body, msg.Imgurl, msg.Timestamp)
	return err
}

func (instance *DBHandler) InsertUser(user models.User) (ret models.User, err error) {
	sql := `
		INSERT INTO users (nickname, email, password, points, age, imgurl, region, about)
			VALUES ( $1, $2, $3, $4, $5, $6, $7, $8)
			RETURNING nickname, email;
`
	err = instance.Connection.QueryRow(sql, user.Nickname, user.Email, user.Password, user.Points, user.Age, user.ImgUrl, user.Region, user.About).Scan(&ret.Nickname, &ret.Email)
	return ret, err
}

func (instance *DBHandler) UpdateUser(id int, user models.EditUser) (err error) {
	sql := `
		UPDATE users SET nickname = $2, email = $3, password = $4, age = $5, imgurl = $6, region = $7, about = $8 
			WHERE id = $1 RETURNING *;
`
	_, err = instance.Connection.Exec(sql, id, user.Nickname, user.Email, user.Password, user.Age, user.ImgUrl, user.Region, user.About)
	return err
}

func (instance *DBHandler) GetUser(id int) (user models.PublicUser, err error) {
	sql := `
		SELECT nickname, email, points, age, imgurl, region, about FROM users 
			WHERE id = $1;
`
	err = instance.Connection.QueryRow(sql, id).Scan(&user.Nickname, &user.Email, &user.Points, &user.Age, &user.ImgUrl, &user.Region, &user.About)
	return user, err
}

func (instance *DBHandler) GetUsers(order string, offsetdb int64, limitdb int64) (users []models.LeaderboardUsers, err error) {
	var rows *pgx.Rows
	fmt.Println(offsetdb, limitdb)
	sql := `
		SELECT nickname, email, points FROM users ORDER BY points DESC LIMIT $1 OFFSET $2;
	`
	rows, err = instance.Connection.Query(sql, limitdb, offsetdb)
	for rows.Next() {
		var u models.LeaderboardUsers
		err = rows.Scan(&u.Nickname, &u.Email, &u.Points)
		users = append(users, u)
		if err != nil {
			return nil, err
		}
	}
	return users, err
}

func (instance *DBHandler) LoginCheck(data models.LoginInfo) (user models.PublicUser, id int, err error) {
	sql := `
		SELECT id, nickname, email, points, age, imgurl, region, about FROM users 
			WHERE nickname = $1 AND password = $2;
`
	err = instance.Connection.QueryRow(sql, data.Nickname, data.Password).Scan(&id, &user.Nickname, &user.Email, &user.Points, &user.Age, &user.ImgUrl, &user.Region, &user.About)
	return user, id, err
}

func (instance *DBHandler) ImgUpdate(id int, img string) (err error) {
	sql := `
		UPDATE users SET imgurl = $2
			WHERE id = $1;
`
	_, err = instance.Connection.Exec(sql, int(id), img)
	return err
}
