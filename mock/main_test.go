package main

import (
	"github.com/golang/mock/gomock"
	mock_main "github.com/maru0804/Go.git/mock/mock"
	"testing"
)

func TestSample(t *testing.T) {
	// mockのコントローラを作成します
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ApiClientインターフェイスのmockを作成します
	mockApiClinet := mock_main.NewMockApiClient(ctrl)
	// 作成したmockに対して期待する呼び出しと返り値を定義します
	// EXPECT()では呼び出されたかどうか
	// Request()ではそのメソッド名が指定した引数で呼び出されたかどうか
	// Return()では返り値を指定します
	mockApiClinet.EXPECT().Request("bar").Return("bar", nil)

	d := &DataRegister{}
	d.client = mockApiClinet // mockを登録
	expected := "bar"

	res, err := d.Register(expected)
	if err != nil {
		t.Fatal("Register error!", err)
	}
	if res != expected {
		t.Fatal("Value does not match.")
	}
}
