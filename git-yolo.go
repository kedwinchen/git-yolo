package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/radovskyb/watcher"
)

func logErrror(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ePrint(err error, msg string) {
	logErrror(err)
	if err != nil {
		log.Println("ERROR START >>> ")
		log.Println(fmt.Sprintf("%s", msg))
		log.Println("ERROR END <<<")
	}
}

func exitFail(msg string) {
	// currently a wrapper, may change functionality later
	log.Fatalln(msg)
}

func exitOnError(err error, msg string) {
	logErrror(err)
	exitFail(msg)
}

func readFile(filepath string) []string {
	fileLines := []string{}

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fileLines
}

func getMessages() []string {
	messageList := []string{}
	msgFileList := []string{}

	yoloDir, err := homedir.Dir()
	logErrror(err)
	yoloDir = filepath.Join(yoloDir, ".gityolo")

	pathInfo, err := os.Stat(yoloDir)
	// exitOnError(err, fmt.Sprintf("FATAL: Directory %s does not exist!", yoloDir))
	if !os.FileMode.IsDir(pathInfo.Mode()) {
		exitFail(fmt.Sprintf("FATAL: Expected %s to be a directory", yoloDir))
	}
	// get a list of all text files in yoloDir
	err = filepath.Walk(yoloDir,
		func(path string, info os.FileInfo, err error) error {
			if err == nil {
				if strings.HasSuffix(path, ".txt") {
					pathInfo, err = os.Stat(path)
					if os.FileMode.IsRegular(pathInfo.Mode()) {
						msgFileList = append(msgFileList, path)
					}
				}
			}
			return err
		})
	logErrror(err)
	for _, theFile := range msgFileList {
		theContents := readFile(theFile)
		for _, theLine := range theContents {
			messageList = append(messageList, theLine)
		}
	}

	return messageList
}

func pickMessage(messageList *[]string, r *rand.Rand) string {
	return (*messageList)[r.Intn(len(*messageList))]
}

func runCmd(cmd *exec.Cmd) {
	output, err := cmd.CombinedOutput()
	ePrint(err, string(output))
}

func GitYolo(messageList *[]string, r *rand.Rand) {
	// add gitignore ignored files
	gitAdd := exec.Command("git", "add", ".", "--force")
	// commit with random messag
	gitCommit := exec.Command("git", "commit", "-m", pickMessage(messageList, r))
	// force push to master. what could possibly go wrong?
	// gitPush := exec.Command("git", "push", "--force", "origin", "master")
	gitPush := exec.Command("git", "push")

	runCmd(gitAdd)
	runCmd(gitCommit)
	runCmd(gitPush)
}

func runWatcher(messageList *[]string, r *rand.Rand) {
	theWatcher := watcher.New()
	defer theWatcher.Close()

	go func() {
		for {
			select {
			case event := <-theWatcher.Event:
				log.Println(event)
				GitYolo(messageList, r)
			case err := <-theWatcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	theWatcher.AddRecursive(".")
	theWatcher.Ignore(".git")
	if err := theWatcher.Start(time.Millisecond); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	messageList := getMessages()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	runWatcher(&messageList, r)
}
