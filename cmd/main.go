package main

import servStart "wheather/pkg/handler"

func main() {
	err := servStart.Start()
	if err != nil {
		return
	}
}
