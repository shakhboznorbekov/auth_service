package handler

import (
	"github.com/auth_service/config"
	"github.com/auth_service/storage"
)

type HandlerV1 struct {
	cfg     *config.Config
	storage storage.StorageI
}

func NewHandlerV1(cfg *config.Config, storage storage.StorageI) *HandlerV1 {
	return &HandlerV1{
		cfg:     cfg,
		storage: storage,
	}
}
