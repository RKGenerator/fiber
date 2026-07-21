package basefunc

import (
	"crypto/pbkdf2"

	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"test-fiber/dto"
	apperrors "test-fiber/errors"
	//"golang.org/x/crypto/pbkdf2"
)

func CheckPaginationRequest(req dto.PaginatorRequest) error {
	if req.Page < 1 || req.PageSize < 1 {
		return apperrors.ErrBadRequest
	}
	return nil
}

func CheckPassword(password, hashed string) bool {
	parts := strings.Split(hashed, "$")
	if len(parts) != 4 {
		return false
	}

	iterations := parts[1]
	salt := parts[2]
	hash := parts[3]

	iter, err := strconv.Atoi(iterations)
	if err != nil {
		return false
	}

	storedHash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		fmt.Println(err)
		return false
	}

	newHash, err := pbkdf2.Key(sha256.New, password, []byte(salt), iter, len(storedHash))
	if err != nil {
		return false
	}

	return subtle.ConstantTimeCompare(newHash, storedHash) == 1
}

func ChoiseOrder(reqOrder *string) (*string, bool) {
	var name_column *string
	var desc bool
	if reqOrder == nil {
		return nil, true
	}

	res := strings.Split(*reqOrder, ".")
	if res[0] != "" {
		name_column = &res[0]
	} else {
		name_column = nil
	}

	if res[1] == "" || res[1] != "asc" {
		desc = true
	} else {
		desc = false
	}

	return name_column, desc
}
