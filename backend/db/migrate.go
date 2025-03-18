package db

// Perform migrations on a supplied models.
// Uses Database as the database connection.
func Migrate(model interface{}) error {
	if err := Database.AutoMigrate(&model); err != nil {
		return err
	}
	return nil
}
