package ui

import (
	"image/color"
)

func init() {
	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(Colors), func(i, j int) { Colors[i], Colors[j] = Colors[j], Colors[i] })
}

var Colors []color.RGBA = []color.RGBA{
	{128, 0, 0, 0xFF},     // maroon    #800000
	{139, 0, 0, 0xFF},     // dark red  #8B0000
	{165, 42, 42, 0xFF},   // brown   #A52A2A
	{178, 34, 34, 0xFF},   // firebrick   #B22222
	{220, 20, 60, 0xFF},   // crimson     #DC143C
	{255, 0, 0, 0xFF},     // red   #FF0000
	{255, 99, 71, 0xFF},   // tomato  #FF6347
	{255, 127, 80, 0xFF},  // coral  #FF7F50
	{205, 92, 92, 0xFF},   // indian red  #CD5C5C
	{240, 128, 128, 0xFF}, // light coral   #F08080
	{233, 150, 122, 0xFF}, // dark salmon   #E9967A
	{250, 128, 114, 0xFF}, // salmon    #FA8072
	{255, 160, 122, 0xFF}, // light salmon  #FFA07A
	{255, 69, 0, 0xFF},    // orange red   #FF4500
	{255, 140, 0, 0xFF},   // dark orange     #FF8C00
	{255, 165, 0, 0xFF},   // orange  #FFA500
	{255, 215, 0, 0xFF},   // gold    #FFD700
	{184, 134, 11, 0xFF},  // dark golden rod    #B8860B
	{218, 165, 32, 0xFF},  // golden rod     #DAA520
	{238, 232, 170, 0xFF}, // pale golden rod   #EEE8AA
	{189, 183, 107, 0xFF}, // dark khaki    #BDB76B
	{240, 230, 140, 0xFF}, // khaki     #F0E68C
	{128, 128, 0, 0xFF},   // olive   #808000
	{255, 255, 0, 0xFF},   // yellow  #FFFF00
	{154, 205, 50, 0xFF},  // yellow green   #9ACD32
	{85, 107, 47, 0xFF},   // dark olive green    #556B2F
	{107, 142, 35, 0xFF},  // olive drab     #6B8E23
	{124, 252, 0, 0xFF},   // lawn green  #7CFC00
	{127, 255, 0, 0xFF},   // chart reuse     #7FFF00
	{173, 255, 47, 0xFF},  // green yellow   #ADFF2F
	{0, 100, 0, 0xFF},     // dark green    #006400
	{0, 128, 0, 0xFF},     // green     #008000
	{34, 139, 34, 0xFF},   // forest green    #228B22
	{0, 255, 0, 0xFF},     // lime  #00FF00
	{50, 205, 50, 0xFF},   // lime green  #32CD32
	{144, 238, 144, 0xFF}, // light green   #90EE90
	{152, 251, 152, 0xFF}, // pale green    #98FB98
	{143, 188, 143, 0xFF}, // dark sea green    #8FBC8F
	{0, 250, 154, 0xFF},   // medium spring green     #00FA9A
	{0, 255, 127, 0xFF},   // spring green    #00FF7F
	{46, 139, 87, 0xFF},   // sea green   #2E8B57
	{102, 205, 170, 0xFF}, // medium aqua marine    #66CDAA
	{60, 179, 113, 0xFF},  // medium sea green   #3CB371
	{32, 178, 170, 0xFF},  // light sea green    #20B2AA
	{47, 79, 79, 0xFF},    // dark slate gray  #2F4F4F
	{0, 128, 128, 0xFF},   // teal    #008080
	{0, 139, 139, 0xFF},   // dark cyan   #008B8B
	{0, 255, 255, 0xFF},   // aqua    #00FFFF
	{0, 255, 255, 0xFF},   // cyan    #00FFFF
	{224, 255, 255, 0xFF}, // light cyan    #E0FFFF
	{0, 206, 209, 0xFF},   // dark turquoise  #00CED1
	{64, 224, 208, 0xFF},  // turquoise  #40E0D0
	{72, 209, 204, 0xFF},  // medium turquoise   #48D1CC
	{175, 238, 238, 0xFF}, // pale turquoise    #AFEEEE
	{127, 255, 212, 0xFF}, // aqua marine   #7FFFD4
	{176, 224, 230, 0xFF}, // powder blue   #B0E0E6
	{95, 158, 160, 0xFF},  // cadet blue     #5F9EA0
	{70, 130, 180, 0xFF},  // steel blue     #4682B4
	{100, 149, 237, 0xFF}, // corn flower blue  #6495ED
	{0, 191, 255, 0xFF},   // deep sky blue   #00BFFF
	{30, 144, 255, 0xFF},  // dodger blue    #1E90FF
	{173, 216, 230, 0xFF}, // light blue    #ADD8E6
	{135, 206, 235, 0xFF}, // sky blue  #87CEEB
	{135, 206, 250, 0xFF}, // light sky blue    #87CEFA
	{25, 25, 112, 0xFF},   // midnight blue   #191970
	{0, 0, 128, 0xFF},     // navy  #000080
	{0, 0, 139, 0xFF},     // dark blue     #00008B
	{0, 0, 205, 0xFF},     // medium blue   #0000CD
	{0, 0, 255, 0xFF},     // blue  #0000FF
	{65, 105, 225, 0xFF},  // royal blue     #4169E1
	{138, 43, 226, 0xFF},  // blue violet    #8A2BE2
	{75, 0, 130, 0xFF},    // indigo   #4B0082
	{72, 61, 139, 0xFF},   // dark slate blue     #483D8B
	{106, 90, 205, 0xFF},  // slate blue     #6A5ACD
	{123, 104, 238, 0xFF}, // medium slate blue     #7B68EE
	{147, 112, 219, 0xFF}, // medium purple     #9370DB
	{139, 0, 139, 0xFF},   // dark magenta    #8B008B
	{148, 0, 211, 0xFF},   // dark violet     #9400D3
	{153, 50, 204, 0xFF},  // dark orchid    #9932CC
	{186, 85, 211, 0xFF},  // medium orchid  #BA55D3
	{128, 0, 128, 0xFF},   // purple  #800080
	{216, 191, 216, 0xFF}, // thistle   #D8BFD8
	{221, 160, 221, 0xFF}, // plum  #DDA0DD
	{238, 130, 238, 0xFF}, // violet    #EE82EE
	{255, 0, 255, 0xFF},   // magenta / fuchsia   #FF00FF
	{218, 112, 214, 0xFF}, // orchid    #DA70D6
	{199, 21, 133, 0xFF},  // medium violet red  #C71585
	{219, 112, 147, 0xFF}, // pale violet red   #DB7093
	{255, 20, 147, 0xFF},  // deep pink  #FF1493
	{255, 105, 180, 0xFF}, // hot pink  #FF69B4
	{255, 182, 193, 0xFF}, // light pink    #FFB6C1
	{255, 192, 203, 0xFF}, // pink  #FFC0CB
	{250, 235, 215, 0xFF}, // antique white     #FAEBD7
	{245, 245, 220, 0xFF}, // beige     #F5F5DC
	{255, 228, 196, 0xFF}, // bisque    #FFE4C4
	{255, 235, 205, 0xFF}, // blanched almond   #FFEBCD
	{245, 222, 179, 0xFF}, // wheat     #F5DEB3
	{255, 248, 220, 0xFF}, // corn silk     #FFF8DC
	{255, 250, 205, 0xFF}, // lemon chiffon     #FFFACD
	{250, 250, 210, 0xFF}, // light golden rod yellow   #FAFAD2
	{255, 255, 224, 0xFF}, // light yellow  #FFFFE0
	{139, 69, 19, 0xFF},   // saddle brown    #8B4513
	{160, 82, 45, 0xFF},   // sienna  #A0522D
	{210, 105, 30, 0xFF},  // chocolate  #D2691E
	{205, 133, 63, 0xFF},  // peru   #CD853F
	{244, 164, 96, 0xFF},  // sandy brown    #F4A460
	{222, 184, 135, 0xFF}, // burly wood    #DEB887
	{210, 180, 140, 0xFF}, // tan   #D2B48C
	{188, 143, 143, 0xFF}, // rosy brown    #BC8F8F
	{255, 228, 181, 0xFF}, // moccasin  #FFE4B5
	{255, 222, 173, 0xFF}, // navajo white  #FFDEAD
	{255, 218, 185, 0xFF}, // peach puff    #FFDAB9
	{255, 228, 225, 0xFF}, // misty rose    #FFE4E1
	{255, 240, 245, 0xFF}, // lavender blush    #FFF0F5
	{250, 240, 230, 0xFF}, // linen     #FAF0E6
	{253, 245, 230, 0xFF}, // old lace  #FDF5E6
	{255, 239, 213, 0xFF}, // papaya whip   #FFEFD5
	{255, 245, 238, 0xFF}, // sea shell     #FFF5EE
	{245, 255, 250, 0xFF}, // mint cream    #F5FFFA
	{112, 128, 144, 0xFF}, // slate gray    #708090
	{119, 136, 153, 0xFF}, // light slate gray  #778899
	{176, 196, 222, 0xFF}, // light steel blue  #B0C4DE
	{230, 230, 250, 0xFF}, // lavender  #E6E6FA
	{255, 250, 240, 0xFF}, // floral white  #FFFAF0
	{240, 248, 255, 0xFF}, // alice blue    #F0F8FF
	{248, 248, 255, 0xFF}, // ghost white   #F8F8FF
	{240, 255, 240, 0xFF}, // honeydew  #F0FFF0
	{255, 255, 240, 0xFF}, // ivory     #FFFFF0
	{240, 255, 255, 0xFF}, // azure     #F0FFFF
	{255, 250, 250, 0xFF}, // snow  #FFFAFA
	{0, 0, 0, 0xFF},       // black   #000000
	{105, 105, 105, 0xFF}, // dim gray / dim grey   #696969
	{128, 128, 128, 0xFF}, // gray / grey   #808080
	{169, 169, 169, 0xFF}, // dark gray / dark grey     #A9A9A9
	{192, 192, 192, 0xFF}, // silver    #C0C0C0
	{211, 211, 211, 0xFF}, // light gray / light grey   #D3D3D3
	{220, 220, 220, 0xFF}, // gainsboro     #DCDCDC
	{245, 245, 245, 0xFF}, // white smoke   #F5F5F5
	{255, 255, 255, 0xFF}, // white     #FFFFFF
}
