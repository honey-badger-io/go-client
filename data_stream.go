package hb

import "github.com/honey-badger-io/go-client/pb"

type SendStream struct {
	s pb.Data_CreateSendStreamClient
}

func (s *SendStream) Close() error {
	if err := s.s.CloseSend(); err != nil {
		return err
	}

	_, err := s.s.CloseAndRecv()

	return err
}

func (s *SendStream) Send(key string, data []byte) error {
	return s.s.Send(&pb.SendStreamReq{
		Item: &pb.DataItem{
			Key:  key,
			Data: data,
		},
	})
}
