package cmd

type RmbrNotesManager struct {
	Repository Repository
	Logger     *StandardLogger
}

func NewRmbrNotesManager(repository Repository, logger *StandardLogger) *RmbrNotesManager {
	return &RmbrNotesManager{
		Repository: repository,
		Logger:     logger,
	}
}

type RmbrNote struct {
	Command     string `yaml:"command"`
	Description string `yaml:"summary"`
	Group       string `yaml:"group"`
}

type RmbrNotes struct {
	Notes []RmbrNote `yaml:"notes"`
}
