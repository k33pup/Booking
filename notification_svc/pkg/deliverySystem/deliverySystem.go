//mock delivery system

package deliverysystem

type IDeliverySystem struct {
	SendMessage func([]byte) error
}

type MockDeliverySystem struct {
}
