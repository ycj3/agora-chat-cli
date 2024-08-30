package agora_chat

import (
	"errors"
	gohttp "net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ycj3/agora-chat-cli/http"
	"go.uber.org/mock/gomock"
)

var _ = Describe("UserManager", func() {
	var (
		ctrl        *gomock.Controller
		mockClient  *http.MockClient[userResponseResult]
		userManager *UserManager
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockClient = http.NewMockClient[userResponseResult](ctrl)
		client := &client{
			userClient: mockClient,
			appToken:   "dummy-token",
			appConfig:  &App{BaseURL: "https://api.example.com"},
		}
		userManager = &UserManager{client: client}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("UserOnlineStatuses", func() {
		var (
			userIDs  []string
			response http.Result[userResponseResult]
		)

		BeforeEach(func() {
			userIDs = []string{"user123", "user456"}
			response = http.Result[userResponseResult]{
				StatusCode: gohttp.StatusOK,
				Data: userResponseResult{
					Data: []userOnlineStatus{
						{"user123": "online"},
						{"user456": "offline"},
					},
				},
			}
		})

		It("should retrieve user online statuses successfully", func() {
			mockClient.EXPECT().
				Send(gomock.Any()).
				Return(response, nil)

			statuses, err := userManager.UserOnlineStatuses(userIDs)
			Expect(err).To(BeNil())
			Expect(statuses).To(HaveLen(2))
			Expect(statuses[0]["user123"]).To(Equal("online"))
			Expect(statuses[1]["user456"]).To(Equal("offline"))
		})

		It("should return error if Send fails", func() {
			mockClient.EXPECT().
				Send(gomock.Any()).
				Return(http.Result[userResponseResult]{}, errors.New("request failed"))

			statuses, err := userManager.UserOnlineStatuses(userIDs)
			Expect(err).To(HaveOccurred())
			Expect(statuses).To(BeNil())
			Expect(err.Error()).To(Equal("request failed: request failed"))
		})
		// Todo
		// It("should return error if response status is not OK", func() {
		// 	response.StatusCode = gohttp.StatusInternalServerError
		// 	response.Data = userResponseResult{
		// 		Error: Error{Exception: "500", ErrorDescription: "Internal Server Error"},
		// 	}
		// 	mockClient.EXPECT().
		// 		Send(gomock.Any()).
		// 		Return(response, nil)

		// 	statuses, err := userManager.UserOnlineStatuses(userIDs)
		// 	Expect(err).To(HaveOccurred())
		// 	Expect(statuses).To(BeNil())
		// 	Expect(err.Error()).To(Equal("Internal Server Error"))
		// })
	})
})
