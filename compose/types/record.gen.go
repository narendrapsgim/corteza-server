package types

// 	Hello! This file is auto-generated.

type (

	// RecordSet slice of Record
	//
	// This type is auto-generated.
	RecordSet []*Record
)

// Walk iterates through every slice item and calls w(Record) err
//
// This function is auto-generated.
func (set RecordSet) Walk(w func(*Record) error) (err error) {
	for i := range set {
		if err = w(set[i]); err != nil {
			return
		}
	}

	return
}

// Filter iterates through every slice item, calls f(Record) (bool, err) and return filtered slice
//
// This function is auto-generated.
func (set RecordSet) Filter(f func(*Record) (bool, error)) (out RecordSet, err error) {
	var ok bool
	out = RecordSet{}
	for i := range set {
		if ok, err = f(set[i]); err != nil {
			return
		} else if ok {
			out = append(out, set[i])
		}
	}

	return
}

// FindByID finds items from slice by its ID property
//
// This function is auto-generated.
func (set RecordSet) FindByID(ID uint64) *Record {
	for i := range set {
		if set[i].ID == ID {
			return set[i]
		}
	}

	return nil
}

// IDs returns a slice of uint64s from all items in the set
//
// This function is auto-generated.
func (set RecordSet) IDs() (IDs []uint64) {
	IDs = make([]uint64, len(set))

	for i := range set {
		IDs[i] = set[i].ID
	}

	return
}
