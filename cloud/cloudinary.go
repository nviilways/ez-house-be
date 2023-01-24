package cloud

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/config"
	"github.com/cloudinary/cloudinary-go/v2"
)

var (
	Cloud *cloudinary.Cloudinary
)

func Connect() (err error) {
	Cloud, err = cloudinary.NewFromParams(config.CloudinaryCloud, config.CloudinaryApiKey, config.CloudinarySecretKey)
	return
}
