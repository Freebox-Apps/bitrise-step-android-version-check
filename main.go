package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

const ChangeValueRegex ="[versionName].*\\'([0-9]+[\\.]*[0-9]+[\\.]*[0-9])\\'"

const (
	DebugEnv   = "debug"
	DebugKeyOk = "yes"
	IsNewEnv   = "IS_NEW_VERSION"
	IsNewKeyOk   = "yes"
	IsNewKeyKo   = "no"
)

func main() {

	gitRep := getGitRepository()
	commits := getCommitsToCompare(gitRep)
	output := getDiff(gitRep, os.Getenv("build_gradle_path"), commits[0], commits[1])
	changes := extractVersions(output)
	displayVersion(changes)

	setOutput(len(changes) > 1)

	os.Exit(0)
}

func setOutput(isNew bool) {
	var value string
	if isNew{
		value = IsNewKeyOk
	}else{
		value = IsNewKeyKo
	}
	cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", IsNewEnv, "--value", value).CombinedOutput()

	if err != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
		os.Exit(1)
	}
}

func extractVersions(output string) [][]string {
	regex := regexp.MustCompile(ChangeValueRegex)
	versions := regex.FindAllStringSubmatch(output, -1)
	return versions
}

func isDebug() bool {
	return os.Getenv(DebugEnv) == DebugKeyOk
}