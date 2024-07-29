/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

type DeviceManager struct {
	client *Client
}

func (dm *DeviceManager) List(userID string) error {
	// path := "users/" + userID + "/push/binding"
	// err := dm.client.REST(http.MethodGet, path, nil)
	// if err != nil {
	// 	fmt.Printf("error list all device info:%v", err)
	// }
	return nil
}
