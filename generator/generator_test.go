package generator

import "testing"

func TestGenerator(t *testing.T) {

	pre := []string{"Big", "Bad"}
	pst := []string{"Beetle", "Borgs"}
	
	a := NewStandardGenerator(&pre, &pst)

	if a.Next() != "BigBeetle" {
		//t.Error("First term was not BigBeetle")
	}

}
