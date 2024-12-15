//mock delivery system

package deliverysystem

type IDeliverySystem interface {
	SendMail(string) error
}
