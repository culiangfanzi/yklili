package test

import (
	"beegostudy/util"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func Test_AnalysisGoTag(t *testing.T) {
	str := util.FileToString(`E:\github\go\gostudy\src\beegostudy\views\platform\user\test.html`)
	t.Fatal(util.AnalysisGoTag(str))
}

func Test_String(t *testing.T) {
	str := "required;rangelength(6,12);email;max(10)"
	strs := strings.Split(str, ";")
	for _, s := range strs {
		re := regexp.MustCompile(`(\()(.*)(\))`)
		//st := re.FindStringSubmatch(s)
		st := re.ReplaceAllString(s, ":$2")
		fmt.Printf("%s\n", st)
	}
	fmt.Println(strings.Index("asd,asd", "."))
	t.Fatal("")
}