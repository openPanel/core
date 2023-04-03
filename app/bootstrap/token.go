package bootstrap

import (
	"github.com/google/uuid"

	"github.com/openPanel/core/app/config"
	"github.com/openPanel/core/app/constant"
)

func createToken() error {
	token := uuid.NewString()
	return config.Save(constant.ConfigKeyAuthorizationToken, token, constant.SharedStore)
}
