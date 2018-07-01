/*
 * Revision History:
 *     Initial: 2018/05/31        Tong Yuehong
 */

package excelfile

import (
	"time"
	"encoding/base64"
	"strings"
	"io/ioutil"
	"fmt"
)

const (
	file = "./file"
	filesuffix = "xlsx"
)
func fileName() string {
	timestamp := time.Now().Unix()

	time := time.Unix(timestamp, 0).Format("2006-01-02 03:04:05 PM")
	time = strings.Replace(time, " ", "", 2)

	return file + time + "." + filesuffix
}

// SaveFile -
func SaveAvatar(userID uint32, image string) (string, error) {
	fileName := fileName()

	img, _ := base64.StdEncoding.DecodeString(image)
	err := ioutil.WriteFile(fileName, []byte(img), 0777)
	if err != nil {
		fmt.Println(err)
	}

	return fileName, err
}
