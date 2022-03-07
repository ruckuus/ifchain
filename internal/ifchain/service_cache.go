package ifchain

func (s *serviceCache) Save(obj *Object) error {
	// Do caching
	s.cache.Set("cache", obj)
	obj.UpdateState("cached")

	s.log.Println("cache: ", obj)
	return s.ServiceIface.Save(obj)
}
