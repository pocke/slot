slot
========

n回でXXXが揃いました!を簡単に実装するためのライブラリ/コマンドラインツール。

Installation
-----------------

```sh
go get github.com/pocke/slot/...
```


Usage
-----

### As a Golang Library

```go
s := slot.New([]string{"h", "o", "g", "e"}, os.Stdout)
i := s.Start()
fmt.Println()
fmt.Printf("%d回でhogeが揃いました!\n", i)
```

### As a command

```sh
$ slot hoge
# => hegeeghegheggegggeehgohhhoeogeggehehoohgeegggghogohoghhoghoeggghgeghhheoogeoheohoge
#    83回で hoge が揃いました!
$ slot ho ge fu ga
# => hogahogegagafugahogefuga
#    12回で hogefuga が揃いました!
```
