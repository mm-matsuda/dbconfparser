package dbconfparser

import (
	"testing"
)

func TestParse (t *testing.T) {
	driver, open, err := Parse("test/dbconf.yml", "test")
	expected1 := "dbms2"
	if driver != expected1 {
		t.Errorf("got: %v\nwant: %v", driver, expected1)
	}

	expected2 := "testuser:password@tcp(127.0.0.1:13332)/database_test?parseTime=true"
	if open != expected2 {
		t.Errorf("got: %v\nwant: %v", open, expected2)
	}

	if err != nil {
		t.Errorf("Error is occured: %v", err)
	}
}

func TestReadFile (t *testing.T) {
	_, _, err := Parse("hogehoge", "hoge")
	expected := "open hogehoge: no such file or directory"
	if err.Error() != expected {
		t.Errorf("got: %v\nwant: %v", err, expected)
	}
}

func TestYamlUnmarshal (t *testing.T) {
	_, _, err := Parse("test/dbconferror.yml", "hoge")
	expected := "yaml: line 1: did not find expected node content"
	if err.Error() != expected {
		t.Errorf("got: %v\nwant: %v", err, expected)
	}
}
