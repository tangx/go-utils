package regex

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_NamedMatch(t *testing.T) {
	expr := `(?P<first>\d+)\.(\d+).(?P<second>\d+)`
	re := MustNew(expr)

	t.Run("Regex embedded Matching", func(t *testing.T) {
		ok := re.MatchString("1234.5678.9")
		NewWithT(t).Expect(ok).To(BeTrue())
	})

	t.Run("Regex Named Match String", func(t *testing.T) {
		str := re.FindStringNamedSubmatch("1234.5678.9").GetString("first")
		NewWithT(t).Expect(str).To(Equal("1234"))
	})

	t.Run("Regex Named Match Bytes", func(t *testing.T) {
		bytes := re.FindNamedSubmatch([]byte("1234.5678.9")).Get("first")
		NewWithT(t).Expect(bytes).To(Equal([]byte("1234")))
	})
}
