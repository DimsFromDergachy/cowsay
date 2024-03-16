package cowsay

func Say (phrase string) (string, error) {
	cow := &Cow{}
	return cow.Say(phrase)
}