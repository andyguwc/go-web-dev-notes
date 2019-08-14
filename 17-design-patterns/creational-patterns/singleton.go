/* Singleton Pattern

Example implementating a counter
*/


package creational

type singleton struct {
	count int
}

// initializes by using a pointer. We cannot initialize a struct to be nil
var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}

	return instance
}

// AddOne() method which increases the count of the counter 
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

func (s *singleton) GetCount() int {
	return s.count
}




