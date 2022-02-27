package bind

import "github.com/webview/webview"

func QuitBindCollection(w webview.WebView) {
	w.Bind("quit", QuitBind(w))
}

func QuitBind(w webview.WebView) interface{} {
	return func() {
		w.Terminate()
	}
}
