package main

import (
	"bytes"
	"fmt"
	"gotable"
	"os"
	"strconv"
	"time"

	"github.com/yosssi/gohtml"
)

var pdfProps = []*gotable.PDFProperty{
	// smart shrinking
	{Option: "--disable-smart-shrinking"},
	// top margin
	{Option: "-T", Value: "15"},
	// use custom dpi setting
	{Option: "--dpi", Value: "256"},
	// header center content
	{Option: "--header-center", Value: "Report Table"},
	// header font size
	{Option: "--header-font-size", Value: "7"},
	// header font
	{Option: "--header-font-name", Value: "opensans"},
	// header spacing
	{Option: "--header-spacing", Value: "3"},
	// bottom margin
	{Option: "-B", Value: "15"},
	// footer spacing
	{Option: "--footer-spacing", Value: "5"},
	// footer font
	{Option: "--footer-font-name", Value: "opensans"},
	// footer font size
	{Option: "--footer-font-size", Value: "7"},
	// footer left content
	{Option: "--footer-left", Value: time.Now().Format(gotable.DATETIMEFMT)},
	// footer right content
	{Option: "--footer-right", Value: "Page [page] of [toPage]"},
	// page size
	{Option: "--page-size", Value: "Letter"},
	// orientation
	{Option: "--orientation", Value: "Landscape"},
}

