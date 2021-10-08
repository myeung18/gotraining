package pipeline

import "context"


type Payload interface {

}


type Processor interface {
	Process(context.Context, Payload) (Payload, error)
}

// ProcessFunc
type ProcessFunc func(context.Context, Payload) (Payload, error)

// Process
func (receiver ProcessFunc) Process(ctx context.Context, p Payload) (Payload, error) {
	receiver(ctx, p)
	return "", nil
}