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

type ConfigKey string

const (
	ConfigKeyNodeInfo   ConfigKey = "local.nodeInfo"
	ConfigKeyNodesCache ConfigKey = "local.nodesCache"
)

const ProtoMixinIdStart = 536870900

const (
	ConfigKeyClusterInfo        ConfigKey = "shared.clusterInfo"
	ConfigKeyAuthorizationToken ConfigKey = "shared.authorizationToken"
)
