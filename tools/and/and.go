package main
  
import (
         "fmt"
         "os"

         v "vision"
       )




func main ( ) {

  if len(os.Args) < 4 {
    panic ( fmt.Errorf ( "Need three file names." ) )
  }

  img_1  := v.Read ( os.Args[1] )
  img_2  := v.Read ( os.Args[2] )

  result_img := img_1.And_gray8 ( img_2 )

  result_img.Write ( os.Args[3] )
}





