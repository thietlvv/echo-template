package users

type Order struct {
	OrderId string
	Amount  int
}

type EventStore struct {
	EventId       string
	EventType     string
	EventData     string
	AggregateId   string
	AggregateType string
}

type ServiceDiscovery struct {
	OrderServiceUri string `json:"order_service_uri"`
}

const (
	Queue     = "Order.OrdersCreatedQueue"
	Subject   = "Order.OrderCreated"
	aggregate = "Order"
	event     = "OrderCreated"
)
