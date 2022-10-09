package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

const databaseFile string = "crud_app.db"

const createTable string = `
CREATE TABLE IF NOT EXISTS posts(
id INTEGER NOT NULL PRIMARY KEY,
title TEXT,
body TEXT,
user_id TEXT NOT NULL
);
`

var Posts *PostTable = NewPostTable()

type PostTable struct {
	db *sql.DB
}

func NewPostTable() *PostTable {
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err := db.Exec(createTable); err != nil {
		log.Fatalln(err)
	}
	return &PostTable{
		db: db,
	}
}

func (pt *PostTable) Create(post Post) (int, error) {
	res, err := pt.db.Exec("INSERT INTO posts VALUES (NULL,?,?,?)", post.Title, post.Body, post.UserId)
	if err != nil {
		return 0, err
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func (pt *PostTable) SelectAll() ([]Post, error) {
	res, err := pt.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for res.Next() {
		i := Post{}
		err := res.Scan(&i.Id, &i.Title, &i.Body, &i.UserId)
		if err != nil {
			return nil, err
		}
		posts = append(posts, i)
	}
	return posts, nil
}
func (pt *PostTable) SelectById(id int) (Post, error) {
	res := pt.db.QueryRow("SELECT * from posts WHERE id=?", id)
	var post Post
	err := res.Scan(&post.Id, &post.Title, &post.Body, &post.UserId)
	return post, err
}
func (pt *PostTable) Delete(id int) bool {
	_, err := pt.db.Exec("DELETE FROM posts WHERE id=?", id)
	if err != nil {
		return false
	}
	return true
}
func (pt *PostTable) Update(id int, post Post) error {
	query := "UPDATE posts SET "
	parts := make([]string, 0, 3)
	args := make([]interface{}, 0, 3)
	if post.Title != nil {
		parts = append(parts, `title = ?`)
		args = append(args, post.Title)
	}
	if post.Body != nil {
		parts = append(parts, `body = ?`)
		args = append(args, post.Body)
	}
	if post.UserId != nil {
		parts = append(parts, `user_id = ?`)
		args = append(args, post.UserId)
	}
	query += strings.Join(parts, ",") + ` WHERE id = ?`
	args = append(args, id)
	_, err := pt.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

type Post struct {
	Id     int     `json:"id,omitempty"`
	Title  *string `json:"title,omitempty"`
	Body   *string `json:"body,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

func NewPost(title string, body string, userId string) Post {
	return Post{
		Title:  &title,
		Body:   &body,
		UserId: &userId,
	}
}
