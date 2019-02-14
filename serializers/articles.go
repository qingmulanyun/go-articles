package serializers

type Articles struct {
	Tag string `json:"tag"`
	Count int `json:"count"`
	Articles []uint `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}