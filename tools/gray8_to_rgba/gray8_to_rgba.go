package main
  
import (
         "fmt"
         "os"

         v "vision"
       )


var fp = fmt.Fprintf





func check ( err error ) {
  if err != nil {
    panic ( err )
  }
}





func main ( ) {

  if len(os.Args) < 3 {
    panic ( fmt.Errorf ( "usage : <GRAY8_FILE> <RGBA_FILE>" ) )
  }

  gray8_file_name := os.Args[1]
  rgb_file_name   := os.Args[2]

  gray_img := v.Read ( gray8_file_name )
  rgb_img  := v.New_image ( v.Image_type_rgba, gray_img.Width, gray_img.Height )

  gray_img.Invert_gray8 ( )

  for y := uint32(0); y < gray_img.Height; y ++ {
    for x := uint32(0); x < gray_img.Width; x ++ {
      g := gray_img.Get_gray8 ( x, y )
      rgb_img.Set_rgba ( x, y, g, 0, 0, 255 )
    }
  }

  rgb_img.Write_rgba_to_tif ( rgb_file_name )
}





