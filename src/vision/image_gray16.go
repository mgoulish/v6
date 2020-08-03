
package vision



import (
        "fmt"
        "image"
        "image/color"
        "os"

        "golang.org/x/image/tiff"
)





func ( img * Image ) Set_gray16 ( x, y uint32, g uint16 ) {
  address := 2 * ( x + y * img.Width )

  if address >= img.n_bytes - 1 {
    panic ( fmt.Errorf ( "Set_gray16: %d %d is not in image. (%d %d)", 
                         x, y, img.Width, img.Height ) )
  }

  img.Pixels [ address     ] = byte(g >> 8)
  img.Pixels [ address + 1 ] = byte(g&0x0F)
}





func ( img * Image ) Get_gray16 ( x, y uint32 ) ( g  uint16 ) {
  address := 2 * ( x + y * img.Width )

  if address >= img.n_bytes - 1 {
    panic ( fmt.Errorf ( "Get_gray16: %d %d is not in image. (%d %d)", 
                         x, y, img.Width, img.Height ) )
  }

  val := uint16(img.Pixels [ address ])
  val <<= 8
  val += uint16(img.Pixels [ address + 1])
  return val
}





func Read_tif_to_gray16 ( file_name string ) ( img * Image ) {
  reader, err := os.Open ( file_name )
  if err != nil {
    fp ( os.Stdout, "image_read error: |%s|\n", err.Error() )
    os.Exit ( 1 )
  }
  defer reader.Close()

  decoded_image, _, err := image.Decode ( reader )
  if err != nil {
    fp ( os.Stdout, "image_read error: |%s|\n", err.Error() )
    os.Exit ( 1 )
  }

  gray_image, ok := decoded_image.(*image.Gray16)
  if ! ok {
    fp ( os.Stdout, "image_read error: Can't get gray16 image.\n" )
    os.Exit ( 1 )
  }

  bounds := gray_image.Bounds()
  width  := uint32(bounds.Max.X)
  height := uint32(bounds.Max.Y)

  img = New_image ( Image_type_gray16, width, height )

  var x, y uint32
  for y = 0; y < uint32(bounds.Max.Y); y ++ {
    for x = 0; x < uint32(bounds.Max.X); x ++ {
      g := gray_image.Gray16At ( int(x), int(y) )
      img.Set_gray16 ( x, y, g.Y )
    }
  }

  return img
}





func ( img * Image ) Write_gray16_to_tif ( file_name string ) {
  fn := "Write_gray16_to_tif"

  width  := int(img.Width)
  height := int(img.Height)
  rect   := image.Rectangle { image.Point{0,0}, image.Point{width, height} }

  tif_image := image.NewGray16 ( rect )

  var x, y uint32
  for y = 0; y < img.Height; y ++ {
    for x = 0; x < img.Width; x ++ {
      g := img.Get_gray16 ( x, y )
      tif_image.SetGray16 ( int(x), int(y), color.Gray16 { g } )
    }
  }

  tif_file, err := os.Create ( file_name )
  check ( err, fn )

  err = tiff.Encode ( tif_file, tif_image, nil )
  check ( err, fn )
}





func ( img * Image ) Threshold_gray16 ( threshold uint16 ) ( dst * Image ) {

  // Gray8 is my equivalent of a binary image.
  dst = New_image ( Image_type_gray8, img.Width, img.Height )

  var x, y uint32
  for y = 0; y < img.Height; y ++ {
    for x = 0; x < img.Width; x ++ {
      if img.Get_gray16 ( x, y ) >= threshold {
        dst.Set_gray8 ( x, y, 255 )
      } else {
        dst.Set_gray8 ( x, y, 0 )
      }
    }
  }

  return dst
}





func ( img * Image ) Shift_X_Gray16 ( shift int ) ( result * Image ) {

  result = New_image ( Image_type_gray16, img.Width, img.Height )

  // Zero out all pixels.
  var x, y uint32
  for y = 0; y < img.Height; y ++ {
    for x = 0; x < img.Width; x ++ {
        result.Set_gray16 ( x, y, 0 )
    }
  }

  if shift < 0 {
    for y = 0; y < img.Height; y ++ {
      dst_x := uint32(0)
      for x = uint32(-shift); x < img.Width; x ++ {
          g := img.Get_gray16 ( x, y )
          result.Set_gray16 ( dst_x, y, g )
          dst_x ++
      }
    }
  } else
  {
    for y = 0; y < img.Height; y ++ {
      dst_x := uint32(shift)
      offset := int(img.Width) - shift
      for x = 0; x < uint32(offset); x ++ {
          g := img.Get_gray16 ( x, y )
          result.Set_gray16 ( dst_x, y, g )
          dst_x ++
      }
    }
  }

  return result
}



func ( img * Image ) Shift_Y_Gray16 ( shift int ) ( result * Image ) {
  result = New_image ( Image_type_gray16, img.Width, img.Height )

  // Zero out all pixels.
  var x, y uint32
  for y = 0; y < img.Height; y ++ {
    for x = 0; x < img.Width; x ++ {
      result.Set_gray16 ( x, y, 0 )
    }
  }

  if shift < 0 {
    dst_y := uint32(0)
    for y = uint32(-shift); y < img.Height - uint32(-shift); y ++ {
      for x = 0; x < img.Width; x ++ {
        g := img.Get_gray16 ( x, y )
        result.Set_gray16 ( x, dst_y, g )
      }
      dst_y ++
    }
  } else {
    dst_y := uint32(shift)
    for y = 0; y < img.Height - uint32(shift); y ++ {
      for x = 0; x < img.Width; x ++ {
        g := img.Get_gray16 ( x, y )
        result.Set_gray16 ( x, dst_y, g )
      }
      dst_y ++
    }
  }

  return result
}





