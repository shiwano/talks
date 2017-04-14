type Ping struct {
	Message string `json:"message" msgpack:"message"`
}

func (t *Ping) Coerce() error {
	return nil
}

func (t *Ping) Bytes(serializer typhenapi.Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}
	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}
	return data, nil
}
