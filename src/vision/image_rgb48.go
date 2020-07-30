
package vision



import (
        _ "golang.org/x/image/tiff"
)





func ( img * Image ) Set_RGB48 ( x, y uint32, r, g, b uint16 ) {
  bpp := Bytes_per_pixel ( img.Image_type )
  address := bpp * ( x + y * img.Width )

  // r ---------------------
  img.Pixels [ address     ] = byte(r >> 8)
  img.Pixels [ address + 1 ] = byte(r&0x0F)

  // g ---------------------
  img.Pixels [ address + 2 ] = byte(g >> 8)
  img.Pixels [ address + 3 ] = byte(g&0x0F)

  // b ---------------------
  img.Pixels [ address + 4 ] = byte(b >> 8)
  img.Pixels [ address + 5 ] = byte(b&0x0F)
}





func ( img * Image ) Get_RGB48 ( x, y uint32 ) ( r, g, b  uint16 ) {
  bpp := Bytes_per_pixel ( img.Image_type )
  address := bpp * ( x + y * img.Width )

  // r ---------------------
  r = uint16(img.Pixels [ address ])
  r <<= 8
  r += uint16(img.Pixels [ address + 1])

  // g ---------------------
  g = uint16(img.Pixels [ address ])
  g <<= 8
  g += uint16(img.Pixels [ address + 1])

  // b ---------------------
  b = uint16(img.Pixels [ address ])
  b <<= 8
  b += uint16(img.Pixels [ address + 1])

  return r, g, b
}





