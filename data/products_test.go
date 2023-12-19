
package data

import "testing"

func TestCheckValidation(t *testing.T){
	p:= &Product{
		Name: "NiKlus ice cream sprinkels ",
		Price: 2.00,
	}
	err:= p.Validate()
	if err != nil{
		t.Fatal(err)
	}
}