package httpx

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

type Person struct {
	Name            string  `map:"name"`
	Age             int     `map:"Age,omitempty"`
	Gender          bool    `map:"gender"`
	Score           float64 `map:""`
	InvisibaleScore float64
	Addr            Address `map:""`
}

type Address struct {
	Name   string `map:"name"`
	Block  string `map:",inline"`
	Street string `map:"street"`
}

func Test_StructHashmap(t *testing.T) {

	p := Person{
		Name:   "zhangsan",
		Age:    20,
		Gender: true,
		Score:  1123.31,
		Addr: Address{
			Name:   "home",
			Block:  "building3",
			Street: "xiangyang",
		},
	}

	params := make(map[string]string)

	err := StructToHashmap(p, params)
	if err != nil {
		logrus.Fatalf("hashmap failed: %v", err)
	}

	// spew.Dump(params)
	/*
		(map[string]string) (len=7) {
		 (string) (len=5) "Score": (string) (len=7) "1123.31",
		 (string) (len=10) "Addr__name": (string) (len=4) "home",
		 (string) (len=5) "Block": (string) (len=9) "building3",
		 (string) (len=12) "Addr__street": (string) (len=9) "xiangyang",
		 (string) (len=4) "name": (string) (len=8) "zhangsan",
		 (string) (len=3) "Age": (string) (len=2) "20",
		 (string) (len=6) "gender": (string) (len=4) "true"
		}
	*/
	t.Run("StructHashMap", func(t *testing.T) {
		NewWithT(t).Expect(params["name"]).To(Equal("zhangsan"))
		NewWithT(t).Expect(params["Age"]).To(Equal("20"))
		NewWithT(t).Expect(params["gender"]).To(Equal("true"))
		NewWithT(t).Expect(params["Score"]).To(Equal("1123.31"))

		NewWithT(t).Expect(params["Addr__name"]).To(Equal("home"))
		NewWithT(t).Expect(params["Addr__street"]).To(Equal("xiangyang"))
		NewWithT(t).Expect(params["Block"]).To(Equal("building3")) // 使用 inline 内联了到最上层
	})
}
