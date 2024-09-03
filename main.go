package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	//? <-uslanozan-> ?\\

	//! < SELECT A BROSWER FOR EACH OS BELOW AND TYPE  >

	//? ---Google Chrome---
	//*  -Windows-
	// exec.Command("cmd", "/C", "start chrome --incognito", <SCANNER NAME>.Text()).Start()
	//*  -macOS-
	// exec.Command("open", "-a", "Google Chrome", "--args", "--incognito", <SCANNER NAME>.Text()).Start()
	//*  -Linux-
	// exec.Command("google-chrome", "--incognito", <SCANNER NAME>.Text()).Start()

	//? ---Firefox---
	//*  -Windows-
	// exec.Command("cmd", "/C", "start firefox --private-window", <SCANNER NAME>.Text()).Start()
	//* -macOS-
	// exec.Command("open", "-a", "Firefox", "--args", "--private-window", <SCANNER NAME>.Text()).Start()
	//* -Linux-
	// exec.Command("firefox", "--private-window", <SCANNER NAME>.Text()).Start()

	//? ---Brave---
	//* -Windows-
	// exec.Command("cmd", "/C", "start brave --incognito", <SCANNER NAME>.Text()).Start()
	//* -macOS-
	// exec.Command("open", "-a", "Brave Browser", "--args", "--incognito", <SCANNER NAME>.Text()).Start()
	//* -Linux-
	// exec.Command("brave-browser", "--incognito", <SCANNER NAME>.Text()).Start()

	//? ---Microsoft Edge---
	//*  -Windows-
	// exec.Command("cmd", "/C", "start msedge --inprivate", <SCANNER NAME>.Text()).Start()
	//*  -macOS-
	// exec.Command("open", "-a", "Microsoft Edge", "--args", "--inprivate", <SCANNER NAME>.Text()).Start()
	//*  -Linux-
	// exec.Command("microsoft-edge", "--inprivate", <SCANNER NAME>.Text()).Start()

	//---------------------------------------------\\
	//? Excel and Zenkit
	dependencies, err := os.Open("sources/dependencies.txt")
	if err != nil {
		fmt.Println("dependencies.txt cannot be opened !")
	}

	scannerD := bufio.NewScanner(dependencies)

	for scannerD.Scan() {
		//! Change here which Operating System you use
		exec.Command("cmd", "/C", "start msedge --inprivate", scannerD.Text()).Start()

		time.Sleep(300 * time.Millisecond)
	}

	dependencies.Close()
	//---------------------------------------------\\
	//? Tor must be here

	//---------------------------------------------\\
	//? Forums
	clearnet, err := os.Open("sources/clearnet.txt")
	if err != nil {
		fmt.Println("clearnet.txt cannot be opened !")
	}

	scannerC := bufio.NewScanner(clearnet)

	for scannerC.Scan() {
		//! Change here which Operating System you use
		exec.Command("cmd", "/C", "start msedge --inprivate", scannerC.Text()).Start()

		time.Sleep(300 * time.Millisecond)
	}
	clearnet.Close()
	//---------------------------------------------\\

}
