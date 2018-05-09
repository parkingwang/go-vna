package vna

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//
// Author: 陈永佳 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

type KVPair struct {
	Key   string
	Value string
}

func (f KVPair) String() string {
	return fmt.Sprintf("%s : %s", f.Key, f.Value)
}

func ReadRecords(name string) ([]KVPair, error) {
	file, err := os.Open(name)
	defer file.Close()

	if nil != err {
		return nil, err
	}

	output := make([]KVPair, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) > 2 && !strings.HasPrefix(txt, "#") {
			columns := strings.Split(txt, ",")
			if 2 == len(columns) {
				output = append(output, KVPair{
					Key:   columns[0],
					Value: columns[1],
				})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return output, nil
}
