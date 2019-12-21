package screen

import (
	"github.com/creack/pty"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

//var ctrlA_D = exec.Command("\001d")

const (
	Detached    = true
	NotDetached = false
)

type Screen struct {
	PID    string
	Name   string
	Time   string
	Host   string
	Status bool
}

func Detach(tty string) error {
	return exec.Command("screen", "-d", tty).Run()
}

func Kill(tty string) error {
	return exec.Command("screen", "-S", tty, "-X", "quit").Run()
}

func List() ([]Screen, error) {
	var listScreen []Screen
	var cmd *exec.Cmd = exec.Command("screen", "-ls")
	var output, err = cmd.Output()
	if err != nil {
		return nil, err
	}
	var list []string = strings.Split(strings.Trim(string(output), "\n"), "\n")
	for i := 0; i < len(list); i++ {
		var values []string = strings.Split(strings.TrimSpace(list[i]), "\t")
		var pidNameHost []string = strings.Split(values[0], ",")
		var screen Screen = Screen{
			PID:  pidNameHost[0],
			Name: pidNameHost[1],
			Host: func() string {
				if len(pidNameHost) == 3 {
					return pidNameHost[2]
				}
				return ""
			}(),
			Time: values[1][1 : len(values[1])-1],
			Status: func() bool {
				if values[2] == "(Detached)" {
					return Detached
				} else {
					return NotDetached
				}
			}(),
		}
		listScreen = append(listScreen, screen)
	}
	return listScreen, nil
}

func Create(tty string, a ...string) error {
	var (
		args = append([]string{"-S", tty}, a...)
		c    = exec.Command("screen", args...)
	)

	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}
	defer func() { _ = ptmx.Close() }()

	go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
	_, _ = io.Copy(os.Stdout, ptmx)

	//go Detach(tty)
	//ptmx, _ = pty.Start(ctrlA_D)
	return nil
}

func Execute(tty string, a ...string) error {
	var err error
	go func() { err = Create(tty, a...) }()
	time.Sleep(time.Millisecond * 500)
	Detach(tty)
	return nil
}

func View(tty string) string {
	return ""
}
