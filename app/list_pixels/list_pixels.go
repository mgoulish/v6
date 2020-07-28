package main
  
import (
         "fmt"
         "os"
         "strconv"

         v "vision"
       )


var fp = fmt.Fprintf


func main ( ) {

  if len(os.Args) < 4 {
    fp ( os.Stdout, "\nlist_pixels error: need <FILE> <LOW> <HIGH>\n\n" )
    os.Exit ( 1 )
  }

  file_name := os.Args[1]

  low, err := strconv.Atoi ( os.Args[2] )
  if err != nil {
    fp ( os.Stdout, "list_pixels err: can't convert arg 2 to int.\n" )
    os.Exit ( 1 )
  }

  high, err := strconv.Atoi ( os.Args[3] )
  if err != nil {
    fp ( os.Stdout, "list_pixels err: can't convert arg 3 to int.\n" )
    os.Exit ( 1 )
  }

  img := v.Read_Image_Gray16 ( file_name )

  var x, y uint32
  for y = 0; y < img.Height; y ++ {
    for x = 0; x < img.Width; x ++ {
      g := img.Get_Gray16 ( x, y )
      if low <= int(g) && int(g) <= high {
        fp ( os.Stdout, "%d %d %d\n", x, y, g )
      }
    }
  }
}





