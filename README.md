# wiredify

transform Japanese va-vi-vu-va-vo and Japanese ba-bi-bu-be-bo mutually. 

## CLI

### wiredify

transform Japanese ba-bi-bu-be-bo to Japanese va-vi-vu-va-vo. 

#### usage

```shell
$ echo "ジェネレーティブ・エーアイ" | go run ./bin/dewiredify/main.go
ジェネレーティヴ・エーアイ
```

### dewiredify

transform Japanese va-vi-vu-va-vo to Japanese ba-bi-bu-be-bo

#### usage

```shell
$ echo "ジェネレーティヴ・エーアイ" | go run ./bin/dewiredify/main.go
ジェネレーティブ・エーアイ
$ echo "ポヴォ" | go run ./bin/dewiredify/main.go
ポボ
$ echo "ヴォーカル" | go run ./bin/dewiredify/main.go
ボーカル
$ echo "カヴァー" | go run ./bin/dewiredify/main.go
カバー
$ echo "ヴォイストレーニング" | go run ./bin/dewiredify/main.go
ボイストレーニング
```

## library

### install

```shell
go get -u "github.com/eniehack/wiredify/pkg/wiredify"
```

### wiredify

```go
import (
  "bufio"
  "bytes"

  "github.com/eniehack/wiredify/pkg/wiredify"
)

...

h := &wiredify.Handler{
  In:  bufio.NewScanner(os.Stdin),
  Out: new(bytes.Buffer),
}
h.Wiredify()
```

### dewiredify

```go
import (
  "bufio"
  "bytes"

  "github.com/eniehack/wiredify/pkg/wiredify"
)

...

h := &wiredify.Handler{
  In:  bufio.NewScanner(os.Stdin),
  Out: new(bytes.Buffer),
}
h.Dewiredify()
```


## other implementations

- [eggplants/wiredify](https://github.com/eggplants/wiredify) - an implementation written by Python.
- [oageo/wiredify](https://github.com/oageo/wiredify) - an implementation written by Rust 

