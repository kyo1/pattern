# pattern

## Installation

```sh
$ go get github.com/kyo1/pattern
```

## How to use

```sh
$ pattern create 32
aabaacaadaaeaafaagaahaaiaajaakaa

$ pattern offset acaa
acaa found at offset 4

$ pattern offset 0x61616361
acaa found at offset 4

$ pattern offset -b 0x61616361
aaca found at offset 3
```

