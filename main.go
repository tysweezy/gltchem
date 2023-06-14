package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"log"
)

type Task struct {
	name     string
	duration string
}

func Boop(soundPath string) {
	afplay, err := exec.LookPath("afplay")

	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(afplay, soundPath)
	cmd.Run()
}

func Notify(title string, message string, soundName string) {
	osa, err := exec.LookPath("osascript")

	if err != nil {
		log.Fatal(err)
	}

	script := fmt.Sprintf("display notification %q with title %q sound name %q", message, title, soundName)
	cmd := exec.Command(osa, "-e", script)
	cmd.Run()
}

func TaskFile(file string) []string {
	var fileTasks []string
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", file)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fileTasks = append(fileTasks, scanner.Text())
	}

	fmt.Println("tasks: ", fileTasks)

	return fileTasks
}

func main() {
	fmt.Println("Running boop")
	// Notify("Reminder", "Do a task", "Blow")
	// Boop("/System/Library/Sounds/Blow.aiff")
	tasks := TaskFile("tasks.txt")
	for _, t := range tasks {
		taskString := strings.Split(t, "|")
		task := Task{name: taskString[0], duration: taskString[1]}
		Notify("Reminder", task.name, "Blow")
	}
}
