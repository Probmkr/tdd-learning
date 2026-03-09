package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount}
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{d.amount * multiplier}
}

func (d *Dollar) Equals(dollar *Dollar) bool {
	return d.amount == dollar.amount
}

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{amount}
}

func (d *Franc) Times(multiplier int) *Franc {
	return &Franc{d.amount * multiplier}
}

func (d *Franc) Equals(franc *Franc) bool {
	return d.amount == franc.amount
}
