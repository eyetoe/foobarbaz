package choices

type Choice struct {
	name   string
	letter string
	action func()
}
