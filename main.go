package main

import (
	"fmt"
	webview "github.com/webview/webview_go"
	"os"
)

var err error

func main() {
	// get the first command line argument
	// if it is not empty, use it as the URL
	url := ""
	if len(os.Args) > 1 {
		url = os.Args[1]
	}
	fmt.Println("URL:", url)
	// check if the URL starts with "http://", if not, add "https://"
	if url != "" {
		url = testUrl(url)
	}
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Browser")
	w.SetSize(1200, 800, webview.HintNone)
	err = w.Bind("updateTitle", func(title string) {
		fmt.Println("Received page title:", title)
		w.SetTitle(title)
	})
	if err != nil {
		fmt.Println("Error binding function:", err)
	}
	err = w.Bind("navigate", func(url string) {
		url = testUrl(url)
		fmt.Println("Navigate to:", url)
		w.Navigate(url)
	})
	if err != nil {
		fmt.Println("Error binding function:", err)
	}
	if url != "" {
		w.Navigate(url)
	} else {
		w.SetHtml("Thanks for using Browser!")
	}
	w.Init(`
        document.addEventListener("DOMContentLoaded", function() {
			if (typeof window.updateTitle === "function") {
    			window.updateTitle(document.title)
				.catch(function(error) {
        			console.error("Title update failed.", error);
    			});
			} else {
              console.error("updateTitle undefined.");
			}
		});
		document.addEventListener("keydown", function(event) {
            if (event.ctrlKey && event.altKey && event.key === "u") {
                promptForURL();
            }
		});
		function promptForURL() {
            var url = prompt("Open page:", "");
			if (url) {
                window.navigate(url);
            }
		}
	`)

	w.Run()
}

func testUrl(url string) string {
	if len(url) < 7 {
		url = "https://" + url
	} else if url[:7] != "http://" {
		url = "https://" + url
	}
	return url
}
