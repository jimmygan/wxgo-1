#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"


typedef struct {
	int _[3];
} SizerFlags;


void wxSizerFlags_Init(SizerFlags* p, int proportion);
void wxSizerFlags_Align(SizerFlags* p, int alignment);
void wxSizerFlags_Border(SizerFlags* p, int direction, int borderinpixels);
void wxSizerFlags_Bottom(SizerFlags* p);
void wxSizerFlags_Center(SizerFlags* p);
void wxSizerFlags_DoubleBorder(SizerFlags* p, int direction);
void wxSizerFlags_Expand(SizerFlags* p);
void wxSizerFlags_FixedMinSize(SizerFlags* p);
void wxSizerFlags_ReserveSpaceEvenIfHidden(SizerFlags* p);
void wxSizerFlags_Left(SizerFlags* p);
void wxSizerFlags_Proportion(SizerFlags* p, int proportion);
void wxSizerFlags_Right(SizerFlags* p);
void wxSizerFlags_Shaped(SizerFlags* p);
void wxSizerFlags_Top(SizerFlags* p);
void wxSizerFlags_TripleBorder(SizerFlags* p, int direction);

int wxSizerFlag_GetDefaultBorder();

#ifdef __cplusplus
}
#endif