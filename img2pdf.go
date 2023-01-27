/***************************************************

Create PDF from a sequence of images.

- support 2 input image formats, like:
    - jpg/jpeg
    - png

***************************************************/

package main

import (
  "flag"
  "github.com/jung-kurt/gofpdf"
  "strings"
)

func main() {
  // Set PDF spec
  pdf := gofpdf.New("P", "mm", "A4", "")
  
  // flag to accept the input images filenames and output PDF filename as command-line arguments
  images := flag.String("images", "", "comma separated list of image file names")
  output := flag.String("output", "images.pdf", "name for the pdf output file name")
  flag.Parse()
  
  imageFiles := strings.Split(*images, ",")
  
  for _, imageFile := range imageFiles {
    pdf.AddPage()
    pdf.Image(imageFile, 10, 10, 150, 0, false, "", 0, "")
  }
  
  pdf.OutputFileAndClose(*output)
  
}
