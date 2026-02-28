package delete

func Delete(cache map[string]string, key string) error {
	_, ok := cache[key]
	if !ok {
		return ErrKeyNotFound
	}

	delete(cache, key)
	return nil
}
