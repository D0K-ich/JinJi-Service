package transactions

import (
	"github.com/D0K-ich/JinJi-Service/network/messages"
)

func(h *Handler) Route(message *messages.Message) (payload any, err error) {

	switch message.SubjectAction() {
	//case "transactions/list"        : payload, err 	= h.Transactions().List(h.UserId())
	//case "transaction/status"       : payload, err 	= h.Transactions().StatusByOrderId(h.UserId(), incoming.String("order_id"))
	//case "transaction/create-ecom"  : payload, err 	= h.Transactions().CreateEcom(h.UserId(), incoming.Float("amount"), incoming.String("comment"))
	//
	//case "tariffs/list"			    : payload, err 	= h.Tariffs().List()
	//case "tariff/change"		    : payload, err 	= h.Tariffs().Change(h.UserId(), incoming.Int("tariff_id"))
	}

	return
}
