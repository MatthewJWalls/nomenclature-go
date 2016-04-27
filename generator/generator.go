package generator

import (
	"encoding/json"
	"io/ioutil"
)

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

	// the path to the file where we're going to
	// persist our state to, so we carry off from
	// the same place next time
	stateFile string
}

// This is used for persisting the state of the generator to
// disk via json marshalling.

type StandardGeneratorRepr struct {
	Prefixes_n int
	Step int
	Depth int
}

// Next() implements the core algorithm for deciding the next
// name to be given by the generator.

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

	this.Save(this.stateFile)	
	
	return prefix+postfix
	
}

// restore generator state from disk

func (this *StandardGenerator) Load(path string) string {

	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}
	
	repr := StandardGeneratorRepr{}

	err = json.Unmarshal(bytes, &repr)

	if err != nil {
		panic("Failed to parse generator state")
	}

	this.prefixes_n = repr.Prefixes_n
	this.step = repr.Step
	this.depth = repr.Depth

	return string(bytes)

}

// persist generator state to disk

func (this *StandardGenerator) Save(path string) string {

	repr := StandardGeneratorRepr{
		this.prefixes_n,
		this.step,
		this.depth,
	}

	bytes, err := json.Marshal(repr)

	if err != nil {
		panic("Failed to persist generator from disk")
	}

	ioutil.WriteFile(path, bytes, 0644)

	return string(bytes)
		
}

// For debugging info

func (this *StandardGenerator) State() StandardGeneratorRepr {

	repr := StandardGeneratorRepr{
		this.prefixes_n,
		this.step,
		this.depth,
	}

	return repr

}

func NewStandardGenerator(prefixes []string, postfixes []string, save string) *StandardGenerator {

	return &StandardGenerator{
		prefixes,
		postfixes,
		0,
		0,
		1,
		save,
	}
	
}
