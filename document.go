package mo


type Document interface {
	AfterInsert(data ResultInsertOne) (err error)
	BeforeInsert()(err error)
	BeforeUpdate() error
	AfterUpdate() error
}
type DefaultLifeCycle struct {

}
func (v *DefaultLifeCycle) BeforeInsert() error {return nil}
func (v *DefaultLifeCycle) AfterInsert(data ResultInsertOne) error {return nil}
func (v *DefaultLifeCycle) BeforeUpdate() error {return nil}
func (v *DefaultLifeCycle) AfterUpdate() error {return nil}

type ManyDocument interface {
	ManyD() (documents []interface{}, err error)
	AfterInsertMany(result ResultInsertMany) (err error)
}