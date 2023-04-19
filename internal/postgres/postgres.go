package postgres

// type Database struct {
// 	Conn *pg.DB
// 	Tx   *pg.Tx
// }

// type debugHook struct{}

// // BeforeQuery ...
// func (debugHook) BeforeQuery(ctx context.Context, evt *pg.QueryEvent) (context.Context, error) {
// 	q, err := evt.FormattedQuery()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if evt.Err != nil {
// 		log.Printf("Error %s executing query:\n%s\n", evt.Err, q)
// 	} else {
// 		log.Printf("%s", q)
// 	}

// 	return ctx, nil
// }

// // AfterQuery
// func (debugHook) AfterQuery(context.Context, *pg.QueryEvent) error {
// 	return nil
// }

// var _ pg.QueryHook = (*debugHook)(nil)

// func NewDatabase() *Database {
// 	db := &Database{}
// 	db.Connect()

// 	return db
// }

// // Connect ...
// func (db *Database) Connect() {
// 	var tlsConfig *tls.Config

// 	if env.Env() == "production" || env.Env() == "staging" || env.Env() == "development" {
// 		tlsConfig = &tls.Config{InsecureSkipVerify: true}
// 	}

// 	db.Conn = pg.Connect(&pg.Options{
// 		ApplicationName: env.AppName(),
// 		User:            env.DBUser(),
// 		Password:        env.DBPassword(),
// 		Addr:            env.DBAddress(),
// 		Database:        env.DBDatabase(),
// 		TLSConfig:       tlsConfig,
// 	})

// 	if env.Env() == "local" || env.Env() == "development" {
// 		db.Conn.AddQueryHook(debugHook{})
// 	}
// }

// // Close ...
// func (db *Database) Close() {
// 	db.Conn.Close()
// }

// func (db *Database) Health() (bool, error) {
// 	if db.Conn != nil {
// 		return true, nil
// 	}

// 	return false, fmt.Errorf("internal-postgres: connection unavailable")
// }

// func (db *Database) Get() (*Database, error) {
// 	if db.Conn != nil {
// 		return db, nil
// 	}

// 	return nil, fmt.Errorf("internal-postgres: connection unavailable")
// }

// func (db *Database) NewTrx() error {
// 	tx, err := db.Conn.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	db.Tx = tx

// 	return nil
// }

// func (db *Database) ErrNoRows() error {
// 	return pg.ErrNoRows
// }
