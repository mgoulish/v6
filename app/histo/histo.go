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
  
  n := 65536
  histo := make ( []int, n )
  var x, y uint32
  for y = 0; y < img.Height; y ++ {
    for x = 0; x < img.Width; x ++ {
      g := img.Get_gray16 ( x, y )
      histo [ g ] ++
    }
  }

  for i := 0; i < n; i ++ {
    if histo[i] > 0 {
      fp ( os.Stdout, "%8d  %d\n", i, histo[i] )
    }
  }
}





