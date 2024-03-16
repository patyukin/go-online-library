package sender

type Sender struct {
}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) Send() error {
	return nil
}
