package resolvers

// TODO implement different permissions and action structures for each resolver group

func IsSignedIn() bool {
	return true
}

func IsBackOfficer() bool {
	return true
}

func IsOwner() bool {
	return true
}

func IsManager() bool {
	return true
}

func CanRead() bool {
	return true
}

func CanWrite() bool {
	return true
}

func CanUpdate() bool {
	return true
}

func CanDelete() bool {
	return true
}
