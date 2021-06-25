package ui

import (
	// "fmt"
	"fmt"
	"log"
	"os"

	// "math"
	// "time"
	// "fmt"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/sarthakvk/onepass/backend"
	"golang.org/x/term"

)

var ftc map[string]interface{}
var KeyEvents <-chan ui.Event
func Base() {
	ftc = map[string]interface{}{
		"login":           login,
		"register":        register,
		"reset passcode":  resetPasscode,
		"search password": searchPassword,
		"add password":    addPassword,
		"remove password": removePassword,
		"set smtp":        setSMTP,
		"logout":          logout,
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	KeyEvents = ui.PollEvents()
	homeScreen()

}

func homeScreen() {
	w, h, _ := term.GetSize(0)
	p := widgets.NewList()
	p.SetRect(0, 0, w, h)
	p.Title = "Welcome"
	_, err := os.Open(backend.DF_NAME)
	if err != nil {
		p.Rows = []string{"1. Register"}
	} else {
		p.Rows = []string{
			"1. Login",
			"2. Forgot Passcode",
		}
	}
	p.SelectedRow = 0
	p.SelectedRowStyle.Bg = ui.ColorGreen
	ui.Render(p)

	for e := range KeyEvents {
		switch e.ID {
		case "<Escape>":
			ui.Clear()
			return
		case "<Down>":
			p.SelectedRow = (p.SelectedRow + 1) % len(p.Rows)

		case "<Up>":
			p.SelectedRow = p.SelectedRow - 1
			if p.SelectedRow < 0 {
				p.SelectedRow += len(p.Rows)
			}

		case "<Enter>":
			takeAction(p)
		}

		p.Title = "Welcome"
		if backend.Logged_in {
			p.Rows = []string{
				"1. Search Password",
				"2. Add Password",
				"3. Remove Password",
				"4. Set SMTP",
				"5. Logout",
				"6. Reset Passcode",
			}
		} else {
			p.Rows = []string{"1. Login", "2. Reset Passcode"}
		}
		ui.Render(p)

	}
}

func takeAction(i *widgets.List) {
	p := i.Rows[i.SelectedRow]
	inp := ""
	for _, v := range strings.Split(p, " ")[1:] {
		inp += v + " "
	}
	inp = inp[:len(inp)-1]
	inp = strings.ToLower(inp)
	execute(inp, i)
}

func execute(inp string, p *widgets.List) {
	ftc[inp].(func(*widgets.List))(p)
}

func login(p *widgets.List) {
	p.Title = "Login"
	p.Rows = []string{
		"Enter Password: ",
	}
	ui.Render(p)
	var pass string
	fmt.Scan(&pass)
	backend.Login([]byte(pass))
}
func register(p *widgets.List) {
	p.Title = "Password length should be in range of [6,12]"
	p.Rows = []string{
		"Enter Passcode: ",
	}
	ui.Render(p)
	var pass string
	fmt.Scan(&pass)
	p.Rows = []string{
		"Confirm Passcode: ",
	}
	ui.Render(p)
	var cpass string
	fmt.Scan(&cpass)

	if pass == cpass && len(pass) > 6 && len(pass) < 12 {
		backend.Register([]byte(pass))
	} else {
		register(p)
	}
}
func resetPasscode(p *widgets.List)  {
	p.Title = "Password length should be in range of [6,12]"
	p.Rows = []string{
		"Old Passcode: ",
	}
	ui.Render(p)
	var pass string
	fmt.Scan(&pass)
	if backend.Login([]byte(pass)) {
		register(p);
	}
}
func addPassword(p *widgets.List)    {
	p.Title = "Add new Password"
	p.Rows = []string{"Name: \nURL: \nPassword: "}
	// for e := range KeyEvents {
		// termbox.
	// }
	var n,u,pa string
	fmt.Scan(&n,&u,&pa)
	backend.AddPassword(n,u,pa)
	
}
func removePassword() {}
func searchPassword() {}
func setSMTP()        {}
func logout()         {}
