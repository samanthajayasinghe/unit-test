package slack

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test_Source(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Slack client Test Suie")
}

var _ = Context("Slack client testing:", func() {
	var (
		client      SlackClient
		channelName string
		message     string
	)

	// Global setup for all Source tests
	BeforeEach(func() {
		client = SlackClient{Key: "test-key"}
		channelName = "#some-channel"
		message = "Hello World!"
	})

	// Global cleanup for all Source tests
	AfterEach(func() {

	})

	When("slack sending message to the channel", func() {

		It("Should fail for empty keys", func() {

			client = SlackClient{}
			res, err := client.SendMessage(channelName, message)

			//check err after sending the message
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).Should(ContainSubstring("key is empty"))

			Expect(res).To(BeNil())

		})

		It("Should fail for empty channel", func() {

			res, err := client.SendMessage("", message)

			//check err after sending the message
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).Should(ContainSubstring("channel name is empty"))

			Expect(res).To(BeNil())

		})

		It("Should fail for empty message", func() {
			res, err := client.SendMessage(channelName, "")

			//check err after sending the message
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).Should(ContainSubstring("message is empty"))

			Expect(res).To(BeNil())
		})
	})

})
