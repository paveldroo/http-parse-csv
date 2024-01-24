package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type stake struct {
	Date time.Time
	Open float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}

func handle(res http.ResponseWriter, req *http.Request) {
	stakes := prs()
	err := tpl.Execute(res, stakes)
	if err != nil {
		log.Fatalln(err)
	}
}

func prs() []stake {
	var stakes []stake

	f, _ := os.Open("table.csv")
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		d, _ := time.Parse("2006-01-02", row[0])
		o, _ := strconv.ParseFloat(row[1], 64)

		stakes = append(stakes, stake{Date: d, Open: o})
	}
	return stakes
}
