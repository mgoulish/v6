package main
  
import (
         "fmt"
         "os"
         "strconv"

         v "vision"
       )




func main ( ) {

  if len(os.Args) < 5 {
    panic ( fmt.Errorf ( "Need src x-shift y-shift dst" ) )
  }

  src := os.Args[1] 
  dst := os.Args[4] 

  x_shift, err := strconv.Atoi ( os.Args[2] )
  if err != nil {
    panic ( fmt.Errorf ( "Can't convert |%s| to int.", os.Args[2] ) )
  }

  y_shift, err := strconv.Atoi ( os.Args[3] )
  if err != nil {
    panic ( fmt.Errorf ( "Can't convert |%s| to int.", os.Args[3] ) )
  }

  img := v.Read ( src )
  x_shifted := img.Shift_X_Gray16 ( x_shift ) 
  x_and_y_shifted := x_shifted.Shift_Y_Gray16 ( y_shift ) 
  x_and_y_shifted.Write ( dst )
}





