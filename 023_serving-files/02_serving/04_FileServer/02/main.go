package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	//Hàm này sẽ loaị bỏ phần đường dẫn /resource trong câu request và export files trong thư muc ./assets
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/toby.jpg">`)
}

/*

./assets/toby.jpg

*/
