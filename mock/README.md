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
---
/opt/homebrew/Cellar/go/1.18/libexec/bin/go tool test2json -t /private/var/folders/k_/d7b9dw9j1fn2zlg4h_qm_zqh99mxd5/T/GoLand/___TestSample_in_github_com_maru0804_Go_git_mock.test -test.v -test.paniconexit0 -test.run ^\QTestSample\E$
=== RUN   TestSample
--- PASS: TestSample (0.00s)
PASS

Process finished with the exit code 0
~~~
