package json

func Wrap(name string, item interface{}) map[string]interface{} {
	wrapped := map[string]interface{}{
		name: item,
	}
	return wrapped
}
