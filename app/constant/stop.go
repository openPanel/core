package constant

type StopID string

const (
	StopIDLocalSqliteDB  StopID = "local_sqlite_db"
	StopIDSharedDqliteDB StopID = "shared_dqlite_db"

	StopIDLogger        StopID = "log_file"
	StopIDCron          StopID = "cron"
	StopIDQUICConnCache StopID = "quic_conn_cache"

	StopIDGRPCUnixServer StopID = "grpc_unix_server"
	StopIDGRPCQUICServer StopID = "grpc_quic_server"
	StopIDGRPCHTTPServer StopID = "grpc_web_server"
)
