
# datahasher.ComputeHash(interface{}) uint64 

- I borrowed [the `ComputeHash` function](https://github.com/vugu/vugu/blob/495882447160f3d5a38ffa6653f98a07881baba5/data-hasher.go) from [Brad Peabody's](https://peabody.io) golang wasm project [Vugu on github](https://github.com/vugu/vugu).  It has the MIT License.


# Usage

```xxhashVal := datahasher.ComputeHash(anyThing)```

# Tests

Run tests like this: ```go test -v```

Tests are written using GoConvey https://github.com/smartystreets/goconvey

