// +build acceptance

package util
import (
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

type Space struct {
	page *agouti.Page
}

func (s Space) ViewApp(appName string) App {
	DelayForRendering()
	Expect(s.page.FindByLink(appName)).To(BeFound())
	Eventually(Expect(s.page.FindByLink(appName).Click()).To(Succeed()))
	return App{s.page}
}
