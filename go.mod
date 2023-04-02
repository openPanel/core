module github.com/openPanel/core

go 1.20

require (
	entgo.io/contrib v0.3.5
	entgo.io/ent v0.11.10
	github.com/canonical/go-dqlite v1.11.7
	github.com/flowchartsman/swaggerui v0.0.0-20221017034628-909ed4f3701b
	github.com/go-co-op/gocron v1.19.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zap/v2 v2.0.0-rc.3
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/julienschmidt/httprouter v1.3.0
	github.com/lorenzosaino/go-sysctl v0.3.1
	github.com/mattn/go-sqlite3 v1.14.16
	github.com/pkg/errors v0.9.1
	github.com/quic-go/quic-go v0.33.0
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.2
	go.uber.org/zap v1.24.0
	google.golang.org/genproto v0.0.0-20230331144136-dcfb400f0633
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

replace (
	entgo.io/ent v0.11.10 => github.com/openPanel/ent v0.0.0-20230402135651-305c3704e8e2
	github.com/canonical/go-dqlite v1.11.7 => github.com/openPanel/go-dqlite v0.0.0-20230331113006-e579e9de3c20
	github.com/flowchartsman/swaggerui v0.0.0-20221017034628-909ed4f3701b => github.com/openPanel/swaggerui v0.0.0-20230401141121-264dd475eced
)

require (
	ariga.io/atlas v0.10.0 // indirect
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/Rican7/retry v0.3.1 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/bufbuild/protocompile v0.5.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20230323073829-e72429f035bd // indirect
	github.com/google/renameio v1.0.1 // indirect
	github.com/hashicorp/hcl/v2 v2.16.2 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jhump/protoreflect v1.15.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/onsi/ginkgo/v2 v2.9.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/qtls-go1-19 v0.3.2 // indirect
	github.com/quic-go/qtls-go1-20 v0.2.2 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/zclconf/go-cty v1.13.1 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29 // indirect
	golang.org/x/exp/typeparams v0.0.0-20230321023759-10a507213a29 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mod v0.9.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	honnef.co/go/tools v0.4.3 // indirect
)
