package main
  
import (
         "fmt"
         "os"

         v "vision"
       )


var fp = fmt.Fprintf





func main ( ) {

  if len(os.Args) < 5 {
    fp ( os.Stdout, "\nrgb48 error: need r g b out\n\n" )
    os.Exit ( 1 )
  }

  r_img := v.Read ( os.Args[1] )
  g_img := v.Read ( os.Args[2] )
  b_img := v.Read ( os.Args[3] )

  rgb48 := v.New_image ( v.Image_type_rgb48, r_img.Width, r_img.Height )

  for y := uint32(0); y < r_img.Height; y ++ {
    for x := uint32(0); x < r_img.Width; x ++ {
      r := r_img.Get_gray16 ( x, y )
      g := g_img.Get_gray16 ( x, y )
      b := b_img.Get_gray16 ( x, y )
      rgb48.Set_rgb48 ( x, y, r, g, b )
    }
  }
  
  rgb48.Write ( os.Args[4] )
}





