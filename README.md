# git-yolo

You only live once.

As they say, commit often, perfect later (I feel like I'm missing a part of the phrase. Oh well, it's probably nothing)...

## Installation

(if the install script is not available yet)

1. Download the source code (or download the binary and skip to step 3)
2. Compile using `go build git-yolo.go`
3. Add the resulting `git-yolo` binary to your path
4. Add an alias for `git yolo` in your [global] `.gitconfig`
5. Copy the `.gityolo` directory to your home directory (wherever that is)

## Usage

In any git directory, run `git yolo`.
This starts `git-yolo` as a daemon-like process (currently, at time of writing, it blocks)

## Why

Why not?

Also, it gave me a reason to write a project in Go.

## WARNING

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

I am not responsible for any consequences that may result from using this software,
professional or otherwise.
