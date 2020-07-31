package main
  
import (
         "fmt"
         "os"
         "strconv"

         v "vision"
       )


var fp = fmt.Fprintf


func check ( err error, msg string ) {
  if err != nil {
    fp ( os.Stdout, "list_pixels error : %s\n", msg )
    os.Exit ( 1 )
  }
}





func main ( ) {

  if len(os.Args) < 4 {
    fp ( os.Stdout, "\nlist_pixels error: need <FILE> <LOW> <HIGH>\n\n" )
    os.Exit ( 1 )
  }

  file_name := os.Args[1]

  low, err := strconv.Atoi ( os.Args[2] )
  check ( err, "Can't convert arg 2 to int." )

  high, err := strconv.Atoi ( os.Args[3] )
  check ( err, "Can't convert arg 3 to int." )

  img := v.Read ( file_name )

  var x, y uint32
  for y = 0; y < img.Height; y ++ {
    for x = 0; x < img.Width; x ++ {
      g := img.Get_gray16 ( x, y )
      if low <= int(g) && int(g) <= high {
        fp ( os.Stdout, "%d %d %d\n", x, y, g )
      }
    }
  }
}





