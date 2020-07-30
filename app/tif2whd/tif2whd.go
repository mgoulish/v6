package main
  
import (
         "bytes"
         "encoding/binary"
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
    fp ( os.Stdout, "\ntif2whd error: usage : tif2whd <TIFF_FILE> <WHD_FILE>\n\n" )
    os.Exit ( 1 )
  }

  tif_file_name := os.Args[1]
  whd_file_name := os.Args[2]

  img := v.Read_tif_to_gray16 ( tif_file_name )

  f, err := os.Create ( whd_file_name )
  check ( err )
  defer f.Close ( )

  var buf bytes.Buffer
  err = binary.Write ( & buf, 
                       binary.BigEndian, 
                       []uint32{img.Image_type, img.Width, img.Height} )
  check ( err )
  err = binary.Write ( & buf,
                       binary.BigEndian,
                       img.Pixels )
  _, err = f.Write ( buf.Bytes() )
  check ( err )
}





