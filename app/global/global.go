package global

import (
	"github.com/openPanel/core/app/generated/db/local"
	"github.com/openPanel/core/app/generated/db/shared"
	"github.com/openPanel/core/app/global/buildType"
)

type application struct {
	DbLocal  *local.Client
	DbShared *shared.Client

	NodeInfo LocalNodeInfo

	Mode buildType.Mode
}

var App = application{}

func init() {
	App.Mode = buildType.BuildMode
}

func IsDev() bool {
	return App.Mode == buildType.ModeDev
}
