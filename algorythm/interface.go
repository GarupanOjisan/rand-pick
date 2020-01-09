package algorythm

type Algorithm interface {
	RandomPick(*map[int]interface{}) (interface{}, error)
	Validate(*map[int]interface{}) error
}
