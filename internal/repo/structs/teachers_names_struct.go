package structs

type TeachersNames struct {
	FakePk int   `reindex:"fake_PK,,pk" json:"-"`
	Names  Names `reindex:"names" json:"names"`
}

type Names []string

func (n *Names) Upsert(s string) {
	for _, v := range *n {
		if v == s {
			return
		}
	}

	*n = append(*n, s)
}
