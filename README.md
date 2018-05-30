### Q. What is the best pattern to implement continuous integration and deployment?

#### A. Continuous integration, also know as CI, is the pattern of having developers frequently check in their code and commonly run steps such as run tests, scan for vulnerabilities, and deploy code.

There are a few common patterns for this. The one we will advocate today is orchestrating a series of independent automation steps. This can be done in a variety of technologies. The orchestration can be open source Jenkins, Kubernetes, or a variety of vendor technologies including CircleCI and others. The automation steps can be groovy (Jenkins), python or plain ol shell script.

For demonstration purposes, setting up CI for a new repo can be as simple as scripting a few steps in shell script and triggering them from Jenkins.

### 1.	Get Hello world/Concept working

Lets start with a simple application. It could be anything we could build from a shell script, in this case it will be a hello world Golang app. In our case we will create a simple Golang function to add one to a passed in number to provide testable functionality as well.

addone.go
```Golang
package mathhelp

func addone(x int) int {
	x = x + 1
	return x
}
```

### 2. Get a test working

Ideally you would be following test driven development and have written the test first. In some cases where there is a great amount of uncertainty it is defendable to get a working concept first and then write a test.

addone-test.go
```Golang
package mathhelp

import "testing"

func TestAddone(t *testing.T) {
	for _, c := range []struct {
		in, want int
	}{
		{1, 2},
		{-5, -4},
		{100, 101},
	} {
		got := addone(c.in)
		if got != c.want {
			t.Errorf("addone(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```


### 3.	Trigger stages with discrete, scriptable steps

Nearly every Continuous Integration tool including Jenkins, etc operate by triggering a series of arbitrary tasks. Having these tasks independent and able to trigger locally will save an immense amount of time over having this configuration exist entirely in the tool such as Jenkins. Entering configuration in Jenkins and triggering the build to trouble shoot. Its 30 sec to run a shell script locally to 3-10 min of turn around waiting for Jenkins to spin up a worker and kick off the script.

build-bin.sh
```
#!/bin/bash

ROOT=$(dirname "${BASH_SOURCE}")/..
source "$ROOT/build/common.sh"

build_bin
```

run-tests.sh
```
#!/bin/bash

ROOT=$(dirname "${BASH_SOURCE}")/..
source "$ROOT/build/common.sh"

run_tests
```

common.sh
```
#!/bin/bash

# This will canonicalize the path
ROOT=$(cd $(dirname "${BASH_SOURCE}")/.. && pwd -P)
cd $ROOT

function run_tests() {
  go test
}

function build_bin() {
  go build
}
```

### 4.	Point preferred Continuous Integration tool to independent scripts and profit

NOTE: This requires the Golang plugin installed in Jenkins and configured in the Jenkins global tools section, not general configuration

Integrating Jenkins or any Continuous Integration tool is now straight forward. Additionally, the configuration now exists in source control for an added benefit instead of living in Jenkins or similar tool.

Adding additional steps is a matter of adding new shell scripts and wiring them into Jenkins. Perhaps one might build a docker container as a step and then deploy to Kubernetes as another. In the future multiple environments can also be shell scripts, perhaps sharing configuration variables so the build process can be clean and easy to follow. A great amount of complexity can be avoided with a little best practice.

Jenkinsfile
```
pipeline {
    agent any

    stages {
        stage('Build Bin') {
            steps {
                sh 'build/build-bin.sh'
            }
        }
        stage('Build War') {
            steps {
                sh 'build/run-tests.sh'
            }
        }
    }
}
```
