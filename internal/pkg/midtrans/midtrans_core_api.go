package midtrans

import (
	"github.com/midtrans/midtrans-go/coreapi"
)

func (coreApi *MidtransCoreApiImpl) ChargeTransaction(request *coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {
	response, err := coreApi.Client.ChargeTransaction(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}