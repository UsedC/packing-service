package server

type GetPackSizesResponse struct {
	PackSizes []int `json:"pack_sizes"`
}

type SetPackSizesRequest struct {
	PackSizes []int `json:"pack_sizes"`
}

func (r *SetPackSizesRequest) Validate() error {
	if len(r.PackSizes) == 0 {
		return ErrEmptyPackSizes
	}

	for _, packSize := range r.PackSizes {
		if packSize <= 0 {
			return ErrInvalidPackSize
		}
	}

	return nil
}

type CalculatePacksRequest struct {
	Total int `json:"total"`
}

func (r *CalculatePacksRequest) Validate() error {
	if r.Total <= 0 {
		return ErrInvalidTotal
	}

	return nil
}

type CalculatePacksResponse struct {
	Packs map[int]int `json:"packs"`
}
