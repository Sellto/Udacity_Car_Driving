package action

import (
  "../lib"
)


var HOG64 = cp.HOGdata {
      ResizeX     : 64,
      ResizeY     : 64,
      CellSize    : 8, // [pixel]
      BlockSize   : 2, // [cell]
      Epsilon     : 1.0,
      Orientation : 9,
    }

var HOG32 = cp.HOGdata {
      ResizeX     : 32,
      ResizeY     : 32,
      CellSize    : 4, // [pixel]
      BlockSize   : 2, // [cell]
      Epsilon     : 1.0,
      Orientation : 9,
}

var HOGALL = cp.HOGdata {
      ResizeX     : 320,
      ResizeY     : 48,
      CellSize    : 8, // [pixel]
      BlockSize   : 2, // [cell]
      Epsilon     : 1.0,
      Orientation : 6,
}

var Available_config = []string{"hog64", "hog32", "hogALL", "gs", "customgs","bw"}
