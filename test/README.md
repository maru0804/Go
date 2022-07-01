# ginkgo
# install
~~~zsh
go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/gomega/...
~~~
# usage
準備
~~~zsh
#テストスイートファイルの作成
ginkgo bootstrap

#テストファイルの作成
ginkgo generate Person(file name)
~~~
実行
~~~zsh
#test実行
 ginkgo -v -cover -coverprofile=./cover.out
~~~
success result
~~~
Random Seed: 1656397051

Will run 0 of 0 specs

Ran 0 of 0 Specs in 0.000 seconds
SUCCESS! -- 0 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS
coverage: [no statements]
composite coverage: [no statements]

Ginkgo ran 1 suite in 3.03224875s
Test Suite Passed
~~~
# rink
- [public](https://onsi.github.io/ginkgo/)
- [tutorial](https://qiita.com/myoshimi/items/62bc89b8065e08834b02)
