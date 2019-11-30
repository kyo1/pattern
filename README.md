# pattern

## Installation

```sh
$ go get -u github.com/kyo1/pattern/cmd/pattern
```

## How to use

```sh
$ pattern create 32
aaaabaaacaaadaaaeaaafaaagaaahaaa

$ pattern offset baaa
acaa found at offset 4

$ pattern offset 0x61616162
baaa found at offset 4

$ pattern offset -b 0x61616162
aaab found at offset 1
```

