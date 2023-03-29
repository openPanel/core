package constant

type Store int

const (
	LocalStore  Store = iota
	SharedStore Store = iota
)

type Key string

const (
	NodeInfoConfigKey    Key = "local.nodeInfo"
	ClusterInfoConfigKey Key = "shared.clusterInfo"
)

const ProtoMixinIdStart = 536870900

type Mode string

const (
	ModeDev  Mode = "dev"
	ModeProd Mode = "prod"
)
