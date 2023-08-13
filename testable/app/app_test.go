package app

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/samanthajayasinghe/unit-test/testable/messenger/slack"
	"github.com/samanthajayasinghe/unit-test/testable/mocks"
)

func Test_Source(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App Test Suie")
}

var _ = Context("Messenger App testing:", func() {
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
		channelName = "#some-channel"
		message = "Hello World!"
		response = slack.SlackMessageResponse{X: "a", Y: "b", Z: "c"}
	})

	// Global cleanup for all Source tests
	AfterEach(func() {
		mockCtrl.Finish()
	})

	When("app sending messages to the channel", func() {
		It("Send message successfully", func() {
			mockMessageClient.EXPECT().SendMessage(channelName, message).Return(response, nil)
			mockLogger.EXPECT().Infoln(gomock.Any()).AnyTimes()
			app := &App{
				MsgClient: mockMessageClient,
				Logger:    mockLogger,
			}
			err := app.SendMessage(channelName, message)

			//check err after sending the message
			Expect(err).To(BeNil())

		})

		It("Failed to send messages", func() {
			mockMessageClient.EXPECT().SendMessage(channelName, message).Return(response, fmt.Errorf("scope_unmatched"))
			mockLogger.EXPECT().Errorln(gomock.Any()).AnyTimes()
			app := &App{
				MsgClient: mockMessageClient,
				Logger:    mockLogger,
			}
			err := app.SendMessage(channelName, message)

			//check the error message
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).Should(ContainSubstring("fail to send message:"))

		})
	})

})
