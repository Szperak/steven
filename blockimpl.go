// Copyright 2015 Matthew Collins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"image"
)

// Stone

type stoneVariant int

const (
	stoneNormal stoneVariant = iota
	stoneGranite
	stoneSmoothGranite
	stoneDiorite
	stoneSmoothDiorite
	stoneAndesite
	stoneSmoothAndesite
)

func (s stoneVariant) String() string {
	switch s {
	case stoneNormal:
		return "stone"
	case stoneGranite:
		return "granite"
	case stoneSmoothGranite:
		return "smooth_granite"
	case stoneDiorite:
		return "diorite"
	case stoneSmoothDiorite:
		return "smooth_diorite"
	case stoneAndesite:
		return "andesite"
	case stoneSmoothAndesite:
		return "smooth_andesite"
	}
	return fmt.Sprintf("stoneVariant(%d)", s)
}

type blockstone struct {
	baseBlock
	Variant stoneVariant `state:"variant,0-6"`
}

func initStone(name string) *BlockSet {
	l := &blockstone{}
	l.init(name)
	set := alloc(l)
	return set
}

func (b *blockstone) String() string {
	return b.Parent.stringify(b)
}

func (b *blockstone) clone() Block {
	return &blockstone{
		baseBlock: *(b.baseBlock.clone().(*baseBlock)),
		Variant:   b.Variant,
	}
}

func (b *blockstone) ModelName() string {
	return b.Variant.String()
}

func (b *blockstone) toData() int {
	data := int(b.Variant)
	return data
}

// Grass

type blockGrass struct {
	baseBlock
	Snowy bool `state:"snowy"`
}

func initGrass() *BlockSet {
	g := &blockGrass{}
	g.init("grass")
	set := alloc(g)
	return set
}

func (g *blockGrass) String() string {
	return g.Parent.stringify(g)
}

func (g *blockGrass) clone() Block {
	return &blockGrass{
		baseBlock: *(g.baseBlock.clone().(*baseBlock)),
		Snowy:     g.Snowy,
	}
}

func (g *blockGrass) ModelVariant() string {
	return fmt.Sprintf("snowy=%t", g.Snowy)
}

func (g *blockGrass) TintImage() *image.NRGBA {
	return grassBiomeColors
}

func (g *blockGrass) toData() int {
	if g.Snowy {
		return -1
	}
	return 0
}

// Tall grass

type tallGrassType int

const (
	tallGrassDeadBush = iota
	tallGrass
	tallGrassFern
)

func (t tallGrassType) String() string {
	switch t {
	case tallGrassDeadBush:
		return "dead_bush"
	case tallGrass:
		return "tall_grass"
	case tallGrassFern:
		return "fern"
	}
	return fmt.Sprintf("tallGrassType(%d)", t)
}

type blockTallGrass struct {
	baseBlock
	Type tallGrassType `state:"type,0-2"`
}

func initTallGrass() *BlockSet {
	t := &blockTallGrass{}
	t.init("tallgrass")
	t.cullAgainst = false
	set := alloc(t)
	return set
}

func (t *blockTallGrass) String() string {
	return t.Parent.stringify(t)
}

func (t *blockTallGrass) clone() Block {
	return &blockTallGrass{
		baseBlock: *(t.baseBlock.clone().(*baseBlock)),
		Type:      t.Type,
	}
}

func (t *blockTallGrass) ModelName() string {
	return t.Type.String()
}

func (t *blockTallGrass) TintImage() *image.NRGBA {
	return grassBiomeColors
}

func (t *blockTallGrass) toData() int {
	return int(t.Type)
}
