package main
  
import (
         "fmt"
         "os"

         v "vision"
       )


var fp = fmt.Fprintf





func check ( err error ) {
  if err != nil {
    panic ( err )
  }
}





func main ( ) {

  if len(os.Args) < 3 {
    panic ( fmt.Errorf ( "usage : <GRAY16> <GRAY16> " ) )
  }

  thresh := uint16(15)

  gray16_1 := os.Args[1]
  gray16_2 := os.Args[2]

  // Read in the gray16 from the original tif,
  // and apply the threshold.
  gray_img_1 := v.Read_tif_to_gray16 ( gray16_1 )
  gray_img_2 := v.Read_tif_to_gray16 ( gray16_2 )

  thresh_1 := gray_img_1.Threshold_gray16 ( thresh )
  thresh_2 := gray_img_2.Threshold_gray16 ( thresh )

  // Invert the images, because the below-threshold pixels
  // are easier to see that way.
  thresh_1.Invert_gray8 ( )
  thresh_2.Invert_gray8 ( )

  anded := thresh_1.And_gray8 ( thresh_2 )

  dilated  := anded.Dilate ( )
  dilated2 := dilated.Dilate ( )
  dilated3 := dilated2.Dilate ( )

  dilated3.Write_gray8_to_tif ( "./result.tif" )
}





