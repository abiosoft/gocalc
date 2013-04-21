package main

import (
	gtk "github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
	"strings"
)

var (
	display   *gtk.Entry // where values are displayed
	inputMode = false
	nums      = "789/456x123-0.=+"
	operators = "/x-+="
)

// End the program
func Quit() {
	gtk.MainQuit()
}

// Action to be performed by each button, returns a handler function
func Input(b *gtk.Button) func() {
	return func() {
		if strings.Index(operators, b.GetLabel()) != -1 {
			val, _ := strconv.ParseFloat(display.GetText(), 32)
			Calculation(float32(val), b.GetLabel())
			display.SetText(GetResult())
			inputMode = false
		} else {
			if inputMode {
				display.SetText(display.GetText() + b.GetLabel())
			} else {
				display.SetText(b.GetLabel())
				inputMode = true
				if result.operator == "=" {
					Reset()
				}
			}
		}
	}
}

func main() {
	gtk.Init(&os.Args)
	display = gtk.NewEntry()
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Simple Go Calculator")
	window.Connect("destroy", Quit, nil)

	// Vertical box containing all components
	vbox := gtk.NewVBox(false, 1)

	// Menu bar
	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	// Add calculator display to vertical box
	display.SetCanFocus(false) // disable focus on calcuator display
	display.SetText("0")
	display.SetAlignment(1.0) //align text to right
	vbox.Add(display)

	// Menu items
	filemenu := gtk.NewMenuItemWithMnemonic("_File")
	menubar.Append(filemenu)
	filesubmenu := gtk.NewMenu()
	filemenu.SetSubmenu(filesubmenu)

	aboutmenuitem := gtk.NewMenuItemWithMnemonic("_About")
	aboutmenuitem.Connect("activate", func() {
		messagedialog := gtk.NewMessageDialog(
			window.GetTopLevelAsWindow(),
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO,
			gtk.BUTTONS_OK,
			"Simple Go Calculator")
		messagedialog.Response(func() {}, nil)
		messagedialog.Run()
		messagedialog.Destroy()
	},
		nil)
	filesubmenu.Append(aboutmenuitem)

	resetmenuitem := gtk.NewMenuItemWithMnemonic("_Reset")
	resetmenuitem.Connect("activate", func() { Reset(); display.SetText("0") }, nil)
	filesubmenu.Append(resetmenuitem)

	exitmenuitem := gtk.NewMenuItemWithMnemonic("E_xit")
	exitmenuitem.Connect("activate", Quit, nil)
	filesubmenu.Append(exitmenuitem)

	// Vertical box containing all buttons
	buttons := gtk.NewVBox(false, 5)

	for i := 0; i < 4; i++ {
		hbox := gtk.NewHBox(false, 5) // a horizontal box for each 4 buttons
		for j := 0; j < 4; j++ {
			b := gtk.NewButtonWithLabel(string(nums[i*4+j]))
			b.Clicked(Input(b), nil) //add click event
			hbox.Add(b)
		}
		buttons.Add(hbox) // add horizonatal box to vertical buttons' box
	}

	vbox.Add(buttons)

	window.Add(vbox)
	window.SetSizeRequest(250, 250)
	window.ShowAll()
	gtk.Main()
}
