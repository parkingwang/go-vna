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

type Fields struct {
	Short string
	Name  string
}

func (f Fields) String() string {
	return fmt.Sprintf("%s : %s", f.Short, f.Name)
}

func ReadFields(name string) ([]Fields, error) {
	file, err := os.Open(name)
	defer file.Close()

	if nil != err {
		return nil, err
	}

	out := make([]Fields, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) > 2 && !strings.HasPrefix(txt, "#") {
			columns := strings.Split(txt, ",")
			if 2 == len(columns) {
				out = append(out, Fields{
					Short: columns[0],
					Name:  columns[1],
				})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return out, nil
}
