package constant

const CORS_ALLOWED_ORIGINS = "CORS_ALLOWED_ORIGINS"

// Application environment variable keys
const APP_DEBUG = "APP_DEBUG"
const APP_ENV = "APP_ENV"
const APP_PORT = "APP_PORT"
const APP_HOST = "APP_HOST"
const APP_URL = "APP_URL"

// Database environment variable keys
const DB_HOST = "DB_HOST"
const DB_USER = "DB_USER"
const DB_NAME = "DB_NAME"
const DB_PASSWORD = "DB_PASSWORD"
const DB_PORT = "DB_PORT"

// Dev Database environment variable keys
const DEV_DB_HOST = "DEV_DB_HOST"
const DEV_DB_USER = "DEV_DB_USER"
const DEV_DB_NAME = "DEV_DB_NAME"
const DEV_DB_PASSWORD = "DEV_DB_PASSWORD"
const DEV_DB_PORT = "DEV_DB_PORT"

// context keys
type WalletAppContextgKey string

const (
	ApplicationContextKey WalletAppContextgKey = "applicationCtx"
)
