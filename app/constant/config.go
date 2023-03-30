package constant

type Mode string

const (
	ModeDev  Mode = "dev"
	ModeProd Mode = "prod"
)

type Store int

const (
	LocalStore  Store = iota
	SharedStore Store = iota
)

type Key string

const (
	ConfigKeyNodeInfo Key = "local.nodeInfo"
)

const ProtoMixinIdStart = 536870900

const (
	ConfigKeyClusterInfo        Key = "shared.clusterInfo"
	ConfigKeyAuthorizationToken Key = "shared.authorizationToken"
)
