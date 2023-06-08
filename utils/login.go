package utils

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/rivo/tview"
)

type Student struct {
	Name    string    `json:"name"`
	NIM     string    `json:"nim"`
	Shift   string    `json:"shift"`
	Session time.Time `json:"session"`
}

func Login() {
	app := tview.NewApplication()

	// Create a new form.
	form := tview.NewForm()

	// Add inputs for name and NIM.
	form.AddInputField("Nama", "", 40, nil, nil).
		AddInputField("NIM", "", 40, nil, nil).
		AddDropDown("Shift", []string{"-Pilih shift-", "A", "B", "C", "D", "E", "F", "G", "H", "I"}, 0, func(option string, optionIndex int) {})

	// Add a button to submit the form.
	form.AddButton("LOGIN", func() {
		// Retrieve the values of the input fields.
		name := form.GetFormItemByLabel("Nama").(*tview.InputField).GetText()
		nim := form.GetFormItemByLabel("NIM").(*tview.InputField).GetText()
		_, shift := form.GetFormItemByLabel("Shift").(*tview.DropDown).GetCurrentOption()
		// Save the student data here.
		student := Student{
			Name:    name,
			NIM:     nim,
			Shift:   shift,
			Session: time.Now().Local(),
		}

		// Marshal the Student struct to JSON.
		data, err := json.Marshal(student)
		if err != nil {
			panic(err)
		}

		// Write the JSON to a file.
		err = ioutil.WriteFile("config.json", data, 0644)
		if err != nil {
			panic(err)
		}
		// Close the application.
		app.Stop()
		println("\033[32m[+] Login success\t\033[0m\n")

	})

	// Add the form to the application.
	if err := app.SetRoot(form, true).Run(); err != nil {
		panic(err)
	}
}
