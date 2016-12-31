package identity

type Identity interface {
  Addr() string
}

type NetIdentity string
func (n NetIdentity) Addr() string {
  return (string)(n)
}
var _ Identity = (*NetIdentity)(nil)
