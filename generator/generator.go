package generator

import "strconv"

// The StandardGenerator takes two word lists and produces
// names made by combining one (or more) items from the first
// list (the prefix list) and one item from the second (postfix)
// list. For names like "BadBadger", "UnknowableGrue", etc.

type StandardGenerator struct {

	// word lists
	prefixes  []string
	postfixes []string

	// which prefix index to use next
	prefixes_n int

	// which postfix index to use next
	step int

	// how many prefixes do we need to generate
	// to guarantee our name is unique?
	depth int
}

func (this *StandardGenerator) Next() string {

	prefix := ""
	postfix := ""
	
	// get next prefix and postfix

	for i := 0; i < this.depth ; i += 1 {
		prefix = prefix + this.prefixes[(this.prefixes_n + i) % len(this.prefixes)]
	}

	postfix = this.postfixes[(this.prefixes_n + this.step) % len(this.postfixes)]

	// increment for next time

	this.prefixes_n = this.prefixes_n + 1

	// increase step if we looped around all the prefixes

	if this.prefixes_n == len(this.prefixes) {
		this.step += 1
		this.prefixes_n = 0
	}
	
	// increase depth if we maxed out our step (all loops done)

	if this.step == len(this.postfixes) {
		this.step = 0
		this.prefixes_n = 0
		this.depth += 1
	}
	
	return prefix+postfix
	
}

// For debugging info
func (this *StandardGenerator) State() string {
	return "prefixes_n: " + strconv.Itoa(this.prefixes_n)
}

func NewStandardGenerator(prefixes []string, postfixes []string) *StandardGenerator {

	return &StandardGenerator{
		prefixes,
		postfixes,
		0,
		0,
		1,
	}
	
}
