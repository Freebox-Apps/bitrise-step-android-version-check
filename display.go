package main

import "fmt"


func displayIsMerge(err error, output string) {
	if err != nil {
		fmt.Printf("Error isMerge: %#v", err)
	}
	fmt.Printf("Commit is a Merge: %t\n ", output != "0")
}

func displayDiff(errDiff error, output string) {
	if errDiff != nil {
		fmt.Printf("Diff Error:\n%s\n", errDiff)
	} else {
		fmt.Printf("Diff Output:\n%s\n", output)
	}
}

func displayVersion(versions [][]string) {
	if len(versions) > 0 {
		fmt.Printf("===============================")
		fmt.Printf(" Version Changed  ")
		fmt.Printf("===============================\n")
		for i := 0; i < len(versions); i++ {
			fmt.Printf("\t\tVersion %d: %s", i, versions[i][1])
		}
	}else{
		fmt.Printf("===============================")
		fmt.Printf("  Same Current Version  ")
		fmt.Printf("===============================\n")
	}
	fmt.Println("")
}