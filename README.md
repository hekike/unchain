# conventional-commits

Inspired by: https://github.com/conventional-changelog

## Detect SemVer change since latest tag

```sh
$ bump /my-dir
major
```

## Parse commits since latest tag

```sh
$ parse /my-dir
hash,semver,type,component,description,body,footer
3d4428781580f92b0a56892978c5a22bc53903c9,major,fix,foo,fifth commit,body,BREAKING CHANGE: blabla
29afc9699602e73418395226f22389a5271c5e58,major,fix,bar,fourth commit,BREAKING CHANGE: blabla,
6289d27b800d3966ec7f14394ff4c48b08dd5976,patch,fix,foo,third commit,body,
998df6abedeeb0e090986b5de3a89e62c03c436d,patch,chore,foo,second commit,,
a4a95856d51dc3018170f2a854581590d1a27687,minor,feat,foo,initial commit,,
```
