package koffing

type showdown interface {
	FromJson(j string) error
	ToJson() (string, error)
	FromShowdown(s string) error
	ToShowdown() (string, error)
	Validate() error
}
