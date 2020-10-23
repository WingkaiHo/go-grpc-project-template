package e2e_test

import (
	"context"
	"fmt"
	"os"
	"strconv"

	pb "github.com/WingkaiHo/go-grpc-project-template/api/proto/demo"
	"github.com/WingkaiHo/go-grpc-project-template/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
)

var _ = Describe("Details service", func() {
	cfg := config.GetConfigByEnv()
	port := ":" + strconv.Itoa(cfg.Port)
	conn, _ := grpc.Dial(port, grpc.WithInsecure())

	client := pb.NewDetailsClient(conn)
	It("Should Details.Get be successful", func() {
		id := uint64(1)
		resp, err := client.Get(context.Background(), &pb.GetDetailRequest{
			Id: id,
		})

		host, _ := os.Hostname()
		name := fmt.Sprintf("%s-%d", host, id)

		Expect(err).To(BeNil())
		fmt.Println(resp, err)
		Expect(resp.Name).To(Equal(name))
	})

	AfterEach(func() {
		conn.Close()
	})
})
