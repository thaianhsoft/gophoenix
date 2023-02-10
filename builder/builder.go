package builder

type Querier interface {
	Query() (query string, args []interface{})
}

type SelectView interface{
	implementView()
}
// selector define for select statement querier
type Selector struct {
	*CoreBuilder
	selects []SelectView
}

func (Selector) implementView() {}
func Select(sl SelectView) *Selector {
	s := &Selector{CoreBuilder: &CoreBuilder{}}
	s.selects = append(s.selects, sl)
}

func (s *Selector) Query() (query string, args []interface{}){

}