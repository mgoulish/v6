package main
  
import (
         "fmt"
         "os"

         v "vision"
       )


var fp = fmt.Fprintf





func main ( ) {

  if len(os.Args) < 2 {
    fp ( os.Stdout, "\nbrightest_point error: need file name\n\n" )
    os.Exit ( 1 )
  }

  file_name := os.Args[1]
  img := v.Read ( file_name )

  brightest_g := uint16(0)
  brightest_x := uint32(0)
  brightest_y := uint32(0)

  for y := uint32(0); y < img.Height; y ++ {
    for x := uint32(0); x < img.Height; x ++ {
      g := img.Get_gray16 ( x, y )

      if g > brightest_g {
        brightest_g = g
        brightest_x = x
        brightest_y = y
      }
    }
  }
  
  fp ( os.Stdout, "\n\n%d at %d %d\n", brightest_g, brightest_x, brightest_y )
}





