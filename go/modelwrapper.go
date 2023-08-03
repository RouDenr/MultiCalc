package main

/*
#cgo CXXFLAGS: -std=c++17
#include "modelwrapper.h"
*/
import "C"

import (
	"github.com/webview/webview"
	//"unsafe"
	"net/http"
	//"strconv"
)

func calculateExpression(expression string) string {
	// В этом примере мы просто вернем результат выражения без реальных вычислений
	// Здесь вы можете вызвать функцию Calculate из динамической библиотеки C++,
	// которая выполнит реальные вычисления.
	return expression
}

func main() {
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		expression := r.FormValue("expression")
		result := calculateExpression(expression)

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(result))
	})

	// Запускаем сервер для обработки вычислений
	go http.ListenAndServe(":8080", nil)

	// Запускаем WebView с HTML-файлом, который будет служить представлением
	debug := true // Установите в false для скрытия консоли разработчика
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle("Simple Calculator")
	w.SetSize(400, 600, webview.HintNone)
	w.Navigate("http://localhost:8080")

	// Запускаем главный цикл WebView
	w.Run()
}
