package koffing

type Team struct {
	Name    string    `json:"name,omitempty"`
	Format  string    `json:"format,omitempty"`
	Folder  string    `json:"folder,omitempty"`
	Pokemon []Pokemon `json:"pokemon,omitempty"`
}
