package main

import (
	"fmt"
	"io/ioutil"
	"time"
	"os"
	"errors"
	"os/exec"
	"syscall"
	"path/filepath"

	"github.com/ropog3ovomou/systray"
	"github.com/ropog3ovomou/DisposableVM/bin/go/icon"
	"github.com/sqweek/dialog"
	"github.com/ropog3ovomou/beeep"
)

func main() {
	onExit := func() {
	}

	systray.Run(onReady, onExit)
}


func onReady() {
	appdata, _ := os.UserConfigDir()
	mainDir := appdata + "/DisposableVM"

	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("DisposableVM")
	systray.SetTooltip("Disposable VM")
	xselect := systray.AddMenuItem("Select VM", "Quit the whole app")
	systray.AddSeparator()
	xstart := systray.AddMenuItem("Start", "my home")
	xdispose := systray.AddMenuItem("Dispose", "my home")
   log := func (action string, msg string) {
		now := time.Now()
		ioutil.WriteFile(mainDir + "/" + fmt.Sprintf(`%d_%s.log`, now.Unix(), action), []byte(msg), 0644)
   }
	run := func (action string, arg string, title string, msg string) {
		cmd := exec.Command("bin/shell/bash.exe", "-c", "./disposable " + action + " \"\\\"" + arg + "\\\"\"")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		bs,err := cmd.CombinedOutput()
		if err == nil {
			xstart.Enable()
			xdispose.Enable()
			beeep.Notify(title, msg, "icon.png")
         log(action, fmt.Sprintf("%s", bs))
		} else {
			dialog.Message("%s\n\n%s", bs, err).Error()
		}
	}
	systray.AddSeparator()
	xquit := systray.AddMenuItem("Quit", "Quit the whole app")
	if dat, err := os.ReadFile(mainDir + "/Current"); err == nil {
		// path/to/whatever exists
		xselect.SetTitle("= " + filepath.Base(string(dat)) + " =")
	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		xstart.Disable()
		xdispose.Disable()
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		dialog.Message("%s", err).Error()
	}
	for {
		select {
		case <-xselect.ClickedCh:
			filename, err := dialog.File().Filter("VMWare Virtual Machine (*.vmx)", "vmx").Load()
			if err == nil {
				run("configure", filename, "VM Selected",
		 		   "Virtual machine template was selected successfully")
				xselect.SetTitle("= " + filepath.Base(filepath.Dir(filename)) + " =")
			} else {
			// 	dialog.Message("%s", err).Error()
			}
		case <-xstart.ClickedCh:
			run("start", "", "VM Created", "Disposable VM was created successfully. Trying to start it...")
			// ok := dialog.Message("%s", "Do you want to continue?").Title("DisposableVM").YesNo()
			// if ok {
			// 	dialog.Message("%s", "That's right!").Title("OK").YesNo()
			// 	//open.Run("https://www.getlantern.org")
			// }
		case <-xdispose.ClickedCh:
			run("dispose", "", "VM Wiped Out", "Virtual Machine was disposed successfully")
			// dialog.Message("%s", "Dispose").Title("DisposableVM").YesNo()
		case <-xquit.ClickedCh:
			systray.Quit()
			return
		}
	}

}
