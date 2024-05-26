package main

import (
	"fmt"
	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	err := pdf.AddTTFFont("Montserrat", "Montserrat-Medium.ttf") // Adjust the path to your font file
	if err != nil {
		fmt.Println("Error loading font:", err)
		return
	}

	pdf.AddPage()

	                    /*############## HEADER - CONTENT ##############*/
	                    
	// Header
	pdf.SetX(65) // Adjust X position as needed
	pdf.SetY(50) // Adjust Y position as needed

	// "Ease" with color #064747 and Montserrat-BoldItalic
	pdf.SetFont("Montserrat", "", 20)
	pdf.SetTextColor(6, 71, 71)
	pdf.Cell(nil, "Ease")

	// "Of" with color #F48132 and Montserrat-BoldItalic
	pdf.SetFont("Montserrat", "", 20)
	pdf.SetTextColor(244, 129, 50)
	pdf.Cell(nil, "Of")

	// "Learning" with color #064747 and Montserrat-BoldItalic
	pdf.SetFont("Montserrat", "", 20)
	pdf.SetTextColor(6, 71, 71)
	pdf.Cell(nil, "Learning")

                                       /*########################################################*/
                                       

	                    /*############## Affordable | Adoptable | Accessible ##############*/

                  // Add subtitle below the header
                  subtitleFont := "Montserrat"
                  subtitleFontSize := 8.0
                  //subtitleColor := color.RGBA{6, 71, 71, 255} // Color: #064747

                  pdf.SetFont(subtitleFont, "", subtitleFontSize)
                  //pdf.SetTextColor(subtitleColor.R, subtitleColor.G, subtitleColor.B)

                  // Position the subtitle below the header
                  subtitleX := 75.0 // Same as X position for the header
                  subtitleY := 73.0 // Adjust Y position to place it below the header

                  // Add the subtitle text
                  subtitle := "Affordable | Adoptable | Accessible"
                  pdf.SetX(subtitleX)
                  pdf.SetY(subtitleY)
                  pdf.Cell(nil, subtitle)
                  pdf.Br(20) // Add vertical spacing


                                       /*########################################################*/



	                      /*############## INVOICE - CONTENT ##############*/
	                    
                  // Add subtitle below the header
                  invoiceFont := "Montserrat"
                  invoiceFontSize := 21.0
                  //subtitleColor := color.RGBA{6, 71, 71, 255} // Color: #064747

                  pdf.SetFont(invoiceFont, "", invoiceFontSize)
                  invoicetitleX := 450.0 // Same as X position for the header
                  invoicetitleY := 90.0 // Adjust Y position to place it below the header

                  // Add the subtitle text
                  invoiceSubtitle := "INOVICE"
                  pdf.SetX(invoicetitleX)
                  pdf.SetY(invoicetitleY)
                  pdf.Cell(nil, invoiceSubtitle)
                  pdf.Br(20) // Add vertical spacing

                                       /*########################################################*/
                                       
                                       
                  // Add Receipt no: EOL01
                  pdf.SetX(450) // Adjust X position for "Receipt no: EOL01"
                  pdf.SetY(110)  // Adjust Y position for "Receipt no: EOL01"
                  pdf.SetFont("Montserrat", "", 10)
                  pdf.SetTextColor(0, 0, 0) // Black color for "Receipt no: EOL01"
                  pdf.Cell(nil, "Receipt no: EOL01")


                  // Add section above the table
                  pdf.SetX(30) // Adjust X position as needed
                  pdf.SetY(120) // Adjust Y position as needed
                  sectionFont := "Montserrat"
                  sectionFontSize := 10.0
                  pdf.SetFont(sectionFont, "B", sectionFontSize)
                  
                  pdf.SetTextColor(0, 0, 0) // Black color
                  pdf.SetX(70) // Adjust X position as needed
                  pdf.SetY(150) // Adjust Y position as needed
                  pdf.Cell(nil, "PAYMENT SUMMARY:")
                  pdf.Br(20) // Add vertical spacing

                  details := map[string]string{
                      "Name"    :      "ABCED",
                      "Ph no"   : "1234567890",
                      "Emailid" :   "abc@gmail.com",
                      "Referal" : "ABCDEF",
                      "Course"  : "Master your interview soft skills(MYISS)",
                  }

	// Order of the details
	detailKeys := []string{"Name", "Ph no", "Emailid", "Referal", "Course"}


	for _, key := range detailKeys {
		value := details[key]
		pdf.SetX(70) // Reset X position for each line
		pdf.Cell(nil, key+": ")
		pdf.SetX(120) // Adjust X position for the value
		pdf.Cell(nil, value)
		pdf.Br(15) // Add vertical spacing between details
	}

	
	// Draw margin lines with color #064747
                  pdf.SetLineWidth(0.5)
                  pdf.SetStrokeColor(6, 71, 71) // Set margin lines color to #064747
                  pageWidth := 595.28           // Width of A4 in points
                  pageHeight := 841.89          // Height of A4 in points
                  margin := 30.0                // Increased margin for better appearance
                  pdf.Line(margin, margin, pageWidth-margin, margin)                     // Top margin line
                  pdf.Line(pageWidth-margin, margin, pageWidth-margin, pageHeight-margin) // Right margin line
                  pdf.Line(pageWidth-margin, pageHeight-margin, margin, pageHeight-margin) // Bottom margin line
                  pdf.Line(margin, pageHeight-margin, margin, margin)                     // Left margin line

                                       /*########################################################*/
                                       
                                       
	                        /*############## TABLE - CONTENT ##############*/
	                        
	                    
                  // Add table headers
                  pdf.SetX(70)             // Adjust X position as needed
                  pdf.SetY(270)                // Adjust Y position as needed
                  headerFont := "Montserrat"
                  headerFontSize := 10.0
                  pdf.SetFont(headerFont, "", headerFontSize)

                  columnWidth := (pageWidth - 2*margin) / 7 // Divide the available width by the number of columns
                  cellHeight := 20.0                        // Set the height for each cell

                  headers := []string{"S.NO", "Course", "Fee", "Amt Paid", "TransID", "Total"}


                  for _, header := range headers {
                      // Draw left border for header
                      pdf.SetLineWidth(0.5)
                      pdf.SetStrokeColor(6, 71, 71)
                      pdf.Line(pdf.GetX(), pdf.GetY(), pdf.GetX(), pdf.GetY()+cellHeight) // Left border for the current header cell
    
    
                  // Draw top border
                      pdf.SetLineWidth(0.5)
                      pdf.SetStrokeColor(6, 71, 71)
                      pdf.Line(pdf.GetX(), pdf.GetY(), pdf.GetX()+columnWidth, pdf.GetY()) // Top border for the current header cell
    
                      pdf.CellWithOption(&gopdf.Rect{W: columnWidth, H: cellHeight}, header, gopdf.CellOption{
                          Border: gopdf.Bottom | gopdf.Right, // Add bottom and right borders
                          Align:  gopdf.Center,
                          Float:  gopdf.Right,
                      })
                  }
                  pdf.Br(20) // Add bottom margin for the table
	

                  // Add table rows (dummy data)
                  rows := [][]string{
                      {"1", "MYISS", "599", "599", "123456", "599/-"},
                      {"2", "", "", "", "", ""},
                      {"Total Amount", "", "", "", "", ""}, 
                  }


                  // Add "599/-" to the last column of the last row
                  rows[len(rows)-1][len(headers)-1] = "599/-"

                  dataFont := "Montserrat"
                  dataFontSize := 8.0

                  pdf.SetFont(dataFont, "", dataFontSize)

                  for i, row := range rows {
                      pdf.SetX(70) // Reset X position for each row

                      for j, cell := range row {
                          pdf.SetLineWidth(0.5)
                          pdf.SetStrokeColor(6, 71, 71)

                  // Conditionally adjust border settings for the first row
                  if i == 0 {
                  // Add top border for the first row
                  pdf.Line(pdf.GetX(), pdf.GetY(), pdf.GetX()+columnWidth, pdf.GetY())
                  }

                  // Conditionally adjust border settings for the last row
                  if i == len(rows)-1 {
                  if j == 0 {
                  // For the first cell in the last row, add bottom border only
                  pdf.Line(pdf.GetX(), pdf.GetY(), pdf.GetX(), pdf.GetY()+cellHeight) // Left border for the current cell
                  pdf.CellWithOption(&gopdf.Rect{W: columnWidth, H: cellHeight}, cell, gopdf.CellOption{
                     Border: gopdf.Bottom | gopdf.Right, // Add bottom and right borders
                     Align:  gopdf.Center,
                     Float:  gopdf.Right,
                 })
                  } else if j == len(row)-1 {
                // For the last cell in the last row, add bottom border and right border
                pdf.CellWithOption(&gopdf.Rect{W: columnWidth, H: cellHeight}, cell, gopdf.CellOption{
                    Border: gopdf.Bottom | gopdf.Right, // Add bottom and right borders
                    Align:  gopdf.Center,
                    Float:  gopdf.Right,
                })
            } else {
                // For other cells in the last row, add bottom border only
                pdf.CellWithOption(&gopdf.Rect{W: columnWidth, H: cellHeight}, cell, gopdf.CellOption{
                    Border: gopdf.Bottom, // Add bottom border only
                    Align:  gopdf.Center,
                    Float:  gopdf.Right,
                })
            }
        } else {
            // For other rows, add bottom and right borders
            if j == 0 {
                // For the first cell in the row, add left border
                pdf.Line(pdf.GetX(), pdf.GetY(), pdf.GetX(), pdf.GetY()+cellHeight) // Left border for the current cell
            }
            pdf.CellWithOption(&gopdf.Rect{W: columnWidth, H: cellHeight}, cell, gopdf.CellOption{
                Border: gopdf.Bottom | gopdf.Right, // Add bottom and right borders
                Align:  gopdf.Center,
                Float:  gopdf.Right,
            })
        }
    }
    pdf.Br(cellHeight) // Move to the next row
}


                                       /*########################################################*/
                                       
                                       

	                    /*############## PAYMENT - CONTENT ##############*/	
	
                  // Add Payment Mode: Paytm
                  pdf.SetX(70) // Adjust X position for "Payment Mode: Paytm"
                  pdf.SetY(410) // Adjust Y position for "Payment Mode: Paytm"
                  pdf.SetFont("Montserrat", "", 10)
                  pdf.SetTextColor(0, 0, 255) // Blue color for "Payment Mode: Paytm"
                  pdf.Cell(nil, "Payment Mode: Paytm")

                  // Add Paytm logo beside the text
	imageWidth := 15.0  // Set the desired width of the image
	imageHeight := 15.0 // Set the desired height of the image
	xPos := pdf.GetX() + 2 // Adjust X position for the image
	yPos := pdf.GetY() - 2 // Adjust Y position for the image (slightly above the baseline of the text)

	paytmLogoPath := "images/successfullTransaction.png" // Path to your Paytm logo
	err = pdf.Image(paytmLogoPath, xPos, yPos, &gopdf.Rect{W: imageWidth, H: imageHeight})
	if err != nil {
		fmt.Println("Error loading Paytm logo:", err)
	} else {
		fmt.Println("Paytm logo added successfully")
	}
	
	
                                       /*########################################################*/
                                       
                                
                          // Add text "FIVE THOUSAND NINE HUNDRED AND NINETY NINE ONLY" above the payment mode
                  pdf.SetFont("Montserrat", "", 8) // Set font and size
                  pdf.SetTextColor(169, 169, 169) // Light grey color
                  pdf.SetX(160) // Adjust X position for the text
                  pdf.SetY(335) // Adjust Y position for the text
                  pdf.Cell(nil, "FIVE HUNDRED AND NINETY NINE RUPEES ONLY")
                  pdf.Br(10) // Add vertical spacing
      
                                
                                       
                                       
                  // Add text below the Payment Mode
                  thankYouText1 := "Thank you for choosing EaseOfLearning. Your enrollment for the course 'Master your interview"
                  thankYouText2 := " soft skills(MYISS)' has been successful. Save this Invoice for future reference."
                  pdf.SetX(110) // Adjust X position for the text
                  pdf.SetY(510) // Adjust Y position for the text
                  pdf.SetFont("Montserrat", "", 8) // Set font and size
                  pdf.SetTextColor(0, 0, 0) // Black color for the text
                  pdf.Cell(nil, thankYouText1)
                  pdf.Br(10) // Add vertical spacing
                  pdf.SetX(110) // Adjust X position for the second part of the text
                  pdf.Cell(nil, thankYouText2)
                  pdf.Br(20) // Add more vertical spacing if needed


                  // Calculate the width and height of the box based on the length of the text
                  boxWidth := 400.0
                  boxHeight := 40.0

                  // Set the position of the box
                  boxX := 100.0
                  boxY := 500.0

                  // Draw the box
                  pdf.SetLineWidth(1) // Set the border width
                  pdf.SetStrokeColor(52, 168, 83) // Border color #34a853
                  pdf.RectFromUpperLeftWithStyle(boxX, boxY, boxWidth, boxHeight, "D") // Draw the rectangle (D for drawing only)

                  // Add thank you text inside the box
                  pdf.SetX(boxX + 10) // Adjust X position for the text inside the box
                  pdf.SetY(boxY + 10) // Adjust Y position for the text inside the box
                                       
                                       
                  // Add image above the e-signature
                  thankyouImagePath := "images/padiSTamp.png" // Path to your thank you image
                  thankyouImageWidth := 100.0  // Set the desired width of the image
                  thankyouImageHeight := 100.0  // Set the desired height of the image
                  thankyouXPos := 200.0         // Adjust X position for the image
                  thankyouYPos := 360.0        // Adjust Y position for the image

                  err = pdf.Image(thankyouImagePath, thankyouXPos, thankyouYPos, &gopdf.Rect{W: thankyouImageWidth, H:   
                  thankyouImageHeight})
                  if err != nil {
                  fmt.Println("Error loading thank you image:", err)
                  } else {
                  fmt.Println("Thank you image added successfully")
                  }
                                       
                              
                  // Add text beside the padiSTamp.png image
                  paidAtText := "Paid at 03:05 PM, 23 May 2024"
                  pdf.SetX(400) // Adjust X position for the text
                  pdf.SetY(thankyouYPos + (thankyouImageHeight / 2)) // Adjust Y position for the text to center it vertically

                  pdf.SetFont("Montserrat", "", 8) // Set font and size
                  pdf.SetTextColor(0, 0, 0) // Black color for the text
                  pdf.Cell(nil, paidAtText)
	
	                    /*############## e-signature - CONTENT ##############*/
	                    
	signatureImagePath := "images/signature.png" // Path to your e-signature image
	signatureImageWidth := 100.0  // Set the desired width of the e-signature image
	signatureImageHeight := 50.0  // Set the desired height of the e-signature image
	signatureXPos := 460.0         // Adjust X position for the e-signature image
	signatureYPos := 565.0        // Adjust Y position for the e-signature image

	err = pdf.Image(signatureImagePath, signatureXPos, signatureYPos, &gopdf.Rect{W: signatureImageWidth, H:   
	signatureImageHeight})
	if err != nil {
		fmt.Println("Error loading e-signature:", err)
	} else {
		fmt.Println("E-signature added successfully")
	}
	
		
	signatureFont := "Montserrat"
	signatureFontSize := 10.0
	pdf.SetFont(signatureFont, "", signatureFontSize)
	pdf.SetTextColor(0, 0, 0) // Black color for e-signature
	pdf.SetX(480) // Adjust X position for the e-signature
	pdf.SetY(600) // Adjust Y position for the e-signature

	// Add signature text
	pdf.SetFont("Montserrat", "", 8) // Reduced font size for the signature text
	pdf.Cell(nil, "RajeshMamidi")
	pdf.Br(15) // Move to the next line
	pdf.SetX(500) // Adjust X position for the designation
	pdf.SetY(610)
	pdf.Cell(nil, "HOD")

                                         /*########################################################*/
                                         
                                         
                                         
                                         
                                         
                                         /*############## NOTE - CONTENT ##############*/
                                         
                  // Add NOTE above the footer
                  noteFont := "Montserrat"
                  noteFontSize := 10.0
                  pdf.SetFont(noteFont, "B", noteFontSize) // Bold font
                  pdf.SetTextColor(255, 0, 0)              // Red color
                  pdf.SetX(70)                             // Adjust X position for NOTE
                  pdf.SetY(signatureYPos + signatureImageHeight + 40) // Adjust Y position for NOTE

                  // First part of the text
                  noteTextPart1 := "NOTE: This is an electronically generated Invoice please feel free to reach us at   easeoflearning06@gmail.com"
                  pdf.Cell(nil, noteTextPart1)

// Second part of the text
noteTextPart2 := "for any payment related queries."
pdf.SetX(90) // Reset X position
pdf.SetY(pdf.GetY() + noteFontSize) // Move to the next line
pdf.Cell(nil, noteTextPart2)
  
   
                                        /*#########################################################*/           
                                        
                                        
                                    // Add margin line horizontally below the e-signature with background color #F48132
pdf.SetFillColor(244, 129, 50) // Set fill color to #F48132
//pdf.SetFillColor(6, 71, 71) // Set fill color to #064747
pdf.RectFromUpperLeftWithStyle(30, signatureYPos+signatureImageHeight+100, 535, 10, "F") // Draw a rectangle for the margin line

                                         /*############## GoGreen-Image - CONTENT ##############*/                                   
                                   
     
                                   /*############## GoGreen-Image - CONTENT ##############*/
                       

	goGreenImagePath := "images/goGreen.png" // Path to your GoGreen image
	goGreenImageWidth := 100.0                // Set the desired width of the GoGreen image
	goGreenImageHeight := 80.0               // Set the desired height of the GoGreen image
	goGreenXPos := 250.0                     // Adjust X position for the GoGreen image
	goGreenYPos := signatureYPos + signatureImageHeight + 100 // Adjust Y position for the GoGreen image

	err = pdf.Image(goGreenImagePath, goGreenXPos, goGreenYPos, &gopdf.Rect{W: goGreenImageWidth, H: goGreenImageHeight})
	if err != nil {
		fmt.Println("Error loading GoGreen image:", err)
	} else {
		fmt.Println("GoGreen image added successfully")
	}
	
	
	
	
	

                                   /*###########################################################*?
                                   
                                   
                                         
                                         /*############## FOOTER - CONTENT ##############*/

                  // Add footer background
                  pdf.SetFillColor(220, 220, 220) // Light gray color for the background
                  pdf.RectFromUpperLeftWithStyle(30, 800, 535, 20, "FD") // Draw a rectangle for the footer background

                  // Add footer text
                  footerFont := "Montserrat"
                  footerFontSize := 8.0
                  footerText := "Generated by EaseOfLearning"
                  pdf.SetFont(footerFont, "", footerFontSize)
                  pdf.SetTextColor(0, 0, 0) // Black color for footer text
                  pdf.SetX(250) // Adjust X position for the footer text
                  pdf.SetY(803) // Adjust Y position for the footer text (slightly lower for centering within the rectangle)
                  pdf.Cell(nil, footerText)

                                         /*########################################################*/
                                         
                                         
	err = pdf.WritePdf("invoice.pdf")
	if err != nil {
		panic(err)
	}

	fmt.Println("PDF generated Successfully")
                  }

