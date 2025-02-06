package app

import "context"

type PackService interface {
	GetPackSizes(context.Context) []int
	SetPackSizes(context.Context, []int)
	CalculatePacks(context.Context, int) (map[int]int, error)
}

type App struct {
	packService PackService
}

func NewApp(packService PackService) *App {
	return &App{packService: packService}
}

func (a *App) GetPackSizes(ctx context.Context) []int {
	return a.packService.GetPackSizes(ctx)
}

func (a *App) SetPackSizes(ctx context.Context, packSizes []int) {
	a.packService.SetPackSizes(ctx, packSizes)
}

func (a *App) CalculatePacks(ctx context.Context, total int) (map[int]int, error) {
	return a.packService.CalculatePacks(ctx, total)
}
