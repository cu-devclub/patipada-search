package communication

import "search-esdb-service/proto/ml_gateway_proto"

type Communication interface {
	Text2Vec(text string) (*ml_gateway_proto.Text2VecResponse, error)
}
