package service

import (
	"context"
	"slices"
)

type PackService struct {
	packSizes      []int
	calculatorFunc func(int, []int) map[int]int
}

func NewPackService(packSizes []int, calculatorFunc func(int, []int) map[int]int) *PackService {
	return &PackService{
		packSizes:      packSizes,
		calculatorFunc: calculatorFunc,
	}
}

func (ps *PackService) GetPackSizes(_ context.Context) []int {
	return ps.packSizes
}

func (ps *PackService) SetPackSizes(_ context.Context, packSizes []int) {
	slices.Sort(packSizes)

	ps.packSizes = packSizes
}

func (ps *PackService) CalculatePacks(_ context.Context, total int) (map[int]int, error) {
	packs := ps.calculatorFunc(total, ps.packSizes)

	if packs == nil {
		return nil, ErrInvalidPackSizes
	}

	return packs, nil
}
