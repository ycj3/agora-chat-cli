package agora_chat

import (
	"errors"
	gohttp "net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ycj3/agora-chat-cli/http"
	"go.uber.org/mock/gomock"
)

var _ = Describe("PushManager", func() {
	var (
		ctrl        *gomock.Controller
		mockClient  *http.MockClient[PushResponseResult]
		pushManager *PushManager
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockClient = http.NewMockClient[PushResponseResult](ctrl)
		client := &client{
			pushClient: mockClient,
			appToken:   "dummy-token",
			appConfig:  &App{BaseURL: "https://api.example.com"},
		}
		pushManager = &PushManager{client: client}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("SyncPush", func() {
		var (
			userID   string
			strategy pushStrategy
			msg      PushMessage
			response http.Result[PushResponseResult]
		)

		BeforeEach(func() {
			userID = "user123"
			strategy = OnlyPushPrivider
			msg = PushMessage{
				Title:    "Test Title",
				Content:  "Test Content",
				SubTitle: "Test SubTitle",
			}
			response = http.Result[PushResponseResult]{
				StatusCode: gohttp.StatusOK,
				Data: PushResponseResult{
					Data: []PushResult{
						{
							PushStatus: "Success",
							Desc:       "Push sent successfully",
							StatusCode: gohttp.StatusOK,
						},
					},
				},
			}
		})

		It("should sync push successfully", func() {
			mockClient.EXPECT().
				Send(gomock.Any()).
				Return(response, nil)

			result, err := pushManager.SyncPush(userID, strategy, msg)
			Expect(err).To(BeNil())
			Expect(result.Data[0].PushStatus).To(Equal("Success"))
			Expect(result.Data[0].Desc).To(Equal("Push sent successfully"))
		})

		It("should return error if Send fails", func() {
			mockClient.EXPECT().
				Send(gomock.Any()).
				Return(http.Result[PushResponseResult]{}, errors.New("request failed"))

			result, err := pushManager.SyncPush(userID, strategy, msg)
			Expect(err).To(HaveOccurred())
			Expect(result.Data).To(BeNil())
			Expect(err.Error()).To(Equal("request failed: request failed"))
		})

		//Todo
		// It("should return error if response status is not OK", func() {
		// 	response.StatusCode = gohttp.StatusInternalServerError
		// 	response.Data = PushResponseResult{
		// 		Data: []pushResult{},
		// 	}
		// 	mockClient.EXPECT().
		// 		Send(gomock.Any()).
		// 		Return(response, nil)

		// 	result, err := pushManager.SyncPush(userID, strategy, msg)
		// 	Expect(err).To(HaveOccurred())
		// 	Expect(result.Data).To(BeNil())
		// })
	})
})
