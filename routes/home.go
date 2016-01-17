package routes

import (
	"../db/documents"
	"../models"
	"../session"
	"fmt"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func IndexHandler(rnd render.Render, s *session.Session, db *mgo.Database) {
	fmt.Println(s.Username)

	postDocuments := []documents.PostDocument{}
	postsCollection := db.C("posts")
	postsCollection.Find(nil).All(&postDocuments)

	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{doc.Id, doc.Title, doc.ContentHtml, doc.ContentMarkdown}
		posts = append(posts, post)
	}
	rnd.HTML(200, "index", posts)
}
