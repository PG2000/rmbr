package cmd_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
	"io/ioutil"
	"rmbr/cmd"
	mock_cmd "rmbr/cmd/mocks"
)

var _ = Describe("Create", func() {

	var (
		mockCtrl *gomock.Controller
		loader   *mock_cmd.MockRepository
		sut      *cobra.Command
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		loader = mock_cmd.NewMockRepository(mockCtrl)
		defer mockCtrl.Finish()

		manager := cmd.NewRmbrNotesManager(loader, fakeLogger())
		sut = cmd.CreateCommand(manager)
	})

	Context("Adds a note with fields command,description and group", func() {

		It("should have all fields", func() {
			// Arrange
			flags := []string{
				"--group",
				"kubernetes",
				"--description",
				"my super cool description",
				"--command",
				"my-command -d",
			}

			// Act
			parseFlagsErrors := sut.ParseFlags(flags)

			loader.EXPECT().New(cmd.RmbrNote{
				Command:     "my-command -d",
				Description: "my super cool description",
				Group:       "kubernetes",
			}).Times(1)

			err := sut.RunE(sut, flags)

			// Assert
			Expect(parseFlagsErrors).ToNot(HaveOccurred())
			Expect(err).NotTo(HaveOccurred())
		})

		It("should have an empty default group", func() {

			// Arrange
			flags := []string{
				"--description",
				"my super cool description",
				"--command",
				"my-command -d",
			}

			// Act
			parseFlagsErrors := sut.ParseFlags(flags)

			loader.EXPECT().New(cmd.RmbrNote{
				Command:     "my-command -d",
				Description: "my super cool description",
				Group:       "",
			}).Times(1)

			err := sut.RunE(sut, flags)

			// Assert
			Expect(parseFlagsErrors).ToNot(HaveOccurred())
			Expect(err).NotTo(HaveOccurred())
		})

		It("should fail when command is empty", func() {
			flags := []string{
				"--description",
				"",
				"--command",
				"",
				"--group",
				"",
			}

			sut.RunE(sut, flags)
		})
	})

})

func fakeLogger() *cmd.StandardLogger {
	logger := cmd.NewLogger()
	logger.SetOutput(ioutil.Discard)
	return logger
}
