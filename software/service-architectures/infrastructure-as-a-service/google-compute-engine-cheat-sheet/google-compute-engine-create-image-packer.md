# CREATE A CUSTOM IMAGE USING PACKER CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Create a custom image using packer is a way
to create a custom image on gce._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-architectures/infrastructure-as-a-service/google-compute-engine-cheat-sheet/google-compute-engine-create-image-packer.md#overview)
* [GIVE PACKER AUTHENTICATION](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-architectures/infrastructure-as-a-service/google-compute-engine-cheat-sheet/google-compute-engine-create-image-packer.md#give-packer-authentication)
* [BASIC GCE PACKER TEMPLATE FILE (A GOOD PLACE TO START)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/service-architectures/infrastructure-as-a-service/google-compute-engine-cheat-sheet/google-compute-engine-create-image-packer.md#basic-gce-packer-template-file-a-good-place-to-start)

Documentation and Reference

* [hello-go-deploy-gce](https://github.com/JeffDeCola/hello-go-deploy-gce)
* [gce template file reference](https://www.packer.io/docs/builders/googlecompute.html)
* [packer](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations/orchestration/builds-deployment-containers/packer-cheat-sheet)

## OVERVIEW

Packer's entire deal is to create custom machine `images`.
This is exactly what we want to do.

The following illustration shows how `packer` controls the automation of
building an `image`.  As you can see, it all stems from one
configuration file `gce-packer-template.json`.

![IMAGE -  google compute engine create custom image packer - IMAGE](../../../../docs/pics/software/service-architectures/gce-create-custom-image-packer.svg)

## GIVE PACKER AUTHENTICATION

Packer needs to be authorized to use your `gce` account.
This is done using a google service account file. We already setup an env
variable `$GCP_JEFFS_SERVICE_ACCOUNT_PATH` that points to the
location of the service account file.

For information how to set this up checkout my cheat sheet
[here](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/service-providers/google-cloud-platform-cheat-sheet#identity--access-management-iam---service-accounts).

## BASIC GCE PACKER TEMPLATE FILE (A GOOD PLACE TO START)

A good place to start is this template file.

This build will,

* Authorize Packer with `gce`.
  * Verify the credential via `$GCP_JEFFS_SERVICE_ACCOUNT_PATH` env variable.
  * Use Project `$GCP_JEFFS_PROJECT_ID` env variable.
* Starts up a temporary VM `instance`  and temporary `boot disk`
  from a source `image`.
  * Use source gce machine `image` `ubuntu-1604-xenial-v20190306`.
* Provisions (Configures and Installs) whatever you want on
  this VM `instance`.
  * Copy a welcome file to /tmp.
  * Add a user named `jeff` via `useradd` command.
    This is done in the script `add-user-jeff.sh`
  * Move the welcome file from /tmp to /home/jeff.
    this is done in script `move-welcome-file.sh`.
* Deletes the temporary VM `instance`.
* Creates the custom `image` based on `boot disk`.
* Deletes the `boot disk`.

The packer command would be,

```bash
packer build -force \
    -var "account_file=$GCP_JEFFS_SERVICE_ACCOUNT_PATH" \
    -var "project_id=$GCP_JEFFS_PROJECT_ID" \
    gce-packer-template.json
```

And the template file `gce-packer-template.json` would be,

```json
{
    "variables": {
        "account_file": "",
        "project_id": "",
        "source_image": "ubuntu-1604-xenial-v20190306",
        "image_name": "hello-go-{{isotime \"20060102\"}}",
        "ssh_username": "packer",
        "zone": "us-west1-a"
    },

    "builders": [
        {
            "type": "googlecompute",
            "ssh_timeout": "10m",
            "account_file":"{{user `account_file`}}",
            "project_id": "{{user `project_id`}}",
            "source_image": "{{user `source_image`}}",
            "image_name": "{{user `image_name`}}",
            "ssh_username": "{{user `ssh_username`}}",
            "zone": "{{user `zone`}}"
        }
    ],

    "provisioners": [
        {
            "type": "file",
            "source": "./install-scripts/welcome.txt",
            "destination": "/tmp/welcome.txt"
        },
        {
            "type": "shell",
            "pause_before": "10s",
            "execute_command": "chmod +x {{ .Path }}; {{ .Vars }} sudo -E {{ .Path }}",
            "scripts": [
                "./install-scripts/add-user-jeff.sh",
                "./install-scripts/move-welcome-file.sh"
            ]
        }
    ]
}
```

Some quicks notes,

The ssh_username is `packer`.  That's how packer can ssh
into your temporary VM instance.

The `execute_command` `sudo -E` indicates to the security
policy that the user wishes to preserve their existing
environment variables.

To see this template working in a real example, go to my repo
[hello-go-deploy-gce](https://github.com/JeffDeCola/hello-go-deploy-gce).
