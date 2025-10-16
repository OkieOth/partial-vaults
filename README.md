![WIP](https://img.shields.io/badge/work%20in%20progress-red)

# partial-vaults

A tool to encrypt or decrypt JSON or YAML files that contain partial Ansible vault encrypted values.

## Overview

partial-vaults allows you to work with files that have a mix of plain text and Ansible Vault encrypted values. Unlike standard Ansible Vault which encrypts entire files, this tool lets you:

- Encrypt specific values within JSON or YAML files
- Decrypt specific values while leaving others encrypted
- Process nested structures with selective encryption


## Usage


## Build

You can build this tool from source as native build or as docker image

```bash
# build binary
# Requirements: installed golang 1.24, make
make build

# build docker image
# Requirements: installed docker
make build-docker
```



# Additional

* Passwort for `resources/tests/partial.yaml`: aa
* Passwort for `resources/tests/partial_encrypted_example.json`: test999
* Passwort for `resources/tests/partial_encrypted_example.yaml`: test999
