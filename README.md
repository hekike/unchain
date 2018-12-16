# conventional-commits

Inspired by: https://github.com/conventional-changelog

## Generate changelog

```sh
$ changelog /my-dir
<a name="2.0.0"></a>
## 2.0.0 (2018-12-16)


#### Bug Fixes

* **foo:** third commit 6289d27b800d3966ec7f14394ff4c48b08dd5976
* **foo:** second commit 998df6abedeeb0e090986b5de3a89e62c03c436d

#### Features

* **foo:** initial commit a4a95856d51dc3018170f2a854581590d1a27687

#### Breaking Changes

* so braking, much pain ecd94da5b9f10c04ce53723729ae7068cc73557e
* blabla 29afc9699602e73418395226f22389a5271c5e58

```

## Detect SemVer change since latest tag

```sh
$ bump /my-dir
major
```

## Parse commits since latest tag

```sh
$ parse /my-dir
hash,semver,type,component,description,body,footer
ecd94da5b9f10c04ce53723729ae7068cc73557e,major,fix,foo,fifth commit,body,BREAKING CHANGE: so braking, much pain
29afc9699602e73418395226f22389a5271c5e58,major,fix,bar,fourth commit,BREAKING CHANGE: blabla,
6289d27b800d3966ec7f14394ff4c48b08dd5976,patch,fix,foo,third commit,body,
998df6abedeeb0e090986b5de3a89e62c03c436d,patch,chore,foo,second commit,,
a4a95856d51dc3018170f2a854581590d1a27687,minor,feat,foo,initial commit,,
```
