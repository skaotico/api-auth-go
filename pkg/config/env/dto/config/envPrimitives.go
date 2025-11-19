package config

import "time"

type Config struct {
	// AppPort es el puerto donde correrá el servidor HTTP.
	// Ejemplo: ":8080".
	AppPort string `envconfig:"APP_PORT" required:"true"`

	// Environment define el modo de ejecución de la aplicación.
	// Ejemplos: "development", "production", "staging".
	Environment string `envconfig:"ENV" required:"true"`

	// JWTSecret define el secreto utilizado para firmar y validar
	// tokens JWT. Es obligatorio por seguridad.
	JWTSecret string `envconfig:"JWT_SECRET" required:"true"`

	// DBHost es la dirección del host de la base de datos.
	DBHost string `envconfig:"DB_HOST" required:"true"`

	// DBPort es el puerto de conexión a la base de datos.
	// Valor por defecto: "5432".
	DBPort string `envconfig:"DB_PORT" default:"5432"`

	// DBUser es el usuario de base de datos utilizado para
	// autenticación.
	DBUser string `envconfig:"DB_USER" required:"true"`

	// DBPass es la contraseña del usuario de base de datos.
	DBPass string `envconfig:"DB_PASS" required:"true"`

	// DBName es el nombre de la base de datos a utilizar.
	DBName string `envconfig:"DB_NAME" required:"true"`

	// ReadTimeout es el tiempo máximo permitido para leer una
	// solicitud HTTP.
	ReadTimeout time.Duration `envconfig:"READ_TIMEOUT" default:"5s"`

	// WriteTimeout es el tiempo máximo permitido para enviar una
	// respuesta HTTP.
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"10s"`

	// JWTExpiration define el tiempo de expiración del token JWT (acceso).
	// Ejemplo: "15m", "1h".
	JWTExpiration time.Duration `envconfig:"JWT_EXPIRATION" default:"15m"`

	// JWTRefreshTTL define el tiempo de vida del refresh token.
	// Ejemplo: "24h", "7d".
	JWTRefreshTTL time.Duration `envconfig:"JWT_REFRESH_TTL" default:"24h"`

	// Version define la versión actual de la aplicación.
	Version string `envconfig:"VERSION" required:"true"`
}
