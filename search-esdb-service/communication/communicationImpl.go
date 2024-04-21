package communication

type CommunicationImpl struct {
	GRPC     GRPCStruct
}

func NewCommunicationImpl(GRPC GRPCStruct) Communication {
	return &CommunicationImpl{
		GRPC:     GRPC,
	}
}
