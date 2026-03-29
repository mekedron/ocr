#import <Foundation/Foundation.h>
#import <Vision/Vision.h>
#import <AppKit/AppKit.h>
#include <stdlib.h>
#include "ocr.h"

char* recognize_text(const char* image_path, const char** languages, int lang_count) {
    @autoreleasepool {
        NSString *path = [NSString stringWithUTF8String:image_path];
        NSImage *image = [[NSImage alloc] initWithContentsOfFile:path];
        if (!image) return NULL;

        NSData *tiffData = [image TIFFRepresentation];
        if (!tiffData) return NULL;

        NSBitmapImageRep *bitmap = [NSBitmapImageRep imageRepWithData:tiffData];
        if (!bitmap) return NULL;

        CGImageRef cgImage = [bitmap CGImage];
        if (!cgImage) return NULL;

        VNRecognizeTextRequest *request = [[VNRecognizeTextRequest alloc] init];
        request.recognitionLevel = VNRequestTextRecognitionLevelAccurate;

        if (lang_count > 0) {
            NSMutableArray *langs = [NSMutableArray arrayWithCapacity:lang_count];
            for (int i = 0; i < lang_count; i++) {
                [langs addObject:[NSString stringWithUTF8String:languages[i]]];
            }
            request.recognitionLanguages = langs;
        }

        VNImageRequestHandler *handler = [[VNImageRequestHandler alloc] initWithCGImage:cgImage options:@{}];
        NSError *error = nil;
        [handler performRequests:@[request] error:&error];
        if (error) return NULL;

        NSMutableString *result = [NSMutableString string];
        for (VNRecognizedTextObservation *observation in request.results) {
            VNRecognizedText *candidate = [[observation topCandidates:1] firstObject];
            if (candidate) {
                if (result.length > 0) [result appendString:@"\n"];
                [result appendString:candidate.string];
            }
        }

        if (result.length == 0) return NULL;
        return strdup([result UTF8String]);
    }
}

char* supported_languages(void) {
    @autoreleasepool {
        VNRecognizeTextRequest *request = [[VNRecognizeTextRequest alloc] init];
        request.recognitionLevel = VNRequestTextRecognitionLevelAccurate;

        NSError *error = nil;
        NSArray<NSString *> *languages = [request supportedRecognitionLanguagesAndReturnError:&error];
        if (error || !languages) return NULL;

        NSString *joined = [languages componentsJoinedByString:@"\n"];
        return strdup([joined UTF8String]);
    }
}
