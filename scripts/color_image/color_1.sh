#! /bin/bash

R_TIF=/home/annex/astronomy/Dark_Horse/2020/07/30/tif/r.tif
G_TIF=/home/annex/astronomy/Dark_Horse/2020/07/30/tif/g.tif
B_TIF=/home/annex/astronomy/Dark_Horse/2020/07/30/tif/b.tif

tif2whd ${R_TIF} ./r.whd
tif2whd ${G_TIF} ./g.whd
tif2whd ${B_TIF} ./b.whd

cp ${R_TIF} ./r.tif

~/bin/shift ./g.whd 1 -5 g_shifted.whd
~/bin/shift ./b.whd 2 -7 b_shifted.whd

whd2tif ./g_shifted.whd ./g_shifted.tif
whd2tif ./b_shifted.whd ./b_shifted.tif


