#ifndef __WXGO_CGOUTIL_H__
#define __WXGO_CGOUTIL_H__

#ifdef __cplusplus
extern "C" {
#endif

// An opaque hadle represents the intermedia between wxString and go string.
// The details of the handled data is known to "cgoutil.cpp" and "cgoutil.go"
// only.
// We only need to modify "cgoutil.cpp" and "cgoutil.go" in order to change the
// internal details of string exchanging, if we use this opaque handle to
// exchange string data.
typedef void* String;
typedef void* WxObjectPtr;
typedef int BOOL;
#define TRUE 1
#define FALSE 0

typedef struct tagPoint {
	int X, Y;
} Point;

typedef struct tagSize {
	int Width, Height;
} Size;

typedef struct tagRect {
	int X, Y, Width, Height;
} Rect;

extern BOOL g_CgoReady;


#ifdef __cplusplus
}
#endif

#endif //__WXGO_CGOUTIL_H__