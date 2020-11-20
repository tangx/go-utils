package md5sha

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_md5sha(t *testing.T) {

	// NewWithT(t)
	t.Run("file md5sha", func(t *testing.T) {
		sha, _ := File("file.txt")
		NewWithT(t).Expect(sha).To(Equal("ba1f2511fc30423bdbb183fe33f3dd0f"))
	})

	t.Run("string md5sha", func(t *testing.T) {
		sha, _ := String("123\n")
		NewWithT(t).Expect(sha).To(Equal("ba1f2511fc30423bdbb183fe33f3dd0f"))
	})
}
