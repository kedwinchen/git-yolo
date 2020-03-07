package main

import (
	"log"
	"math/rand"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
)

func getMessages() []string {
	return []string{"stuff"}
}

func pickMessage(messageList *[]string, r *rand.Rand) string {
	return (*messageList)[r.Intn(len(*messageList))]
}

func runCmd(cmd *exec.Cmd) {
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

func GitYolo(messageList *[]string, r *rand.Rand) {
	gitAdd := exec.Command("git", "add", ".", "-f")
	gitCommit := exec.Command("git", "commit", "-m", pickMessage(messageList, r))
	gitPush := exec.Command("git", "push", "--force", "origin", "master")

	runCmd(gitAdd)
	runCmd(gitCommit)
	runCmd(gitPush)
}

func runWatcher(messageList *[]string, r *rand.Rand) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(".")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {
	messageList := getMessages()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	runWatcher(&messageList, r)
}
