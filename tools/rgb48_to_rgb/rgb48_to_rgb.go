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
    panic ( fmt.Errorf ( "usage : <RGB48_FILE> <RGBA_FILE>" ) )
  }

  rgb48_img := v.Read ( os.Args[1] )
  rgb_img   := v.New_image ( v.Image_type_rgba, rgb48_img.Width, rgb48_img.Height )

  for y := uint32(0); y < rgb48_img.Height; y ++ {
    for x := uint32(0); x < rgb48_img.Width; x ++ {
      r1, g1, b1 := rgb48_img.Get_rgb48 ( x, y )

      // fp ( os.Stdout, "MDEBUG pixel  x %d y %d : %d %d %d\n", x, y, r1, g1, b1 )

      biggest := r1
      if g1 > biggest {
        biggest = g1
      }
      if b1 > biggest {
        biggest = b1
      }

      /*
      if biggest < 100 {
        rgb_img.Set_rgba ( x, y, 0, 0, 0, 255 )
        continue
      }
      */

      multiplier := 65536.0/float64(biggest)

      // fp ( os.Stdout, "MDEBUG biggest: %d  multiplier %f\n", biggest, multiplier )

      r16 := multiplier * float64(r1)
      g16 := multiplier * float64(g1)
      b16 := multiplier * float64(b1)

      // fp ( os.Stdout, "MDEBUG 16:  %f %f %f \n", r16, g16, b16 )

      r8 := uint8 ( r16 / 256.0 )
      g8 := uint8 ( g16 / 256.0 )
      b8 := uint8 ( b16 / 256.0 )
      // fp ( os.Stdout, "MDEBUG 8:  %f %f %f \n", r8, g8, b8 )

      rgb_img.Set_rgba ( x, y, r8, g8, b8, 255 )
    }
  }

  rgb_img.Write_rgba_to_tif ( os.Args[2] )
}





