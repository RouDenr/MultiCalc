package main

import (
	"github.com/webview/webview"
	"net/http"
	"io/ioutil"
	"fmt"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	expression, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid expression", http.StatusBadRequest)
		return
	}

	// Вставьте здесь код для вычисления выражения из модели
	result := calculateExpression(string(expression))

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}

func calculateExpression(expression string) string {
	// Вставьте здесь код для вычисления выражения из модели (C++ модель или обертки)
	// В этом примере вернем просто результат выражения (как будто без вычисления)
	return expression
}

func main() {
	http.HandleFunc("/calculate", calculateHandler)

	// Запустим WebView с HTML-файлом, который будет служить представлением
	debug := true // Установите в false для скрытия консоли разработчика
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle("Simple Calculator")
	w.SetSize(400, 600, webview.HintNone)
	w.Navigate("file://" + "index.html")

	// Запускаем главный цикл WebView
	w.Run()
}
