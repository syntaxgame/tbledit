package editor

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Export(inFile, outFile string) {

	data, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Fatal(err)
	}

	f := excelize.NewFile()

	buff := NewBuffer(data)
	colSize := buff.Read(UINT32).(uint32)

	cols, headers := []uint32{}, []string{}
	for i := uint32(0); i < colSize; i++ { // get columns

		col := buff.Read(UINT32).(uint32)
		headers = append(headers, typeTitles[ColType(col)])
		cols = append(cols, col)
	}

	f.SetSheetRow("Sheet1", "A1", &headers) // set column headers

	rowSize := buff.Read(UINT32).(uint32)
	for i := uint32(0); i < rowSize; i++ {

		row := []interface{}{}
		for _, col := range cols {

			colType := ColType(col)
			if colType == STRING {
				textSize := buff.Read(UINT32).(uint32)
				text := BytesToString(buff.ReadN(uint64(textSize)))
				row = append(row, text)

			} else {
				row = append(row, buff.Read(ColType(col)))
			}
		}

		f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i+2), &row) // set row
	}

	f.SaveAs("test.xlsx")
}

func BytesToString(data []byte) string {

	s := ""
	for _, d := range data {
		if d < 128 {
			s += fmt.Sprintf("%c", d)
		} else {
			//s += fmt.Sprintf("%c", ansiChars[d-128])
			s += "?"
		}
	}
	return s
}
