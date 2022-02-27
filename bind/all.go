package bind

import (
	"github.com/webview/webview"
)

func AllBindCollection(w webview.WebView) {
	FileBindCollection(w)
	GroupBindCollection(w)
	ImageBindCollection(w)
	MessageBindCollection(w)
	QuitBindCollection(w)
	RoomBindCollection(w)
	UserBindCollection(w)
}
