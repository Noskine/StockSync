package database

type (
	Migration struct {
		Up string
	}
)

var migrations = []Migration{
	{
		Up: "CREATE TABLE IF NOT EXISTS users (id uuid NOT NULL, name varchar NOT NULL );",
	},
}

func RunMigrations(migrations []Migration) error {
	for _, migration := range migrations {
		_, err := Db.Exec(migration.Up)
		if err != nil {
			return err
		}
	}

	return nil
}
