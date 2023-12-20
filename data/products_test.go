
package data

import "testing"

func TestCheckValidation(t *testing.T){
	p:= &Product{
		Name: "NiKlus ice cream & sprinkels ",
		Price: 2.00,
		SKU: "shi-nob-uye",//three groups of 3 letters lowercase just like in regex
	}
	err:= p.Validate()
	if err != nil{
		t.Fatal(err)
	}
}