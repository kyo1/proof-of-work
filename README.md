# Proof-of-Work

## Installation

```sh
go get -u github.com/kyo1/proof-of-work
```

## Usage

```sh
$ alias pow='proof-of-work'

$ pow -a md5 --output-prefix=000
bHC

$ pow -a md5 --output-prefix=000 --dump
md5(bHC) = 0001978d2e6c3473f66d8d58c0c20bc3

$ pow -a sha512 --input-prefix=aaa --output-prefix=00000 --dump
sha512(aaadBbD) = 00000a65fe21ac6e7be245d92ac1ca4c7eb03ea4fc3f8f6d3eecc4089c13a3d78bf9343391832f6731401a63c2c91c05ac4104af737bd288536a5627f4499877
```
