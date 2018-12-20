# unchain ⛓️

[![CircleCI](https://circleci.com/gh/hekike/unchain.svg?style=svg&circle-token=c7e325c818a4865e9660d944d97fb5fae4b37043)](https://circleci.com/gh/hekike/unchain)

Release tool with automatic changelog generation and next SemVer version calculation based on conventional commits.

## Install

```sh
go get github.com/hekike/unchain
```

## Usage

See help for available commands and parameters.

```sh
unchain -h
```

## Commands

### release

Run in your terminal:

```sh
unchain release
```

- `dir` optional *(default: workdir)*, directory

#### How It Works

Automatically detects the last version and bumps the `patch`, `minor` or `major` semver component based on the conventional commits since that release.
If there is no commit found related to previous version it will release `1.0.0`.

*What It Does*

* Detects the next SemVer version based on commit history
* Detects the previous version from release commits made by this tool or from package.json if any
* Creates or prepends `CHANGELOG.md`
* *(optional)* Execs `npm version` if finds package.json
* Git tags release
* *(optional)* `npm publish` if finds package.json
* Runs `git push` to sync with remote

*CHANGELOG.md example*

```
<a name="2.0.0"></a>
## 2.0.0 (2018-12-16)


#### Bug Fixes

* **http:** fix DNS resolution 6289d27b800d3966ec7f14394ff4c48b08dd5976
* **http:** fix port parsing 998df6abedeeb0e090986b5de3a89e62c03c436d

#### Features

* **http:** add TLS support a4a95856d51dc3018170f2a854581590d1a27687

#### Breaking Changes

* make host config required ecd94da5b9f10c04ce53723729ae7068cc73557e
* make port config required 29afc9699602e73418395226f22389a5271c5e58

```

*Commits example*

- (optional, npm only): chore(package): bump version to 1.0.0
- (always): chore(changelog): update for version 1.0.0

*Tag created*

- `1.0.0` (with package.json, v1.0.0)

Skips non API facing commits from the changelog like `test`, `chore` and `refactor`.

## semver

Detect SemVer change since latest Git Tag.

```sh
$ unchain semver
major
```

## parse

Parse commits since latest Git Tag.

```sh
$ unchain parse
hash,semver,type,component,description,body,footer
ecd94da5b9f10c04ce53723729ae7068cc73557e,major,fix,foo,fifth commit,body,BREAKING CHANGE: so breaking
29afc9699602e73418395226f22389a5271c5e58,major,fix,bar,fourth commit,BREAKING CHANGE: blabla,
6289d27b800d3966ec7f14394ff4c48b08dd5976,patch,fix,foo,third commit,body,
998df6abedeeb0e090986b5de3a89e62c03c436d,patch,chore,foo,second commit,,
a4a95856d51dc3018170f2a854581590d1a27687,minor,feat,foo,initial commit,,
```

## Background

Inspired by:

- https://github.com/conventional-changelog
- https://github.com/Unleash/unleash

Follows:

- https://semver.org
- https://www.conventionalcommits.org
