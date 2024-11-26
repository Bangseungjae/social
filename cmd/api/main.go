package main

import (
	db2 "github.com/Bangseungjae/social/internal/db"
	"github.com/Bangseungjae/social/internal/env"
	mailer2 "github.com/Bangseungjae/social/internal/mailer"
	internalstore "github.com/Bangseungjae/social/internal/store"
	"go.uber.org/zap"
	"time"
)

const version = "0.0.1"

//	@title			GopherSocial API
//	@description	API for GopherSocial, a social network for gopher
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	cfg := config{
		addr:        env.GetString("ADDR", ":8080"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:5173"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp: time.Hour * 24 * 3, // 3 days
			sendGrid: sendGridConfig{
				fromEmail: env.GetString("FROM_EMAIL", ""),
				apiKey:    env.GetString("SENDGRID_API_KEY", ""),
			},
		},
		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_PASS", "admin"),
			},
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db2.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	defer db.Close()
	logger.Info("database connection pool establish")

	if err != nil {
		logger.Panic(err)
	}

	store := internalstore.NewStorage(db)

	mailer := mailer2.NewSendgrid(cfg.mail.sendGrid.apiKey, cfg.mail.sendGrid.fromEmail)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
		client: mailer,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
