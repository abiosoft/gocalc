package main

import (
	"gtk"
	"os"
	"unsafe"
	"strings"
	"strconv"
)

var (
	screen    *gtk.GtkEntry //where values are displayed
	inputMode = true
	nums      = "789/456x123-0.=+"
	operators = "/x-+="
)

//to eliminate long declarations of
//func(w *gtk.GtkWidget, args []unsafe.Pointer)
func action(f func()) gtk.CallbackFunc {
	return func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		f()
	}
}
//closes the program
func quit() gtk.CallbackFunc {
	return action(func() { gtk.MainQuit() })
}
//action to be performed by each button
func input(b *gtk.GtkButton) gtk.CallbackFunc {
	return action(func() {
		if strings.Index(operators, b.GetLabel()) != -1 {
			val, _ := strconv.Atof(screen.GetText())
			calculation(val, b.GetLabel())
			screen.SetText(getResult())
			inputMode = false
		} else {
			if inputMode {
				screen.SetText(screen.GetText() + b.GetLabel())
			} else {
				screen.SetText(b.GetLabel())
				inputMode = true
				if Result.operator == "=" {
					reset()
				}
			}
		}
	})
}
//main function
func main() {
	gtk.Init(&os.Args)
	screen = gtk.Entry()
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("Simple Go Calculator")
	window.Connect("destroy", quit(), nil)

	//vertical box containing all components
	vbox := gtk.VBox(false, 1)

	//menu bar
	menubar := gtk.MenuBar()
	vbox.PackStart(menubar, false, false, 0)

	//add screen to vertical box
	screen.SetCanFocus(false) //disable focus on calcuator screen
	vbox.Add(screen)

	//menu items
	filemenu := gtk.MenuItemWithMnemonic("_File")
	menubar.Append(filemenu)
	filesubmenu := gtk.Menu()
	filemenu.SetSubmenu(filesubmenu)

	aboutmenuitem := gtk.MenuItemWithMnemonic("_About")
	aboutmenuitem.Connect("activate", action(func() {
		messagedialog := gtk.MessageDialog(
			window.GetTopLevelAsWindow(),
			gtk.GTK_DIALOG_MODAL,
			gtk.GTK_MESSAGE_INFO,
			gtk.GTK_BUTTONS_OK,
			"Simple Go Calculator")
		messagedialog.Response(action(func() {}), nil)
		messagedialog.Run()
		messagedialog.Destroy()
	}),
		nil)
	filesubmenu.Append(aboutmenuitem)

	resetmenuitem := gtk.MenuItemWithMnemonic("_Reset")
	resetmenuitem.Connect("activate", action(func() { reset(); screen.SetText("0") }), nil)
	filesubmenu.Append(resetmenuitem)

	exitmenuitem := gtk.MenuItemWithMnemonic("E_xit")
	exitmenuitem.Connect("activate", quit(), nil)
	filesubmenu.Append(exitmenuitem)

	//vertical box containing all buttons
	buttons := gtk.VBox(false, 5)

	for i := 0; i < 4; i++ {
		hbox := gtk.HBox(false, 5) //a horizontal box for each 4 buttons
		for j := 0; j < 4; j++ {
			b := gtk.ButtonWithLabel(string(nums[i*4+j]))
			b.Clicked(input(b), nil) //add click event
			hbox.Add(b)
		}
		buttons.Add(hbox) //add horizonatal box to vertical buttons' box
	}

	vbox.Add(buttons)

	window.Add(vbox)
	window.SetSizeRequest(250, 250)
	window.ShowAll()
	gtk.Main()
}
