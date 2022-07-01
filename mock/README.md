## mock test

## useage
0. 環境構築
~~~bash
go get github.com/golang/mock/gomock
// mock 環境の設定
go install github.com/golang/mock/mockgen
go get github.com/golang/mock/mockgen
~~~
1. mockファイル作成
~~~zsh
mockgen -source=main.go -destination mock/mock_main.go
~~~
2. mockを使ったテストの実装
3. 実行
~~~zsh
go test -v -run TestSample
~~~
