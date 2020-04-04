# FILE:		Makefile
# PROJECT:	git-yolo
# ORIGINATOR:	CHEN, Kedwin
# EMAIL:	(redacted)
# DESCRIPTION:	GNU Makefile for the git-yolo project

MAIN := git-yolo
SRC := ${MAIN}.go
BIN_LINUX := ${MAIN}-linux-amd64
BIN_WINDOWS := ${MAIN}-windows-amd64.exe
BIN_MACOS := ${MAIN}-macos-amd64
BIN_ALL := ${BIN_LINUX} ${BIN_WINDOWS} ${BIN_MACOS}


.PHONY: ${MAIN}
${MAIN}: clean linux windows macos

.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 go build -buildmode=pie -o ${BIN_LINUX} ${SRC}

.PHONY: windows
windows:
	GOOS=windows GOARCH=amd64 go build -o ${BIN_WINDOWS} ${SRC}

.PHONY: macos
macos:
	GOOS=darwin GOARCH=amd64 go build -o ${BIN_MACOS} ${SRC}

.PHONY: clean
clean:
	rm -f ${BIN_ALL}
