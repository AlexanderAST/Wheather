package handler

type City struct {
	City string `json:"city" binding:"required"`
}
