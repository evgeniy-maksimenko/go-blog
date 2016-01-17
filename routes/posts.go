package routes

import (
	"../db/documents"
	"../models"
	"../session"
	"../utils"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"net/http"
)

func WriteHandler(rnd render.Render, s *session.Session) {
	if !s.IsAuth {
		rnd.Redirect("/")
	}
	model := models.EditPostModel{}
	model.IsAuth = s.IsAuth
	model.Post = models.Post{}
	rnd.HTML(200, "write", model)
}

func EditHandler(rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database, s *session.Session) {
	if !s.IsAuth {
		rnd.Redirect("/")
	}
	postsCollection := db.C("posts")
	id := params["id"]

	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}
	post := models.Post{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	model := models.EditPostModel{}
	model.IsAuth = s.IsAuth
	model.Post = post
	rnd.HTML(200, "write", model)
}

func ViewHandler(rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database, s *session.Session) {
	postsCollection := db.C("posts")
	id := params["id"]

	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}
	post := models.Post{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	model := models.ViewPostModel{}
	model.IsAuth = s.IsAuth
	model.Post = post
	rnd.HTML(200, "view", model)
}

func SavePostHandler(rnd render.Render, r *http.Request, db *mgo.Database, s *session.Session) {
	if !s.IsAuth {
		rnd.Redirect("/")
	}
	postsCollection := db.C("posts")
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("content")
	contentHtml := utils.ConvertMarkdownHtml(contentMarkdown)

	postDocuments := documents.PostDocument{id, title, contentHtml, contentMarkdown}
	if id != "" {
		postsCollection.UpdateId(id, postDocuments)
	} else {
		id = utils.GenerateId()
		postDocuments.Id = id
		postsCollection.Insert(postDocuments)
	}

	rnd.Redirect("/")
}

func DeleteHandler(rnd render.Render, r *http.Request, params martini.Params, db *mgo.Database, s *session.Session) {
	if !s.IsAuth {
		rnd.Redirect("/")
	}
	postsCollection := db.C("posts")
	id := params["id"]
	if id == "" {
		rnd.Redirect("/")
		return
	}
	postsCollection.RemoveId(id)

	rnd.Redirect("/")
}

func GetHtmlHandler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")
	html := utils.ConvertMarkdownHtml(md)

	rnd.JSON(200, map[string]interface{}{"html": html})
}
