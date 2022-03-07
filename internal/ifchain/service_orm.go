package ifchain

func (s *serviceOrm) Save(obj *Object) error {
	obj.UpdateState("saved")

	s.log.Println("orm: ", obj)
	err := s.db.Set(obj.Name, obj)
	if err != nil {
		return err
	}
	return nil
}
