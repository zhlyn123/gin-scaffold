package database

func (m *Mysql) Health() error {
	db, err := m.DB.DB()
	if err != nil {
		return err
	}

	return db.Ping()
}