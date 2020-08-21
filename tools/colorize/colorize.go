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





func choose_color ( gray_val uint16 ) ( r, g, b uint16 ) {
  color_map := [][]uint16{
    // threshold, r, g, b
    {500,  128,   0,   0},      // dark red
    {1000,  255,   0,   0},      // red
    {5000,  255, 130,   0},      // orange
    {10000,  228, 255,   0},      // yellow
    {20000,    0, 255,   0},      // green
    {30000,    0,   0, 255},      // blue
    {40000,  238, 130, 238},      // violet
    {50000, 255, 255, 255},       // white
  }

  for _, c := range color_map {
    if gray_val >= c[0] {
      r = c[1]
      g = c[2]
      b = c[3]
    }
  }

  return r, g, b
}





func main ( ) {

  if len(os.Args) < 3 {
    panic ( fmt.Errorf ( "usage : <GRAY16_FILE> <RGBA_FILE>" ) )
  }

  gray16_file_name := os.Args[1]
  rgb_file_name   := os.Args[2]

  gray_img := v.Read ( gray16_file_name )
  rgb_img  := v.New_image ( v.Image_type_rgba, gray_img.Width, gray_img.Height )

  for y := uint32(0); y < gray_img.Height; y ++ {
    for x := uint32(0); x < gray_img.Width; x ++ {
      g := gray_img.Get_gray16 ( x, y )
      r, g, b := choose_color ( g )
      rgb_img.Set_rgba ( x, y, uint8(r), uint8(g), uint8(b), 255 )
      
    }
  }

  rgb_img.Write_rgba_to_tif ( rgb_file_name )
}





