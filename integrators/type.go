package integrators

type Integrators int

const (
	Discourse Integrators = iota
	Twitter
	Playstore
	Intercom
)

func (i Integrators) String() string {
	return [...]string{"Discourse", "Twitter", "Playstore", "Intercom"}[i]
}
