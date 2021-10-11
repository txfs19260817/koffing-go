package koffing

type Showdown interface {
	FromJson(j string) error
	ToJson() (string, error)
	FromShowdown(s string) error
	ToShowdown() (string, error)
	Validate() error
}
