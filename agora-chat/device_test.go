package agora_chat

import (
	"errors"
	gohttp "net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ycj3/agora-chat-cli/http"
	"go.uber.org/mock/gomock"
)

var _ = Describe("DeviceManager", func() {
	var (
		mockCtrl      *gomock.Controller
		mockClient    *http.MockClient[deviceResponseResult]
		deviceManager *DeviceManager
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockClient = http.NewMockClient[deviceResponseResult](mockCtrl)
		client := &client{
			appConfig: &App{
				BaseURL: "https://api.example.com",
			},
			appToken:     "dummy-token",
			deviceClient: mockClient,
		}
		deviceManager = &DeviceManager{client: client}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("AddPushDevice", func() {
		It("should add a push device successfully", func() {
			mockResponse := http.Result[deviceResponseResult]{
				StatusCode: gohttp.StatusOK,
				Data: deviceResponseResult{
					Entities: []DeviceInfo{
						{DeviceID: "device1", DeviceToken: "token1", NotifierName: "APNS"},
					},
				},
			}

			mockClient.EXPECT().Send(gomock.Any()).Return(mockResponse, nil)

			devices, err := deviceManager.AddPushDevice("user1", "device1", "token1", "APNS")
			Expect(err).ToNot(HaveOccurred())
			Expect(devices).To(HaveLen(1))
			Expect(devices[0].DeviceID).To(Equal("device1"))
		})

		It("should return an error if the request fails", func() {
			mockClient.EXPECT().Send(gomock.Any()).Return(http.Result[deviceResponseResult]{}, errors.New("request failed"))

			devices, err := deviceManager.AddPushDevice("user1", "device1", "token1", "APNS")
			Expect(err).To(HaveOccurred())
			Expect(devices).To(BeNil())
		})
	})

	Describe("RemovePushDevice", func() {
		It("should remove a push device successfully", func() {
			mockResponse := http.Result[deviceResponseResult]{
				StatusCode: gohttp.StatusOK,
				Data: deviceResponseResult{
					Entities: []DeviceInfo{
						{DeviceID: "device1", DeviceToken: "", NotifierName: "APNS"},
					},
				},
			}

			mockClient.EXPECT().Send(gomock.Any()).Return(mockResponse, nil)

			devices, err := deviceManager.RemovePushDevice("user1", "device1", "APNS")
			Expect(err).ToNot(HaveOccurred())
			Expect(devices).To(HaveLen(1))
			Expect(devices[0].DeviceID).To(Equal("device1"))
			Expect(devices[0].DeviceToken).To(BeEmpty())
		})

		It("should return an error if the request fails", func() {
			mockClient.EXPECT().Send(gomock.Any()).Return(http.Result[deviceResponseResult]{}, errors.New("request failed"))

			devices, err := deviceManager.RemovePushDevice("user1", "device1", "APNS")
			Expect(err).To(HaveOccurred())
			Expect(devices).To(BeNil())
		})
	})

	Describe("ListPushDevice", func() {
		It("should list push devices successfully", func() {
			mockResponse := http.Result[deviceResponseResult]{
				StatusCode: gohttp.StatusOK,
				Data: deviceResponseResult{
					Entities: []DeviceInfo{
						{DeviceID: "device1", DeviceToken: "token1", NotifierName: "APNS"},
						{DeviceID: "device2", DeviceToken: "token2", NotifierName: "FCM"},
					},
				},
			}

			mockClient.EXPECT().Send(gomock.Any()).Return(mockResponse, nil)

			devices, err := deviceManager.ListPushDevice("user1")
			Expect(err).ToNot(HaveOccurred())
			Expect(devices).To(HaveLen(2))
			Expect(devices[0].DeviceID).To(Equal("device1"))
			Expect(devices[1].DeviceID).To(Equal("device2"))
		})

		It("should return an error if the request fails", func() {
			mockClient.EXPECT().Send(gomock.Any()).Return(http.Result[deviceResponseResult]{}, errors.New("request failed"))

			devices, err := deviceManager.ListPushDevice("user1")
			Expect(err).To(HaveOccurred())
			Expect(devices).To(BeNil())
		})
	})
})
