package service

import (
	"ContentManager/api/operate"
	"ContentManager/internal/biz"
)

type AppService struct {
	operate.UnimplementedAppServer

	uc *biz.GreeterUsecase
}

func NewAppSerivce(uc *biz.GreeterUsecase) *AppService {
	return &AppService{uc: uc}
}
