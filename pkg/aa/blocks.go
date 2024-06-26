// apparmor.d - Full set of apparmor profiles
// Copyright (C) 2021-2024 Alexandre Pujol <alexandre@pujol.io>
// SPDX-License-Identifier: GPL-2.0-only

package aa

const (
	HAT Kind = "hat"
)

// Hat represents a single AppArmor hat.
type Hat struct {
	RuleBase
	Name  string
	Rules Rules
}

func (r *Hat) Validate() error {
	return nil
}

func (p *Hat) Less(other any) bool {
	o, _ := other.(*Hat)
	return p.Name < o.Name
}

func (p *Hat) Equals(other any) bool {
	o, _ := other.(*Hat)
	return p.Name == o.Name
}

func (p *Hat) String() string {
	return renderTemplate(p.Kind(), p)
}

func (p *Hat) Constraint() constraint {
	return blockKind
}

func (p *Hat) Kind() Kind {
	return HAT
}
