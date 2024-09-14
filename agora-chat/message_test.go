package agora_chat

import (
	"errors"
	gohttp "net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ycj3/agora-chat-cli/http"
	"go.uber.org/mock/gomock"
)

var _ = Describe("MessageManager", func() {
	var (
		ctrl           *gomock.Controller
		mockClient     *http.MockClient[messageResponseResult]
		messageManager *MessageManager
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockClient = http.NewMockClient[messageResponseResult](ctrl)
		client := &client{
			messageClient: mockClient,
			appToken:      "dummy-token",
			appConfig:     &App{BaseURL: "https://api.example.com"},
		}
		messageManager = &MessageManager{client: client}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("SendUsersMessage", func() {
		var (
			message  *Message
			response http.Result[messageResponseResult]
		)

		BeforeEach(func() {
			message = &Message{
				From: "user1",
				To:   []string{"user2"},
				Type: MessageTypeText,
				Body: TextMessageBody{
					Msg: "Hello",
				},
			}
			response = http.Result[messageResponseResult]{
				StatusCode: gohttp.StatusOK,
				Data: messageResponseResult{
					Data: map[string]string{"message": "Message sent successfully"},
				},
			}
		})

		It("should send user message successfully", func() {
			mockClient.EXPECT().
				Send(gomock.Any()).
				Return(response, nil)

			result, err := messageManager.SendUsersMessage(message)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(map[string]string{"message": "Message sent successfully"}))
		})

		It("should return error if Send fails", func() {
			mockClient.EXPECT().
				Send(gomock.Any()).
				Return(http.Result[messageResponseResult]{}, errors.New("request failed"))

			result, err := messageManager.SendUsersMessage(message)
			Expect(err).To(HaveOccurred())
			Expect(result).To(BeNil())
			Expect(err.Error()).To(Equal("request failed: request failed"))
		})

		// Todo
		// It("should return error if response status is not OK", func() {
		// 	response.StatusCode = gohttp.StatusInternalServerError
		// 	response.Data = messageResponseResult{
		// 		Data: map[string]string{},
		// 	}
		// 	mockClient.EXPECT().
		// 		Send(gomock.Any()).
		// 		Return(response, nil)

		// 	result, err := messageManager.SendUsersMessage(from, to, msgType, body, options)
		// 	Expect(err).To(HaveOccurred())
		// 	Expect(result).To(BeNil())
		// })

	})
})
