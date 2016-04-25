package generator

import (
	"testing"
)

func TestGeneratorNames(t *testing.T) {

	pre := []string{"Big", "Bad"}
	pst := []string{"Beetle", "Borgs"}
	
	generator := NewStandardGenerator(pre, pst)
	names := map[string]int{}

	// generate 1000 names and ensure there are no dupes
	
	for i := 0; i < 1000; i += 1 {

		name := generator.Next()
		
		if _, ok := names[name]; ok {
			t.Errorf("Duplicate name was given.")
		} else {
			names[name] = 0;
		}
		
	}
	
}
