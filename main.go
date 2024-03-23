package main

import (
	"fmt"
	webview "github.com/webview/webview_go"
)

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Browser")
	w.SetSize(1200, 800, webview.HintNone)
	w.SetHtml("Thanks for using Browser!")
	w.Navigate("https://go.dev")
	err := w.Bind("updateTitle", func(title string) {
		fmt.Println("Received page title:", title)
		w.SetTitle(title)
	})
	if err != nil {
		fmt.Println("Error binding function:", err)
	}
	w.Init(`
		document.addEventListener("DOMContentLoaded", function() {
			if (typeof window.updateTitle === "function") {
    			window.updateTitle(document.title)
				.then(function(result) {
        			console.log("Title updated.", result);
    			})
				.catch(function(error) {
        			console.error("Title update failed.", error);
    			});
			} else {
    			console.error("updateTitle function was not defined.");
			}
		});
	`)
	w.Run()
}
