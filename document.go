package mo


type Document interface {
	AfterInsert(data ResultInsertOne) (err error)
	BeforeInsert()(err error)
	BeforeUpdate() (err error)
	AfterUpdate() (err error)
}
type DefaultLifeCycle struct {

}
func (v *DefaultLifeCycle) BeforeInsert() error {return nil}
func (v *DefaultLifeCycle) AfterInsert(data ResultInsertOne) error {return nil}
func (v *DefaultLifeCycle) BeforeUpdate() (err error) {return nil}
func (v *DefaultLifeCycle) AfterUpdate() (err error) {return nil}

type ManyDocument interface {
	ManyD() (documents []interface{}, err error)
	AfterInsertMany(result ResultInsertMany) (err error)
}