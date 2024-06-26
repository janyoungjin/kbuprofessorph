package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ReadJson(filename string) (map[string]interface{}, error) {
	file, err := os.Open(filename) //파일을 연다
	if err != nil {
		return nil, err
	}
	defer file.Close()                   //프로그램이 끝나면 파일을 닫는다
	content, err := ioutil.ReadAll(file) //파일 내용을 읽는다
	if err != nil {
		return nil, err
	}

	var info map[string]interface{}      //info라는 map을 만든다
	err = json.Unmarshal(content, &info) //json내용을 map으로 저장한다
	if err != nil {
		return nil, err
	}
	return info, nil
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var name string
		switch r.Method {
		case "POST":
			name = r.PostFormValue("name")
		default:
			http.Error(w, "접근 불가 메서드", http.StatusMethodNotAllowed)
			return
		}
		fmt.Println(name)
		professormap, err := ReadJson("professorinfo.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
			return
		}
		value, ishas := professormap[name]
		if !ishas {
			http.Error(w, fmt.Sprintf("%s 교수님을 찾을수 없습니다.", name), http.StatusNotFound)
			return
		}
		fmt.Fprint(w, value)

	})
	http.ListenAndServe(":8000", nil)
}
