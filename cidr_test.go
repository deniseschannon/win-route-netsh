package winroutenetsh

import (
	"io/ioutil"
	"testing"
)

func Test_SingleRow(t *testing.T) {
	data, err := ioutil.ReadFile("./test/testSingleRow")
	if err != nil {
		t.Fatal(err)
	}
	row, err := parseRouteRow(data)
	if err != nil {
		t.Fatal(err)
	}
	println(row.ToPowershellString())
	newS, err := row.ToNewString()
	if err != nil {
		t.Fatal(err)
	}
	println(newS)
}

func Test_Rows(t *testing.T) {
	data, err := ioutil.ReadFile("./test/testRows")
	rows, err := parseRows(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 {
		t.Fail()
	}
}
