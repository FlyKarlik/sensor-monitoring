package migrator

func (a *App) Up() error {
	if err := a.migrator.Up(); err != nil {
		a.logger.Error("App", "App.Start", "Failed to migrate database", err)
		return err
	}

	a.logger.Info("App", "App.Start", "Database migrated successfully")

	return nil
}
