package ifchain

func (s *serviceImpl) Save(obj *Object) error {
	// implement some code
	// transform object, update flag for the next
	// layer - the validation layer
	obj.UpdateState("implemented")
	s.log.Println("impl: ", obj)
	return s.ServiceIface.Save(obj)
}
