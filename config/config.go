package config

type Config struct {
	DBConfig *DBConnectionConfig
	// DevDBConfig which use for Database migrator to replay existing generated migrations, and perform a diff
	// between ent schema, and comes up a new migration file.
	DevDBConfig *DBConnectionConfig
	RedisConfig *RedisConfig
	Server      *Server
}
