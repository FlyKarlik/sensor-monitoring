package migrator

func (a *App) Down() error {
	if err := a.migrator.Down(); err != nil {
		a.logger.Error("App", "App.Start", "Failed to migrate database", err)
		return err
	}

	a.logger.Info("App", "App.Start", "Database migrated successfully")

	return nil
}
