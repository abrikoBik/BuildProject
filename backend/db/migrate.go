package db

// Perform migrations on all supplied models.
// Uses Database as the database connection.
func Migrate(models ...interface{}) error {
	for _, model := range models {
		if err := Database.AutoMigrate(&model); err != nil {
			return err
		}
	}
	return nil
}
