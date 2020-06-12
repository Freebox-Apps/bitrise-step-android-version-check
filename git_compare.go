package main

import (
	"fmt"
	"github.com/Freebox-CI/bitrise-step-android-version-check/git"
	"os"
)

func getGitRepository() git.Git {
	gitRep, errNew := git.New(".")
	if errNew != nil {
		fmt.Printf("Couldn't find repository error: %#v", errNew)
		os.Exit(1)
	}
	return gitRep
}

func getCommitsToCompare(gitRep git.Git) [2]string {
	var commits [2]string
	if isMerge(&gitRep, "HEAD") {
		commits[0] = "HEAD^1"
		commits[1] = "HEAD^2"
	} else {
		commits[0] = "HEAD"
		commits[1] = "HEAD^1"
	}
	return commits
}

func getDiff(gitRep git.Git, buildGradlePath string, commitOne string, commitTwo string) string {
	debug := isDebug()
	diff := gitRep.Diff(commitOne, commitTwo, buildGradlePath, "--unified=0")
	if debug {
		fmt.Printf("Executing Command: %s\n", diff.GetCmd().Args)
	}
	output, errDiff := diff.RunAndReturnTrimmedCombinedOutput()
	if debug {
		displayDiff(errDiff, output)
	}
	return output
}

func isMerge(g *git.Git, commit string) bool {
	output, err := g.RevList(commit, "--no-walk", "--count", "--merges").RunAndReturnTrimmedOutput()
	displayIsMerge(err, output)
	if err != nil {
		os.Exit(1)
	}
	return output != "0"
}