package bind

import "github.com/webview/webview"

func QuitBind(w webview.WebView) interface{} {
	return func() {
		w.Terminate()
	}
}
