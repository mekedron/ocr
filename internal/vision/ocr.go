package vision

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Vision -framework AppKit -framework Foundation
#include <stdlib.h>
#include "ocr.h"
*/
import "C"
import "unsafe"

// RecognizeText runs OCR on the image at the given path using Apple Vision framework.
func RecognizeText(imagePath string, languages []string) string {
	cPath := C.CString(imagePath)
	defer C.free(unsafe.Pointer(cPath))

	var cLangs **C.char
	cLangSlice := make([]*C.char, len(languages))
	for i, lang := range languages {
		cLangSlice[i] = C.CString(lang)
	}
	defer func() {
		for _, p := range cLangSlice {
			C.free(unsafe.Pointer(p))
		}
	}()
	if len(cLangSlice) > 0 {
		cLangs = &cLangSlice[0]
	}

	result := C.recognize_text(cPath, cLangs, C.int(len(languages)))
	if result == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}

// SupportedLanguages returns a newline-separated list of supported language codes.
func SupportedLanguages() string {
	result := C.supported_languages()
	if result == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(result))
	return C.GoString(result)
}
