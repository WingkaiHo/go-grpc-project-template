package e2e_test

import (
	"testing"
	"time"

	"github.com/WingkaiHo/go-grpc-project-template/app"

	"github.com/WingkaiHo/go-grpc-project-template/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func() {
	cfg := config.GetConfigByEnv()
	go app.StartApp(cfg)
	time.Sleep(5 * time.Second)
})

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}
