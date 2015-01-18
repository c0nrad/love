package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	SMSMode = false
)

func main() {
	//LearnFile("./data/whatisai");
	//LearnFile("./data/lady_lazarus")
	LearnFile("./data/pickup")
	//LearnFile("./data/bell_jar");
	LearnFile("./data/2chainz")
	Rap()
	//dumpBrain();

	if SMSMode {
		fmt.Println("Starting SMS mode")
		setupTwilioListeners()
		return
	}

	//Chat();
}

func ReadThoughts() {
	currLink := Links[0]
	thought := ""
	for {
		thought, currLink = NextThought(currLink, 15)
		fmt.Println(thought)

		time.Sleep(2000 * time.Millisecond)
	}
}

func Chat() {
	currLink := Links[0]
	thought := ""
	bio := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		question, _, err := bio.ReadLine()
		if err != nil {
			fmt.Println(err)
		}
		TriggerSentence(string(question), 2)

		thought, currLink = NextThought(currLink, 15)
		fmt.Println(thought)

		currLink = NextLink(currLink)
	}

}

func Rap() {
	currLink := Links[0]
	thought := ""
	for i := 0; i < 20; i++ {
		thought, currLink = NextThought(currLink, 6)
		fmt.Println(thought)
		currLink = NextLink(currLink)

	}
}
