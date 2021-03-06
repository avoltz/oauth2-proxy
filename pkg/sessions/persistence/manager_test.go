package persistence

import (
	"testing"
	"time"

	"github.com/oauth2-proxy/oauth2-proxy/pkg/apis/options"
	sessionsapi "github.com/oauth2-proxy/oauth2-proxy/pkg/apis/sessions"
	"github.com/oauth2-proxy/oauth2-proxy/pkg/logger"
	"github.com/oauth2-proxy/oauth2-proxy/pkg/sessions/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestManager(t *testing.T) {
	logger.SetOutput(GinkgoWriter)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Persistence Manager SessionStore")
}

var _ = Describe("Persistence Manager SessionStore Tests", func() {
	var ms *tests.MockStore
	BeforeEach(func() {
		ms = tests.NewMockStore()
	})
	tests.RunSessionStoreTests(
		func(_ *options.SessionOptions, cookieOpts *options.Cookie) (sessionsapi.SessionStore, error) {
			return NewManager(ms, cookieOpts), nil
		},
		func(d time.Duration) error {
			ms.FastForward(d)
			return nil
		})
})
