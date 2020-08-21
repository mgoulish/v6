package main
  
import (
         "fmt"
         "os"
         "strconv"

         v "vision"
       )


var fp = fmt.Fprintf





func main ( ) {

  if len(os.Args) < 6 {
    fp ( os.Stdout, "\nlist_pixels error: need file name, x, y, w, h\n\n" )
    os.Exit ( 1 )
  }

  file_name := os.Args[1]
  img := v.Read ( file_name )

  x, _ := strconv.Atoi ( os.Args[2] )
  y, _ := strconv.Atoi ( os.Args[3] )
  w, _ := strconv.Atoi ( os.Args[4] )
  h, _ := strconv.Atoi ( os.Args[5] )

  fp ( os.Stdout, "MDEBUG %d %d %d %d\n", x, y, w, h )

  for y1 := uint32(y); y1 < uint32(y+h); y1 ++ {
    for x1 := uint32(x); x1 < uint32(x+w); x1 ++ {
      g := img.Get_gray16 ( x1, y1 )

      fp ( os.Stdout, "%d %d %d\n", x1, y1, g )
    }
  }
}





