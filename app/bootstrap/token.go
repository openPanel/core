package bootstrap

import (
	"github.com/google/uuid"

	"github.com/openPanel/core/app/config"
	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/global/log"
)

func createToken() error {
	token := uuid.NewString()
	log.Info("Generated new token: " + token)
	return config.Save(constant.ConfigKeyAuthorizationToken, token, constant.SharedStore)
}
