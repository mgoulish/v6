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

  if len(os.Args) < 5 {
    fp ( os.Stdout, "\nlist_pixels error: need <FILE_1> <FILE_2>  <LOW> <HIGH>\n\n" )
    os.Exit ( 1 )
  }

  file_1 := os.Args[1]
  file_2 := os.Args[2]

  low, err := strconv.Atoi ( os.Args[3] )
  check ( err, "Can't convert arg 3 to int." )

  high, err := strconv.Atoi ( os.Args[4] )
  check ( err, "Can't convert arg 4 to int." )

  img_1 := v.Read ( file_1 )
  img_2 := v.Read ( file_2 )

  var x, y uint32
  for y = 0; y < img_1.Height; y ++ {
    for x = 0; x < img_1.Width; x ++ {
      g_1 := img_1.Get_gray16 ( x, y )
      g_2 := img_2.Get_gray16 ( x, y )

      if low <= int(g_1) && int(g_1) <= high && low <= int(g_2) && int(g_2) <= high {
        fp ( os.Stdout, "%d %d %d %d\n", x, y, g_1, g_2 )
      }
    }
  }
}





