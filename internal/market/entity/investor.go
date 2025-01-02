package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}

type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}

func NewInvestor(id, name string) *Investor {
	return &Investor{
		ID:            id,
		Name:          name,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}

func (i *Investor) UpsertAssetPosition(assetID string, shares int) {
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetID, shares))
	} else {
		assetPosition.Shares += shares
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for asset := range i.AssetPosition {
		if i.AssetPosition[asset].AssetID == assetID {
			return i.AssetPosition[asset]
		}
	}
	return nil
}
