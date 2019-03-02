package types

type Project struct {
	ID        uint
	ProjectID uint // Can have nested projects; only one parent
	Name      string
}

type Show struct {
	ID        uint
	Name      string
	ProjectID uint
}
