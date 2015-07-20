slot
========

n回でXXXが揃いました!を簡単に実装するためのライブラリ。

Installation
-----------------

```sh
go get github.com/pocke/slot/...
```


Usage
-----

### As a Golang Library

```go
s := New([]string{"h", "o", "g", "e"}, os.Stdout)
i := s.Start()
fmt.Println()
fmt.Printf("%d回でhogeが揃いました!\n", i)
```

### As a command

```sh
$ slot hoge
# => hegeeghegheggegggeehgohhhoeogeggehehoohgeegggghogohoghhoghoeggghgeghhheoogeoheohoge
#    83回で hoge が揃いました!
```
