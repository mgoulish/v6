package main
  
import (
         "fmt"
         "os"

         v "vision"
       )


var fp = fmt.Fprintf


func main ( ) {

  if len(os.Args) < 2 {
    fp ( os.Stdout, "\nhisto error: need file name.\n\n" )
    os.Exit ( 1 )
  }

  file_name := os.Args[1]

  img := v.Read ( file_name )

  fp ( os.Stdout, 
       "MDEBUG img type %s, width %d, height %d\n",
       v.Image_type_name ( img.Image_type ),
       img.Width,
       img.Height,
     )
  
  var x, y uint32

  x = 1721
  y =  902

  g := img.Get_gray16 ( x, y )

  fp ( os.Stdout, "MDEBUG pixel at %d, %d is %d\n", x, y, g )
}





