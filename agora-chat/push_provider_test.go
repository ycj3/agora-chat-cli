package agora_chat

import (
	"errors"
	gohttp "net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ycj3/agora-chat-cli/http"
	"go.uber.org/mock/gomock"
)

var _ = Describe("ProviderManager", func() {
	var (
		mockCtrl        *gomock.Controller
		mockClient      *http.MockClient[PrividerResponseResult]
		providerManager *ProviderManager
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockClient = http.NewMockClient[PrividerResponseResult](mockCtrl)
		client := &client{
			appConfig: &App{
				BaseURL: "https://api.example.com",
			},
			appToken:       "dummy-token",
			providerClient: mockClient,
		}
		providerManager = &ProviderManager{client: client}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("InsertPushProvider", func() {
		It("should insert a push provider successfully", func() {
			provider := PushProvider{
				Type: PushProviderAPNS,
				Name: "TestProvider",
				Env:  EnvDevelopment,
				ApnsPushSettings: &APNSConfig{
					TeamId: "team-id",
					KeyId:  "key-id",
				},
			}

			mockResponse := http.Result[PrividerResponseResult]{
				StatusCode: gohttp.StatusOK,
				Data: PrividerResponseResult{
					Entities: []PushProvider{
						provider,
					},
				},
			}

			mockClient.EXPECT().Send(gomock.Any()).Return(mockResponse, nil)

			result, err := providerManager.UpsertPushProvider(provider)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Entities).To(HaveLen(1))
			Expect(result.Entities[0].Name).To(Equal("TestProvider"))
		})

		It("should return an error if the request fails", func() {
			provider := PushProvider{
				Type: PushProviderAPNS,
				Name: "TestProvider",
			}

			mockClient.EXPECT().Send(gomock.Any()).Return(http.Result[PrividerResponseResult]{}, errors.New("request failed"))

			result, err := providerManager.UpsertPushProvider(provider)
			Expect(err).To(HaveOccurred())
			Expect(result.Entities).To(BeNil())
		})
	})

	Describe("DeletePushProvider", func() {
		It("should delete a push provider successfully", func() {
			uuid := "test-uuid"

			mockResponse := http.Result[PrividerResponseResult]{
				StatusCode: gohttp.StatusOK,
				Data: PrividerResponseResult{
					Entities: []PushProvider{},
				},
			}

			mockClient.EXPECT().Send(gomock.Any()).Return(mockResponse, nil)

			result, err := providerManager.DeletePushProvider(uuid)
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Entities).To(HaveLen(0))
		})

		It("should return an error if the request fails", func() {
			uuid := "test-uuid"

			mockClient.EXPECT().Send(gomock.Any()).Return(http.Result[PrividerResponseResult]{}, errors.New("request failed"))

			result, err := providerManager.DeletePushProvider(uuid)
			Expect(err).To(HaveOccurred())
			Expect(result.Entities).To(BeNil())
		})
	})

	Describe("ListPushProviders", func() {
		It("should list push providers successfully", func() {
			provider := PushProvider{
				Type: PushProviderAPNS,
				Name: "TestProvider",
			}

			mockResponse := http.Result[PrividerResponseResult]{
				StatusCode: gohttp.StatusOK,
				Data: PrividerResponseResult{
					Entities: []PushProvider{
						provider,
					},
				},
			}

			mockClient.EXPECT().Send(gomock.Any()).Return(mockResponse, nil)

			result, err := providerManager.ListPushProviders()
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Entities).To(HaveLen(1))
			Expect(result.Entities[0].Name).To(Equal("TestProvider"))
		})

		It("should return an error if the request fails", func() {
			mockClient.EXPECT().Send(gomock.Any()).Return(http.Result[PrividerResponseResult]{}, errors.New("request failed"))

			result, err := providerManager.ListPushProviders()
			Expect(err).To(HaveOccurred())
			Expect(result.Entities).To(BeNil())
		})
	})
})
