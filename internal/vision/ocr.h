#ifndef OCR_H
#define OCR_H

char* recognize_text(const char* image_path, const char** languages, int lang_count);
char* supported_languages(void);

#endif
