# CONCOURSE CHEAT SHEET

`concourse` _is a pipelined CI (Continuous Integration) for
Software Builds._

Concourse's main goal is to run tasks.

[GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/)

## INSTALL AND RUN CONCOURSE ON VIRTUALBOX

Will run concourse on VirtualBox using vagrant.

Get vargrant for windows

https://www.vagrantup.com

Get the concourse vagrant files. 

To run, from windows cmd prompt type,

```bash
vagrant up
```

To upgrade (will destroy current pipelines).

```bash
vagrant box update --box concourse/lite     # gets the newest Vagrant box
vagrant destroy                             # remove the old Vagrant box
vagrant up 
```

Reinstal the new fly.exe for windows and linux.


## INSTALL FLY ON WINDOWS

Open concourse,

http://192.168.100.4:8080/

In bottom right download latest verions of fly.exe for linux,

Place in,

```bash
C:\Program Files (x86)\Concourse\fly.exe
```

Login fly to concourse,

```bash
"C:\Program Files (x86)\Concourse\fly.exe" -t ci login -c http://192.168.100.4:8080/
```

Check version,

```bash
"C:\Program Files (x86)\Concourse\fly.exe" -version
```

## INSTALL FLY ON LINUX

Open concourse,

http://192.168.100.4:8080/

In bottom right download latest verions of fly.exe for linux,

Place in,

```bash
mkdir $HOME/bin
install $HOME/Downloads/fly $HOME/bin
```

Login fly to concourse,

```bash
fly -t ci login -c http:// 192.168.100.4:8080/
```

Check version,

```bash
fly -version
```

## SIMPLE TEST CASE

Concourse job is to run tasks.  So lets run a task that
does a search and replace on a file.

There are 3 main files:

1. scripts/search_and_replace.sh

```bash
#!/bin/bash
# my-go-examples readme-github-pages.sh

ls -lat
```

2. tasks/task-search-and-replace.yml

```bash
# my-go-examples task-readme-github-pages.yml

platform: linux

image_resource:
  type: docker-image
  source:
    repository: golang
    tag: 1.7.1

inputs:
- name: my-go-examples

outputs:
- name: my-go-examples-updated

run:
path: ./my-go-examples/ci/scripts/readme-github-pages.sh
```

3. pipeline.yml

```bash
# my-go-examples pipeline.yml

jobs:

- name: job-readme-github-pages
  plan:
  - get: my-go-examples
    trigger: true
  - task: task-readme-github-pages
    file: my-go-examples/ci/tasks/task-readme-github-pages.yml
  - put: my-go-examples
    params:
      repository: my-go-examples-updated

- name: job-unit-tests
  plan:
  - get: my-go-examples
    trigger: true
  - put: resource-slack-alert
    params:
      channel: '#jeff-builds'
      text: "From my-go-examples: STARTED job-unit-tests in concourse ci."
  - put: resource-repo-status
    params: { state: "pending", description: "STARTED job-unit-tests in concourse ci.", commit: "my-go-examples" }
  - task: task-unit-tests
    file: my-go-examples/ci/tasks/task-unit-tests.yml
    on_success:
      do:
      - put: my-go-examples
        params:
          repository: my-go-examples  
      - put: resource-slack-alert
        params:
          channel: '#jeff-builds'
          text_file: coverage-results/test_coverage.txt
          text: |
            From my-go-examples: PASSED job-unit-tests in concourse ci. 
            $TEXT_FILE_CONTENT
      - put: resource-repo-status
        params: { state: "success", description: "PASSED job-unit-tests in concourse ci", commit: "my-go-examples" }
    on_failure:
      do:
      - put: resource-slack-alert
        params:
          channel: '#jeff-builds'
          text: "From my-go-examples: FAILED job-unit-tests in concourse ci."
      - put: resource-repo-status
        params: { state: "failure", description: "FAILED job-unit-tests in concourse ci.", commit: "my-go-examples" }

resource_types:

- name: slack-notification
  type: docker-image
  source:
    repository: cfcommunity/slack-notification-resource
    tag: latest
- name: github-status
  type: docker-image
  source:
    repository: dpb587/github-status-resource
    tag: master

resources:

- name: my-go-examples
  type: git
  source:
    #uri: https://github.com/jeffdecola/my-go-examples
    uri: git@github.com:jeffdecola/my-go-examples.git
    branch: master
    private_key: {{git_private_key}}
- name: resource-slack-alert
  type: slack-notification
  source:
    url: {{slack_url}} 
- name: resource-repo-status 
  type: github-status
  source:
    repository: jeffdecola/my-go-examples 
access_token: {{repo_github_token}}

```


