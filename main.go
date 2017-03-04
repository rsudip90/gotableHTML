package main

import (
	"fmt"
	gt "gotable"
	"os"
	"strconv"

	"github.com/yosssi/gohtml"
)

func main() {

	const (
		lorem = `Lorem ipsum dolor sit amet, elementum fermentum suspendisse,illum est curabitur a sem justo rhoncus. Ac iaculis nec amet nisl,
		scelerisque sed quis nec dignissim dolor tempor, pellentesque tortor phasellus ut,
		risus eros sed primis vestibulum, vestibulum viverra. Maecenas orci scelerisque.
		Lectus cursus lorem etiam adipisicing, enim per tellus,
		purus mauris id dapibus qui curabitur nam, tincidunt nec gravida curabitur.`
	)

	var t gt.Table
	t.Init()
	t.SetSection1(lorem)
	t.SetTitle("Go Table")
	t.AddColumn("Line No", 7, gt.CELLSTRING, gt.COLJUSTIFYLEFT)
	t.AddColumn("Unit", 14, gt.CELLSTRING, gt.COLJUSTIFYLEFT)
	t.AddColumn("Amount", 10, gt.CELLINT, gt.COLJUSTIFYLEFT)
	t.AddColumn("Description", 80, gt.CELLSTRING, gt.COLJUSTIFYLEFT)
	rs := t.CreateRowset()
	for i := 0; i < 5; i++ {
		t.AddRow()
		t.Puts(-1, 0, strconv.Itoa(i+1))
		t.Puts(-1, 1, "Unit - "+strconv.Itoa(i+1))
		t.Puti(-1, 2, int64(i*1000))
		t.Puts(-1, 3, lorem)
		t.AppendToRowset(rs, i)
	}
	t.AddLineAfter(t.RowCount() - 1)
	t.InsertSumRowsetCols(rs, t.RowCount(), []int{2})

	// generate text file
	tableText, err := t.SprintTable(gt.TABLEOUTTEXT)
	tf, err := os.Create("table.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer tf.Close()
	tf.WriteString(tableText)
	fmt.Println("Text File generated successfully for table.")

	// generate html file
	c := []*gt.CSSProperty{}
	c = append(c, &gt.CSSProperty{Name: "border", Value: "1px black solid"})
	c = append(c, &gt.CSSProperty{Name: "color", Value: "red"})
	t.SetRowCSS(1, c)
	t.SetColWidth(1, 150, "px")

	c = []*gt.CSSProperty{}
	c = append(c, &gt.CSSProperty{Name: "color", Value: "blue"})
	c = append(c, &gt.CSSProperty{Name: "font-style", Value: "italic"})
	c = append(c, &gt.CSSProperty{Name: "font-size", Value: "20px"})
	t.SetTitleCSS(c)

	c = []*gt.CSSProperty{}
	c = append(c, &gt.CSSProperty{Name: "color", Value: "orange"})
	c = append(c, &gt.CSSProperty{Name: "font-style", Value: "italic"})
	t.SetHeaderCSS(c)
	c = append(c, &gt.CSSProperty{Name: "background-color", Value: "blue"})
	t.SetHeaderCSS(c)

	tableHTML, err := t.SprintTable(gt.TABLEOUTHTML)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	f, err := os.Create("table.html")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	// write formatted html output
	f.WriteString(gohtml.Format(tableHTML))
	fmt.Println("HTML File generated successfully for table.")
}
