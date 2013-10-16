#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

typedef struct {
	int Width, Height;
} Size;

typedef struct {
	int X, Y;
} Point;

typedef struct {
	int X, Y, Width, Height;
} Rect;

Size Get_wxDefaultSize();
Point Get_wxDefaultPosition();


#ifdef __cplusplus
}
#endif