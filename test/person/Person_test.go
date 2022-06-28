package person_test

import (
	"errors"
	"github.com/maru0804/Go.git/test/person"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Person", func() {
	Context("Test for Name", func() { // Nameに関するテストです．文字列は変更できます
		p := person.Person{}
		/* Person.SetNameメソッドのテスト */
		It("Test for SetName", func() {
			p.SetName("Alice") // SetNameメソッドを実行します
		})
		/* Person.GetNameメソッドのテスト(正常系) */
		It("Test for GetName (Normal)", func() {
			p := person.Person{}
			p.SetName("Bob")              // SetNameメソッドでnameにBobをセットします
			name, err := p.GetName()      // GetNameメソッドでnameを読み出します
			Expect(err).To(BeNil())       // エラーは発生しない(nil)ことが期待されます
			Expect(name).To(Equal("Bob")) // nameはBobだと期待されます
		})
		/* Person.GetNameメソッドのテスト(エラー系) */
		It("Test for GetName (Error)", func() {
			p := person.Person{}
			_, err := p.GetName() // エラーなはずなので，nameは使用しません
			// nameが設定されていないので，エラーが返る
			Expect(err).NotTo(BeNil())                           // エラーが発生します(err != nil)
			Expect(err).To(Equal(errors.New("name is not set"))) // エラー内容が正しいか確認します
		})
	})

})
