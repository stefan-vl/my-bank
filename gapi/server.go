package gapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/stefan-vl/my-bank/db/sqlc"
	"github.com/stefan-vl/my-bank/pb"
	"github.com/stefan-vl/my-bank/token"
	"github.com/stefan-vl/my-bank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &Server{store: store, config: config, tokenMaker: tokenMaker}

	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
