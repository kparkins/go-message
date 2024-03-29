package message

type Feed[T any] struct {
	message *message[T]
}

func NewFeed[T any](feed *Publisher[T]) *Feed[T] {
	return &Feed[T]{
		message: feed.head,
	}
}

func (s *Feed[T]) Value() T {
	var value T
	if s.message != nil {
		return s.message.data
	}
	return value
}

func (s *Feed[T]) Updated() chan struct{} {
	if s.message != nil {
		return s.message.ready
	}
	// if we have unsubscribed, return a closed channel so the
	// client is immediately unblocked
	value := make(chan struct{})
	close(value)
	return value
}

func (s *Feed[T]) Next() bool {
	if s.message == nil {
		return false
	}
	finished := s.message.finished
	s.message = s.message.next
	return !finished
}

func (s *Feed[T]) Finished() bool {
	if s.message == nil {
		return true
	}
	return s.message.finished
}

func (s *Feed[T]) Unsubscribe() {
	s.message = nil
}
