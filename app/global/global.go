package global

import (
	"go.uber.org/zap"

	"github.com/openPanel/core/app/generated/db/local"
	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/global/buildType"
)

type application struct {
	DbL *local.Client
	DbS *shared.Client

	Logger *zap.SugaredLogger

	Mode buildType.Mode
}

var App = application{}

func init() {
	App.Mode = buildType.BuildMode
}

func IsDev() bool {
	return App.Mode == buildType.ModeDev
}
