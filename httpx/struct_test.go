package httpx

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

type Person struct {
	Name            string  `map:""`
	Age             int     `map:"Age,omitempty"`
	Gender          bool    `map:"gender"`
	Score           float64 `map:""`
	InvisibaleScore float64
}

func Test_StructHashmap(t *testing.T) {

	p := Person{
		Name:   "zhangsan",
		Age:    20,
		Gender: true,
		Score:  1123.31,
	}

	params := make(map[string]string)

	err := StructToHashmap(p, params)
	if err != nil {
		logrus.Fatalf("hashmap failed: %v", err)
	}

	// spew.Dump(params)
	// gomega.NewWithT()
	t.Run("StructHashMap", func(t *testing.T) {
		NewWithT(t).Expect(params["name"]).To(Equal("zhangsan"))
		NewWithT(t).Expect(params["Age"]).To(Equal("20"))
		NewWithT(t).Expect(params["gender"]).To(Equal("true"))
		NewWithT(t).Expect(params["score"]).To(Equal("1123.31"))
	})
}
