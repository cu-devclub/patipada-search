package handlers

import (
	"data-management/errors"
	"data-management/messages"
	"data-management/ratings/models"
	"data-management/ratings/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ratingHandler struct {
	ratingsUsecase usecases.Usecase
}

func NewRatingHandler(ratingsUsecase usecases.Usecase) Handlers {
	return &ratingHandler{
		ratingsUsecase: ratingsUsecase,
	}
}

func (r *ratingHandler) InsertRating(c *gin.Context) {
	// binding models
	var rating models.Rating

	handlerOpts := NewHandlerOpts(c)
	handlerOpts.Params = rating

	if err := c.ShouldBindJSON(&rating); err != nil {
		r.errorResponse(c, handlerOpts, http.StatusBadRequest, messages.BAD_REQUEST)
		return
	}

	handlerOpts.Params = rating

	// call usecase
	id, err := r.ratingsUsecase.InsertRating(&rating)
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			r.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
			return
		} else {
			r.errorResponse(c, handlerOpts, 500, messages.INTERNAL_SERVER_ERROR)
			return
		}
	}

	// return response
	resp := ResponseOptions{
		Response: id,
	}

	r.successResponse(c, *handlerOpts, http.StatusCreated, resp)
}

func (h *ratingHandler) GetRatings(c *gin.Context) {
	handlerOpts := NewHandlerOpts(c)

	// call usecase
	ratings, err := h.ratingsUsecase.GetRatings()
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			h.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
			return
		} else {
			h.errorResponse(c, handlerOpts, 500, messages.INTERNAL_SERVER_ERROR)
			return
		}
	}

	// return response
	resp := ResponseOptions{
		Response: ratings,
		OptionalResponse: &RatingsLog{
			Amount: len(ratings),
		},
	}

	h.successResponse(c, *handlerOpts, 200, resp)
}

func (h *ratingHandler) GetAverageRatings(c *gin.Context) {
	handlerOpts := NewHandlerOpts(c)

	// call usecase
	average, err := h.ratingsUsecase.GetAverageRatings()
	if err != nil {
		if er, ok := err.(*errors.RequestError); ok {
			h.errorResponse(c, handlerOpts, er.StatusCode, er.Error())
			return
		} else {
			h.errorResponse(c, handlerOpts, 500, messages.INTERNAL_SERVER_ERROR)
			return
		}
	}

	// return response
	resp := ResponseOptions{
		Response: average,
	}

	h.successResponse(c, *handlerOpts, 200, resp)

}
