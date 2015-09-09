type Reader interface {
    Read(p []byte) (n int, err error)
}

type FixedByte byte

func (fb FixedByte) Read(p []byte) (n int, err error) {
  for i, _ := range p {
    p[i] = fb
  }

  return len(p), nil
}