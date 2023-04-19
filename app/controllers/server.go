package controllers

import (
	"fmt"
	"net/http"
	"qiita_app/config"
)

// WebServerを起動する
func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// UR取り込み、Handler実行
	http.HandleFunc("/", top)
	// 問題ありの場合、エラーを返す
	// 第一引数はポート番号、何も指定しない場合はlocalhost
	// 第二引数をnilにした場合、登録してないurlにアクセスしたらデフォルトで404を返す
	return http.ListenAndServe(":"+config.Config.PortNum, nil)
}

// ハンドラ関数で使われる部分を共通化
func generateHTML(w http.ResponseWriter, data interface{}, fileNames ...string) {
	var files []string
	// 引数のファイルNameを取り出す
	for _, file := range fileNames {
		// sprintfでテンプレート化
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// テンプレートをキャッシュする
	// エラーの場合はPanic
	// templates := template.Must(template.ParseFiles(files...))
	// 実行、テンプレート指定
	// defineしたファイルは明示的にファイル指定
	// templates.ExecuteTemplate(w, "layout", data)
}
