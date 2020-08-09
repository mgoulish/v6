package main
  
import (
         "fmt"
         "os"
         "strconv"

         v "vision"
       )




func main ( ) {

  if len(os.Args) < 4 {
    panic ( fmt.Errorf ( "Need src thresh dst." ))
  }

  thresh, err := strconv.Atoi ( os.Args[2] )
  if err != nil {
    panic ( fmt.Errorf ( "Can't convert |%s| to int.", os.Args[2] ) )
  }

  img := v.Read ( os.Args[1] )
  result_img := img.Threshold_gray16 ( uint16(thresh) ) 
  result_img.Write ( os.Args[3] )
  result_img.Write_gray8_to_tif ( "./result.tif" )
}





