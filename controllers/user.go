package controllers

import (
	database "cravefeed_backend/database"
	helpers "cravefeed_backend/helper"
	"cravefeed_backend/interfaces"
	"cravefeed_backend/prisma/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	pClient := database.PClient
	allUsers, err := pClient.Client.User.FindMany().Exec(pClient.Context)
	if err != nil {
		fmt.Println("Cannot fetch users")
		return

	}
	usersMap := make(map[string]interface{})
	usersMap["users"] = allUsers
	err = helpers.WriteJSON(w, http.StatusOK, usersMap)
	if err != nil {
		fmt.Println("Cannot form response")
		return
	}

}

func GetName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var nameReq interfaces.NameRequest
	err := json.NewDecoder(r.Body).Decode(&nameReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, map[string]string{
		"Name": nameReq.Name,
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) { //Removed optional field in Avatar/Bio for testing
	defer r.Body.Close()
	pClient := database.PClient
	var userData interfaces.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&userData)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdUser, err := pClient.Client.User.CreateOne(
		db.User.Email.Set(userData.Email),
		db.User.Password.Set(userData.Password),
		db.User.Bio.Set(userData.Bio),
		db.User.Avatar.Set(userData.Avatar),
		db.User.FirstName.Set(userData.FirstName),
		db.User.LastName.Set(userData.LastName),
	).Exec(pClient.Context)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	helpers.WriteJSON(w, http.StatusOK, createdUser)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	pClient := database.PClient
	var postData interfaces.CreatePostRequest
	err := json.NewDecoder(r.Body).Decode(&postData)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdPost, err := pClient.Client.Post.CreateOne(
		db.Post.Title.Set(postData.Title),
		db.Post.Description.Set(postData.Description),
		db.Post.Longitude.Set(postData.Longitude),
		db.Post.Latitude.Set(postData.Latitude),
		db.Post.Pictures.Set(postData.Pictures),
		db.Post.UserID.Set(postData.UserID),
		db.Post.City.Set(postData.City),
	).Exec(pClient.Context)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	helpers.WriteJSON(w, http.StatusOK, createdPost)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    pClient := database.PClient
    var commentData interfaces.CreateCommentRequest
    err := json.NewDecoder(r.Body).Decode(&commentData)

    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    createdComment, err := pClient.Client.Comment.CreateOne(
        db.Comment.Content.Set(commentData.Content),
        db.Comment.PostID.Set(commentData.PostID),
        db.Comment.UserID.Set(commentData.UserID),
    ).Exec(pClient.Context)

    if err != nil {
        http.Error(w, "Failed to create comment", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    helpers.WriteJSON(w, http.StatusOK, createdComment)
}

func HandleLikeRequest(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    pClient := database.PClient
    var likeReq interfaces.LikeRequest
    err := json.NewDecoder(r.Body).Decode(&likeReq)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    createdLike, err := pClient.Client.Like.CreateOne(
        db.Like.PostID.Set(likeReq.PostID),
        db.Like.UserID.Set(likeReq.UserID),
    ).Exec(pClient.Context)

    if err != nil {
        http.Error(w, "Failed to create like", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    helpers.WriteJSON(w, http.StatusOK, createdLike)
}

func HandleFollowRequest(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    pClient := database.PClient
    var followReq interfaces.FollowRequest
    err := json.NewDecoder(r.Body).Decode(&followReq)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    createdFollow, err := pClient.Client.Follow.CreateOne(
        db.Follow.FollowerID.Set(followReq.FollowerID),
        db.Follow.FollowingID.Set(followReq.FollowingID),
    ).Exec(pClient.Context)

    if err != nil {
        http.Error(w, "Failed to create follow", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    helpers.WriteJSON(w, http.StatusOK, createdFollow)
}
