package cmd

type FakeRepository struct {
	Note *RmbrNotes
	HasBeenCalled bool
}

func (m FakeRepository) New(note RmbrNote) error {
	return nil
}

func (m FakeRepository) Load() (*RmbrNotes, error) {
	return m.Note, nil
}

func (m FakeRepository) Write(applications *RmbrNotes) error {
	return nil
}

func NewFakeRepository(notes *RmbrNotes) (Repository, error) {
	loader := &FakeRepository{}
	loader.Note = notes
	return loader, nil
}
