package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

type pribors [][]string

var Tbpribors map[string]*pribors //Приборы сгруппированы по первому полю файла
var T *template.Template

func indexHandler(w http.ResponseWriter, r *http.Request) {
	T.ExecuteTemplate(w, "finder", nil)
}

func findHandler(w http.ResponseWriter, r *http.Request) {
	ftext := r.FormValue("ftext")  //Получили текст введенный пользователем
	result, ok := Tbpribors[ftext] //Ищем список приборов по ключевому слову
	if ok {
		T.ExecuteTemplate(w, "finder", *result)
	} else {
		T.ExecuteTemplate(w, "index", "Ничего не найдено")
	}
}
func main() {
	err := errors.New("sd")
	T, err = template.ParseFiles("template/index.html", "template/finder.html", "template/result.html")
	if err != nil {
		log.Panic(err.Error())
	}

	go func() {
		for {
			Tbpribors = make(map[string]*pribors)
			lines := readCSV("test.csv")
			for _, line := range *lines {
				keys := strings.Split(line, ",")
				if _, ok := Tbpribors[keys[0]]; ok {
					value := append(*Tbpribors[keys[0]], keys)
					Tbpribors[keys[0]] = &value
				} else {
					p := make(pribors, 0)
					p = append(p, keys)
					Tbpribors[keys[0]] = &p
				}
			}
			log.Println("Файл обновлен")
			time.Sleep(time.Minute)
		}
	}()

	//.......
	log.Println("Выполняется разбор файла с данными...")
	time.Sleep(time.Second * 5) //Подождем перед запуском пока выполнится парсинг файла
	log.Println("Слушаем порт 8080")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/FindPr", findHandler)
	http.ListenAndServe(":8080", nil)
}
