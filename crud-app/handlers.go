package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func PostController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPosts(w, r)
		break
	case http.MethodPost:
		create(w, r)
		break
	case http.MethodDelete:
		deletePost(w, r)
		break
	case http.MethodPut:
		update(w, r)
	case http.MethodPatch:
		update(w, r)
	default:
		http.Error(w, "Invalid request method.", 405)
		break
	}
}

func parseId(r *http.Request) (int, error) {
	idFromPath := strings.TrimPrefix(r.URL.Path, "/posts/")

	id, err := strconv.Atoi(idFromPath)
	if err != nil {
		return 0, http.ErrNoLocation
	}
	return id, nil
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/posts")
	idFromPath := strings.TrimPrefix(path, "/")
	if idFromPath == "" {
		// get all posts
		allPosts, err := Posts.SelectAll()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error getting posts", 500)
		}
		err = json.NewEncoder(w).Encode(allPosts)
		if err != nil {
			http.Error(w, "Error getting posts", 500)
		}
	} else {
		// get one post
		id, err := strconv.Atoi(idFromPath)
		if err != nil {
			http.Error(w, "Invalid request", 400)
			return
		}
		post, err := Posts.SelectById(id)
		if err != nil {
			http.Error(w, "Error getting posts", 500)
			return
		}
		err = json.NewEncoder(w).Encode(post)
		if err != nil {
			http.Error(w, "Error getting posts", 500)
			return
		}
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	var post Post
	id, err := parseId(r)
	if err != nil {
		http.Error(w, "Error updating post", 500)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error updating posts", 500)
		return
	}
	if err := Posts.Update(id, post); err != nil {
		log.Println(err)
		http.Error(w, "Error getting posts", 500)
		return
	}
	postById, err := Posts.SelectById(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error getting posts", 500)
		return
	}
	err = json.NewEncoder(w).Encode(postById)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error getting posts", 500)
	}

}

func create(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatalln(err)
	}
	id, err := Posts.Create(post)
	if err != nil {
		log.Fatalln(err)
		return
	}
	createdPost, err := Posts.SelectById(id)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = json.NewEncoder(w).Encode(createdPost)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		http.Error(w, "Invalid request", 405)
		return
	}
	hasDeleted := Posts.Delete(id)
	if hasDeleted {
		return
	} else {
		http.Error(w, "Failed to delete post", 500)
	}

}
