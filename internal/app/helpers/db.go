package helpers

//// InitPostgres is used to initiate to db postgres connection
//func (h *Helpers) InitPostgres(cfg *config.Config) *sql.DB {
//
//	var dbCfg = cfg.DatabaseCfg[0]
//	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbCfg.DatabaseCfg., cfg.DBPass, cfg.Host, cfg.Port, cfg.DBName, cfg.SslMode)
//
//	conn, err := sql.Open(`postgres`, connStr)
//	if err != nil {
//		log.Fatal(err)
//		panic(err)
//	}
//	err = conn.Ping()
//
//
//	if err != nil {
//		log.Fatal(err)
//		os.Exit(1)
//	}
//
//	maxOpenConnes, _ := strconv.Atoi(cfg.MaxOpenConns)
//	maxIdleConnes, _ := strconv.Atoi(cfg.MaxIdleConns)
//	connMaxLifetime, _ := strconv.Atoi(cfg.ConnMaxLifetime)
//
//	conn.SetMaxOpenConns(maxOpenConnes)
//	conn.SetMaxIdleConns(maxIdleConnes)
//	conn.SetConnMaxLifetime(time.Duration(connMaxLifetime) *time.Second)
//	h.dbConn = conn
//	return conn
//}
//
//// BeginTrx is helper to begin the database transaction
//func (h *Helpers) BeginTrx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
//	trx, err := h.dbConn.BeginTx(ctx, opts)
//	return trx, err
//}