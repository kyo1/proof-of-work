# Proof-of-Work

## Installation

```sh
go get -u github.com/kyo1/proof-of-work
```

## Usage

```sh
$ alias pow='proof-of-work'

$ pow -a md5 --output-prefix=000
2Da

$ pow -a md5 --output-prefix=000 --dump
md5(2Da) = 000a7e887bdb8f67cc3f82c1e177dbb2

$ pow -a sha512 --input-prefix=aaa --output-prefix=00000 --dump
sha512(aaa3UVd) = 00000b26b6ee740f12be4d540c4fe39a419d5320088c1b70e3ee3b0d753c91d170959daa2e7ab6222f5b7c8df11bdd6dc440dd8e7277f57bf60b705dac25af9a
```
