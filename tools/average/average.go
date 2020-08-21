package main
  
import (
         "fmt"
         "os"

         v "vision"
       )


var fp = fmt.Fprintf





func main ( ) {

  if len(os.Args) < 5 {
    fp ( os.Stdout, "\naverage error: need r g b out\n\n" )
    os.Exit ( 1 )
  }

  r_img := v.Read ( os.Args[1] )
  g_img := v.Read ( os.Args[2] )
  b_img := v.Read ( os.Args[3] )

  average_img := v.New_image ( v.Image_type_gray16, r_img.Width, r_img.Height )

  var sum_pixel uint64

  for y := uint32(0); y < r_img.Height; y ++ {
    for x := uint32(0); x < r_img.Width; x ++ {
      r := r_img.Get_gray16 ( x, y )
      g := g_img.Get_gray16 ( x, y )
      b := b_img.Get_gray16 ( x, y )
      sum_pixel = uint64(r) + uint64(g) + uint64(b)

      average_img.Set_gray16 ( x, y, uint16(sum_pixel/3) )
    }
  }
  
  average_img.Write ( os.Args[4] )
}





