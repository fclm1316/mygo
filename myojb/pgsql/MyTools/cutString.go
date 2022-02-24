package MyTools

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

func CutClearString(schema string, fname string, sqlstr string) {
	s := strings.NewReader(sqlstr)
	br := bufio.NewReader(s)
	var cleanSql bytes.Buffer
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		tocleanstr := CleanStr(strings.ToUpper(strings.Split(string(line), "--")[0]))
		cleanSql.WriteString(tocleanstr)
	}

	compileSqlstr1 := fmt.Sprintf("INSERT INTO (%s.%s).*?;|INSERT INTO (%s).*?;", strings.ToUpper(schema),
		strings.ToUpper(fname), strings.ToUpper(fname))
	complieSql := regexp.MustCompile(compileSqlstr1)
	compileMatchSql := complieSql.FindAllString(cleanSql.String(), -1)
	for a := range compileMatchSql {
		fmt.Println(compileMatchSql[a])
	}

}

func CleanStr(str string) string {
	str = strings.ReplaceAll(str, "(", "")
	// 括号内为 $1
	re1 := regexp.MustCompile(`''([a-zA-Z0-9_]+)''`)
	str = re1.ReplaceAllString(str, "'$(1)'")
	return str
}
