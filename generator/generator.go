package generator

// The StandardGenerator takes two word lists and produces
// names made by combining one (or more) items from the first
// list (the prefix list) and one item from the second (postfix)
// list. For names like "BadBadger", "UnknowableGrue", etc.

type StandardGenerator struct {
	prefixes  *[]string
	postfixes *[]string
}

func (this StandardGenerator) Next() string {
	return "test"
}

func NewStandardGenerator(prefixes *[]string, postfixes *[]string) StandardGenerator {

	f := StandardGenerator{ prefixes, postfixes }
	return f
	
}
