# carrier

Opinionated platform that runs on Kubernetes, that takes you from App to URL in one step.

![CI](https://github.com/SUSE/carrier/workflows/CI/badge.svg)

<img src="./docs/carrier.svg" width="50%" height="50%">

## Principles

- must fit in less than 4GB of RAM
- must install in less than 5 minutes when images are warm
- must install with a one-line command and zero config
- must completely uninstall and leave the cluster in its previous state with a one-line command
- must work on local clusters (edge friendly)

### Guidelines

- if possible, choose components that are written in go
- all acceptance tests should run in less than 10 minutes
- all tests should be able to run on the minimal cluster 

## Usage
### Install

```bash
$ carrier install
```
### Uninstall

```bash
$ carrier uninstall
```
---
### Push an app

Run the following command inside of any supported app directory (e.g. inside [sample-app directory](sample-app)).

```bash
$ carrier push NAME
```

### Delete an app

```bash
$ carrier delete NAME
```
---
### Create a separate org

```bash
$ carrier create-org NAME
```

### Target an org

```bash
$ carrier target NAME
```
---
### List all commands

```bash
$ carrier help
```
### Detailed help for each command

```bash
$ carrier COMMAND --help
```