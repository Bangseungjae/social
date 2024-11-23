package main

import (
	"Bangseungjae/social/internal/store"
	"fmt"
	"net/http"
)

type UpdateCommentPayload struct {
	Content *string `json:"content" validate:"omitempty,max=500"`
}

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	var payload UpdateCommentPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Content == nil {
		app.badRequestResponse(w, r, fmt.Errorf("comment not be nil"))
		return
	}

	ctx := r.Context()

	comment := store.Comment{
		PostID:  post.ID,
		UserID:  1,
		Content: *payload.Content,
	}

	if err := app.store.Comments.Create(ctx, &comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
