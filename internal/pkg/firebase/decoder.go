package firebase

import (
	"encoding/base64"

	"github.com/HarsaEdu/harsa-api/configs"
)

func GetDecodedFireBaseKey(c configs.FirebaseConfig) ([]byte, error) {

	fireBaseAuthKey := c.FirebaseAuthKey

	decodedKey, err := base64.StdEncoding.DecodeString(fireBaseAuthKey)
	if err != nil {
		return nil, err
	}

	return decodedKey, nil
}
