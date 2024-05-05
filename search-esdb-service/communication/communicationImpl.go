package communication

import "search-esdb-service/proto/ml_gateway_proto"

type CommunicationImpl struct {
	GRPC GRPCInterface
}

func NewCommunicationImpl(GRPC GRPCInterface) Communication {
	return &CommunicationImpl{
		GRPC: GRPC,
	}
}

func (c *CommunicationImpl) Text2Vec(text string) (*ml_gateway_proto.Text2VecResponse, error) {
	return c.GRPC.Text2Vec(text)
}
