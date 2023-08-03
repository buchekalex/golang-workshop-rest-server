package v1

import (
	"errors"
	"go-mongo-crud-rest-api/internal/entitiy"
	"go-mongo-crud-rest-api/internal/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	user usecase.User
}

func NewServer(user usecase.User) *Server {
	return &Server{
		user: user,
	}
}

func (s Server) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}

func (s Server) GetUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument email"})
		return
	}

	user, err := s.user.GetUser(ctx, email)
	if err != nil {
		log.Printf("mongo GetUser: error: %s", err.Error())
		if errors.Is(err, usecase.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) CreateUser(ctx *gin.Context) {
	var user entitiy.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if _, err := s.user.CreateUser(ctx, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) UpdateUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument email"})
		return
	}
	var user entitiy.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	user.Email = email

	if _, err := s.user.UpdateUser(ctx, &user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (s Server) DeleteUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument email"})
		return
	}

	if err := s.user.DeleteUser(ctx, email); err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (s Server) FindUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument email"})
		return
	}

	user, err := s.user.GetUser(ctx, email)
	if err != nil {
		log.Printf("mongo GetUser: error: %s", err.Error())
		if errors.Is(err, usecase.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
