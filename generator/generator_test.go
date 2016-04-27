package generator

import (
	"testing"
)

func TestGeneratorNames(t *testing.T) {

	pre := []string{"Big", "Bad"}
	pst := []string{"Beetle", "Borgs"}
	
	generator := NewStandardGenerator(pre, pst, ".state")
	names := map[string]int{}

	for i := 0; i < 1000; i += 1 {

		name := generator.Next()
		
		if _, ok := names[name]; ok {
			t.Errorf("Duplicate name was given.")
		} else {
			names[name] = 0;
		}
		
	}
	
}

func TestGeneratorPersistence(t *testing.T) {

	pre := []string{"Big", "Bad"}
	pst := []string{"Beetle", "Borgs"}
	
	generator := NewStandardGenerator(pre, pst, ".state")

	out := generator.Save("ignore")

	if out != "{\"Prefixes_n\":0,\"Step\":0,\"Depth\":1}" {
		t.Errorf("Bad json")
	}

	generator.Next()

	if generator.State().Prefixes_n != 1 {
		t.Errorf("Generator state did not change")
	}
	
	generator.Load("ignore")

	if generator.State().Prefixes_n != 0 {
		t.Errorf(
			"Generator load did not restore values, had %d",
			generator.State().Prefixes_n,
		)
	}

}
