package main

import (
	"log"
	"math/rand"
	"os/exec"
	"time"

	"github.com/radovskyb/watcher"
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
	watcher := watcher.New()
	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Event:
				log.Println("event:", event)
				log.Println(pickMessage(messageList, r))
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	watcher.AddRecursive(".")
	watcher.Ignore(".git")
}

func main() {
	messageList := getMessages()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	runWatcher(&messageList, r)
}
