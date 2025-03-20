# Golang Handson

https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E-%E3%83%8F%E3%83%B3%E3%82%BA%E3%82%AA%E3%83%B3-%E6%8E%8C%E7%94%B0%E6%B4%A5%E8%80%B6%E4%B9%83/dp/4798063991

## Note

### Session 4

When following the setup procedure outlined in this book, I encountered the following error while running `go run main.go`:

```sh
# github.com/go-gl/gl/v3.1/gles2
../../../../go/pkg/mod/github.com/go-gl/gl@v0.0.0-20190320180904-bf2b1f2f34d7/v3.1/gles2/package.go:38:11: fatal error: 'KHR/khrplatform.h' file not found
   38 |  #include <KHR/khrplatform.h>
      |           ^~~~~~~~~~~~~~~~~~~
1 error generated.
```

To resolve this issue, I switched to the latest version of Fyne (version 2) for the setup and confirmed that it works correctly.

```sh
# Run the command as described in the official documentation.
# https://docs.fyne.io/started/
go get fyne.io/fyne/v2@latest
go install fyne.io/fyne/v2/cmd/fyne@latest

fyne --version
#=> fyne version v2.5.5
```

```go
// 4-1
import (
	// Specify installed version 2.
	// https://github.com/june23hy/golang-hands-on?tab=readme-ov-file
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	w.SetContent(widget.NewLabel("Hello Fyne!"))

	w.ShowAndRun()
}
```

```sh
go mod tidy
```
