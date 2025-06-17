package entity

type Availability bool

const (
	Available   Availability = true
	Unavailable Availability = false
)

type ImageUrlEntity struct {
	Url          string
	Availability Availability
}

func NewImageUrl(url string, availability Availability) *ImageUrlEntity {
	return &ImageUrlEntity{
		Url:          url,
		Availability: availability,
	}
}
