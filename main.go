package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"os"
	"runtime"
	"strings"
)

func main() {
	var formatBtnName = "Format"
	var convertBtnName = "ConvertToJSON"
	var pasteBtnName = "Paste"
	var copyBtnName = "Copy"

	if strings.EqualFold(runtime.GOOS, "windows") {
		os.Setenv(FONT_ENV_NAME, FONT_WIN_MEIRYO_PATH)
		formatBtnName = "整形"
		convertBtnName = "JSON化"
		pasteBtnName = "貼り付け"
		copyBtnName = "コピー"
	}

	myApp := app.New()

	myWindow := myApp.NewWindow(APP_NAME + " " + VERSION)
	myWindow.Resize(fyne.NewSize(INTI_WINDOW_WIDTH, INTI_WINDOW_HEIGHT))

	input := widget.NewMultiLineEntry()

	if DEBUG {
		input.SetText(DEBUG_SAMPLE_JSON)
	}

	output := widget.NewMultiLineEntry()
	output.Disable()

	formatBtn := widget.NewButton(formatBtnName, func() {
		it := []byte(input.Text)

		j, err := FormatJSON(it, PREFIX_DEFAULT, INDENT_DEFAULT)
		if err != nil {
			output.SetText("")
			fmt.Println(err)
			return
		}

		output.SetText(string(j))
	})
	formatBtn.Disable()

	toJSONBtn := widget.NewButton(convertBtnName, func() {
		output.SetText(convJSON(input.Text))
	})
	toJSONBtn.Disable()

	pasteBtn := widget.NewButton(pasteBtnName, func() {
		clipboardTest, err := clipboard.ReadAll()
		if err != nil {
			output.SetText("")
			fmt.Println(err)
			return
		}

		input.SetText(clipboardTest)
	})

	copyBtn := widget.NewButton(copyBtnName, func() {
		if err := clipboard.WriteAll(output.Text); err != nil {
			output.SetText("")
			fmt.Println(err)
			return
		}
	})

	input.OnChanged = func(v string) {
		if strings.EqualFold(v, "") {
			formatBtn.Disable()
			toJSONBtn.Disable()
		}

		byteValue := []byte(v)
		if j, err := FormatJSON(byteValue, PREFIX_DEFAULT, INDENT_DEFAULT); err != nil {
			formatBtn.Disable()
			output.SetText(convJSON(v))
			toJSONBtn.Enable()
		} else {
			output.SetText(string(j))
			formatBtn.Enable()
			toJSONBtn.Disable()
		}
	}

	buttonContainer := container.NewGridWrap(fyne.NewSize(125, 50), formatBtn, toJSONBtn, pasteBtn, copyBtn)
	mainContainer := container.NewGridWithColumns(2, input, output)

	c := container.New(layout.NewBorderLayout(buttonContainer, nil, nil, nil), buttonContainer, mainContainer)

	myWindow.SetContent(c)
	myWindow.Show()

	myApp.Run()
}

func convJSON(text string) string {
	xml, err := XML2JSON(text, PREFIX_DEFAULT, INDENT_DEFAULT)
	if err != nil {
		return ""
	}
	return string(xml)
}