func main() {

	const (
		lorem = `Lorem      ipsum dolor sit amet, elementum
		fermentum suspendisse,illum est curabitur a sem justo rhoncus. Ac iaculis nec amet nisl,
		scelerisque sed quis nec dignissim dolor tempor, pellentesque tortor phasellus ut,
		risus eros sed primis vestibulum, vestibulum viverra. Maecenas orci scelerisque.
		Lectus cursus lorem etiam adipisicing, enim per tellus,
		purus mauris id dapibus qui curabitur nam, tincidunt nec gravida curabitur.`
	)

	var t gotable.Table
	t.Init()
	t.SetTitle("Go Table")
	t.SetSection1("Section One")
	t.SetSection1("Section Two")

	t.AddColumn("Line No", 7, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Unit", 14, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	t.AddColumn("Amount", 10, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	t.AddColumn("Description", 80, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
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
	t.AddLineAfter(t.RowCount() - 1)
	t.AddRow()
	t.Puts(-1, 0, "lastRow")
	t.Puts(-1, 1, "Unit - "+"LastRow")
	t.Puti(-1, 2, int64(1000000))
	t.Puts(-1, 3, lorem)

	// ==========
	// TEXT Output
	// ==========

	// generate text file
	tf, err := os.Create("table.txt")
	if err != nil {
		fmt.Printf("Error while creating TEXT output file: %s\n", err.Error())
		os.Exit(1)
	}
	defer tf.Close()

	if err := t.FprintTable(tf); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Text File generated successfully for table.")

	// ==========
	// CSV Output
	// ==========

	// generate text file
	cf, err := os.Create("table.csv")
	if err != nil {
		fmt.Printf("Error while creating CSV output file: %s\n", err.Error())
		os.Exit(1)
	}
	defer cf.Close()

	if err := t.CSVprintTable(cf); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("CSV File generated successfully for table.")

	// apply some css for html output
	c := []*gotable.CSSProperty{}
	c = append(c, &gotable.CSSProperty{Name: "border", Value: "1px black solid"})
	c = append(c, &gotable.CSSProperty{Name: "color", Value: "red"})
	t.SetRowCSS(1, c)
	// FIX the unit coversation bug
	t.SetColHTMLWidth(1, 15, "px")

	c = []*gotable.CSSProperty{}
	// c = append(c, &gotable.CSSProperty{Name: "color", Value: "blue"})
	// c = append(c, &gotable.CSSProperty{Name: "font-style", Value: "italic"})
	// c = append(c, &gotable.CSSProperty{Name: "font-size", Value: "20px"})
	t.SetTitleCSS(c)

	c = []*gotable.CSSProperty{}
	c = append(c, &gotable.CSSProperty{Name: "color", Value: "orange"})
	c = append(c, &gotable.CSSProperty{Name: "font-style", Value: "italic"})
	t.SetHeaderCSS(c)
	c = append(c, &gotable.CSSProperty{Name: "background-color", Value: "blue"})
	t.SetHeaderCSS(c)

	c = []*gotable.CSSProperty{}
	c = append(c, &gotable.CSSProperty{Name: "color", Value: "white"})
	c = append(c, &gotable.CSSProperty{Name: "background-color", Value: "black"})
	t.SetSection1CSS(c)

	c = []*gotable.CSSProperty{}
	c = append(c, &gotable.CSSProperty{Name: "vertical-align", Value: "top"})
	t.SetAllCellCSS(c)

	// ==========
	// HTML Output
	// ==========

	hf, err := os.Create("tableCSS.html")
	if err != nil {
		fmt.Printf("Error while creating HTML output file: %s\n", err.Error())
		os.Exit(1)
	}
	defer hf.Close()

	// write formatted html output
	if err = t.HTMLprintTable(hf); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("HTML File with custom css generated successfully for table.")

	// ==========
	// PDF Output
	// ==========

	var tbl gotable.Table
	tbl.Init()

	tbl.SetSection1("Section1\tSection1\tSection1\tSection1\tSection1\t")
	tbl.SetSection2("Section2\tSection2\tSection2\tSection2\tSection2\t")
	tbl.SetTitle("Accord Sample Report")
	tbl.AddColumn("No", 7, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Unit", 14, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)
	tbl.AddColumn("Amount", 6, gotable.CELLINT, gotable.COLJUSTIFYRIGHT)
	tbl.AddColumn("Description", 80, gotable.CELLSTRING, gotable.COLJUSTIFYLEFT)

	bf, err := os.Create("tableBlank.html")
	if err != nil {
		fmt.Printf("Error while creating PDF output file: %s\n", err.Error())
		os.Exit(1)
	}
	defer bf.Close()

	if err := tbl.HTMLprintTable(bf); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("BLANK file generated successfully for table.")

	rs = tbl.CreateRowset()
	for i := 0; i < 25; i++ {
		tbl.AddRow()
		tbl.Puts(-1, 0, strconv.Itoa(i+1))
		tbl.Puts(-1, 1, "Unit-"+strconv.Itoa(i+1))
		tbl.Puti(-1, 2, int64(i*1000))
		tbl.Puts(-1, 3, lorem)
		tbl.AppendToRowset(rs, i)
	}
	tbl.AddLineAfter(tbl.RowCount() - 1)
	tbl.InsertSumRowsetCols(rs, tbl.RowCount(), []int{2})

	tbl.TightenColumns()

	pf, err := os.Create("table.pdf")
	if err != nil {
		fmt.Printf("Error while creating PDF output file: %s\n", err.Error())
		os.Exit(1)
	}
	defer pf.Close()

	if err := tbl.PDFprintTable(pf, pdfProps); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("PDF file generated successfully for table.")

	hf, err = os.Create("table.html")
	if err != nil {
		fmt.Printf("Error while creating HTML output file: %s\n", err.Error())
		os.Exit(1)
	}
	defer hf.Close()

	// write formatted html output
	if err = tbl.HTMLprintTable(hf); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("HTML File generated successfully for table.")

	// ============
	// custom html template
	// ===========
	var customHTMLBuffer bytes.Buffer
	for i := 0; i < 10; i++ {
		var temp bytes.Buffer
		var t gotable.Table
		t = tbl
		if i == 0 {
			t.SetHTMLTemplate("templates/first.html")
		} else if i == 9 {
			t.SetHTMLTemplate("templates/last.tmpl")
		} else {
			t.SetHTMLTemplate("templates/middle.html")
		}
		t.HTMLprintTable(&temp)
		customHTMLBuffer.Write(temp.Bytes())
	}
	chf, err := os.Create("customTemplateHTML.html")
	if err != nil {
		fmt.Printf("Error while creating HTML output file: %s\n", err.Error())
		os.Exit(1)
	}
	defer chf.Close()
	chf.WriteString(gohtml.Format(customHTMLBuffer.String()))
	fmt.Println("HTML file with custom template path created successfully.")
}
