package app

import (
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/samanthajayasinghe/unit-test/testable/messenger/slack"
	"github.com/samanthajayasinghe/unit-test/testable/mocks"
)

func Test_Source(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Slack Test sute")
}

var _ = Context("Alack app testing:", func() {
	var (
		mockCtrl          *gomock.Controller
		mockMessageClient *mocks.MockMessageClientInterface
		mockLogger        *mocks.MockLoggerInterface
		channelName       string
		message           string
		response          slack.SlackMessageResponse
	)

	// Global setup for all Source tests
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockMessageClient = mocks.NewMockMessageClientInterface(mockCtrl)
		mockLogger = mocks.NewMockLoggerInterface(mockCtrl)
		channelName = "some-channel"
		message = "Hello World!"
		response = slack.SlackMessageResponse{X: "a", Y: "b", Z: "c"}
	})

	// Global cleanup for all Source tests
	AfterEach(func() {
		mockCtrl.Finish()
	})

	When("at lease one release exists", func() {
		It("App.Start shall have called service.Connect and service.SendMessage", func() {
			mockMessageClient.EXPECT().SendMessage(channelName, message).Return(response, nil)
			mockLogger.EXPECT().Infoln(gomock.Any()).AnyTimes()
			app := &App{
				MsgClient: mockMessageClient,
				Logger:    mockLogger,
			}
			app.SendMessage(channelName, message)
		})
	})

})
