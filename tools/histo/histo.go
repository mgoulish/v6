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
    fp ( os.Stdout, "\nhisto error: need file name, x, y, w, h.\n\n" )
    os.Exit ( 1 )
  }

  file_name := os.Args[1]
  img := v.Read ( file_name )

  x, _ := strconv.Atoi ( os.Args[2] )
  y, _ := strconv.Atoi ( os.Args[3] )
  w, _ := strconv.Atoi ( os.Args[4] )
  h, _ := strconv.Atoi ( os.Args[5] )

  histo := img.Histogram_gray16 ( uint32(x), uint32(y), uint32(w), uint32(h) )

  //last_nonzero := 0
  for i := 0; i < len(histo); i ++ {
    if histo[i] > 0 {
      fp ( os.Stdout, "%8d  %d\n", i, histo[i] )
      /*
      gap := i - last_nonzero
      if gap > 10 {
        fp ( os.Stdout, "gap ending at %d, size %d\n", i, gap )
      }
      last_nonzero = i
      */
    }
  }
}





