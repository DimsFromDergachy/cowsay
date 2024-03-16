package cowsay

type Cow struct {

}

func (cow *Cow) Say(phrase string) (string, error) {
	return "There is a cowsay: '" + phrase + "'", nil
}