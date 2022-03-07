package ifchain

import "errors"

func (s *serviceValidator) Save(obj *Object) error {
	if obj.state != "implemented" {
		return errors.New("object not implemented")
	}

	if obj.Name == "" {
		return errors.New("name is required")
	}

	if obj.Version == "" {
		return errors.New("version is required")
	}

	obj.UpdateState("validated")

	s.log.Println("validator: ", obj)
	return s.ServiceIface.Save(obj)
}
