
package vision



import (
        "fmt"
        "image"
        "os"

        tif "golang.org/x/image/tiff"
)





func ( img * Image ) Set_gray8 ( x, y uint32, g uint8 ) {
  address := x + y * img.Width 
  img.Pixels [ address ] = g
}





func ( img * Image ) Get_gray8 ( x, y uint32 ) ( g  uint8 ) {
  address := x + y * img.Width
  val := img.Pixels [ address ]
  return val
}





func (img * Image) Write_gray8_to_tif ( file_name string ) {

  output_file, err := os.Create ( file_name )
  if err != nil {
    panic ( fmt.Errorf ( "image_write error: |%s|\n", err.Error()))
  }
  defer output_file.Close()

  output_image := image.NewRGBA ( image.Rect(0, 0, int(img.Width), int(img.Height)) )

  for y := uint32(0); y < uint32(img.Height); y ++ {
    for x := uint32(0); x < uint32(img.Width); x ++ {
      g := img.Get_gray8 ( x, y )
      addr := 4 * ( (y * img.Width) + x )
      output_image.Pix [ addr     ] = g
      output_image.Pix [ addr + 1 ] = g
      output_image.Pix [ addr + 2 ] = g
      output_image.Pix [ addr + 3 ] = 255
    }
  }

  tif.Encode ( output_file, output_image, nil )
}





func ( img_1 * Image ) And_gray8 ( img_2 * Image ) ( result * Image ) {
  
  if img_1.Image_type != Image_type_gray8 || img_2.Image_type != Image_type_gray8 {
    panic ( "Images must both be gray8." )
  }

  if img_1.Height != img_2.Height || img_1.Width != img_2.Width {
    panic ( fmt.Errorf ( "Both images must be same size." ) )
  }

  result = New_image ( Image_type_gray8, img_1.Width, img_1.Height )

  var x, y uint32
  for y = 0; y < img_1.Height; y ++ {
    for x = 0; x < img_1.Width; x ++ {
      if img_1.Get_gray8(x, y) == 255 && img_2.Get_gray8(x, y) == 255 {
        result.Set_gray8 ( x, y, 255 )
      } else {
        result.Set_gray8 ( x, y, 0 )
      }
    }
  }

  return result
}





