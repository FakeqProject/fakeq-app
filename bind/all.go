package bind

import (
	"github.com/webview/webview"
)

func AllBindCollection(w webview.WebView) {
	w.Bind("quit", QuitBind(w))
}
