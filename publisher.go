package colectivo

import "errors"

type Publisher struct {
	Subs []Subscriber
}

func (p *Publisher) Publish(e Event) error {
	if e.Callback != nil {
		if err := e.Callback(); err != nil {
			return err
		}
	}

	var err error
	for _, sub := range p.Subs {
		if serr := sub.Handle(e); serr != nil {
			err = errors.Join(serr)
		}
	}

	return err
}
