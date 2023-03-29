package global

import (
	"go.uber.org/zap"

	"github.com/openPanel/core/app/constant"
	"github.com/openPanel/core/app/generated/db/local"
	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/global/buildType"
)

type application struct {
	DbLocal  *local.Client
	DbShared *shared.Client

	NodeInfo    LocalNodeInfo
	ClusterInfo ClusterInfo

	Mode constant.Mode

	Logger *zap.Logger
}

var App = application{}

func init() {
	App.Mode = buildType.BuildMode
}

func IsDev() bool {
	return App.Mode == constant.ModeDev
}
