package database

type (
	Migration struct {
		Up string
	}
)

func RunMigrations(migrations []Migration) error {
	conn := OpenConn()
	defer conn.Close()

	for _, migration := range migrations {
		_, err := conn.Exec(migration.Up)
		if err != nil {
			return err
		}
	}

	return nil
}
