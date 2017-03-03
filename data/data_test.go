package data

import (
    "testing"
	"time"
	"log"
)

func TestFolderCheck(t *testing.T){
	f:= Folder{Id: 1, ParentId: 0, Name: "", LastModified: time.Now()}
	
	_, err := f.Check()
	if err == nil{
		t.Error("This folder has an empty name.")
	}
}

var adNameTests = []struct {
    in  string
    out bool
}{
    {"", false},
    {"Okname", true},
    {"Badword132456", false},
}

var adUrlTests = []struct {
    in  string
    out bool
}{
    {"", false},
    {"http://celtra.com", true},
    {"htp:/google.si", false},
	{"http://_g#$%oogle.com", false},
}


func TestAdCheck(t *testing.T){
	a:= Ad{Id: 1, FolderId: 0, Name: "Name", Url: "http://www.validurl.com", LastModified: time.Now()}
	//test a range of different names
	for _, tt := range adNameTests {
		a.Name = tt.in;		
		isOk, err := a.Check()
		
		if isOk != tt.out{
			log.Println(err)
			t.Error("Ad NAME test failed for input: " + tt.in)
		}
	}
	
	a.Name = "Name"
	//test a range of different URL inputs
	for _, tt := range adUrlTests {
		a.Url = tt.in;		
		isOk, err := a.Check()
		
		if isOk != tt.out{
			log.Println(err)
			t.Error("Ad URL test failed for input: " + tt.in)
		}
	}
}