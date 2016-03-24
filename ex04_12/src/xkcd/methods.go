package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func FetchComicInfo(begNum, endNum int, finishWithLackIndex bool) []*ComicInfo {
	infos := make([]*ComicInfo, 0, endNum - begNum + 1)

	for n := begNum; n <= endNum; n++ {
		info, err := requestComicInfo(n)

		if err != nil {
			if finishWithLackIndex {
				break
			} else {
				continue
			}
		} else {
			infos = append(infos, info)
		}
	}

	return infos
}

func LoadComicInfo(path string) ([]*ComicInfo, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	var infos []*ComicInfo
	if err := json.NewDecoder(file).Decode(&infos); err != nil {
		return nil, err
	} else {
		return infos, nil
	}
}

func SaveComicInfo(path string, infos []*ComicInfo) error {
	file, err := os.Create(path)

	if err != nil {
		return  err
	}

	return json.NewEncoder(file).Encode(infos)
}

func Filtering(items []*ComicInfo, filter func(*ComicInfo) bool) []*ComicInfo {
	result := []*ComicInfo{}

	for _, item := range items {
		if(filter(item)) {
			result = append(result, item)
		}
	}

	return result
}


func requestComicInfo(num int) (*ComicInfo, error) {
	resp, err := http.Get(getURL(num))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search info failed : num = %d", num)
	}

	var result ComicInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result.validate(num), nil
}
