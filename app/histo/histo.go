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

  img := v.Read_Image_Gray16 ( file_name )

  n := 65536

  histo := make ( []int, n )

  var x, y uint32
  for y = 0; y < img.Height; y ++ {
    for x = 0; x < img.Width; x ++ {
      g := img.Get_Gray16 ( x, y )
      histo [ g ] ++
    }
  }

  for i := 0; i < n; i ++ {
    if histo[i] > 0 {
      fp ( os.Stdout, "%8d  %d\n", i, histo[i] )
    }
  }
}





