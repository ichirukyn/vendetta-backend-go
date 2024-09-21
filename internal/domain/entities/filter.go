package entities

type WhereQueryFilter = map[string]interface{}

type Filter struct {
	Limit  int
	Offset int
	Order  string
}
