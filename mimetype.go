package mimetype

import (
	"fmt"
	"strings"
)

// ByExtension will take in an extension, with or without the leading "." and will return the mimetype for it if we have one.
func ByExtension(ext string) (string, error) {

	if strings.HasPrefix(ext, ".") {
		ext = strings.TrimPrefix(ext, ".")

	}

	mimeTypes := map[string]string{
		// Image
		"jfi": "image/jpeg",
		"jfif": "image/jpeg",
		"jpeg": "image/jpeg",
		"jpg":  "image/jpeg",
		"png":  "image/png",
		"gif":  "image/gif",
		"webp": "image/webp",
		"cr2":  "image/x-canon-cr2",
		"tif":  "image/tiff",
		"bmp":  "image/bmp",
		"heif": "image/heif",
		"jxr":  "image/vnd.ms-photo",
		"psd":  "image/vnd.adobe.photoshop",
		"ico":  "image/vnd.microsoft.icon",
		"dwg":  "image/vnd.dwg",
		// Video
		"mp4":  "video/mp4",
		"m4v":  "video/x-m4v",
		"mkv":  "video/x-matroska",
		"webm": "video/webm",
		"mov":  "video/quicktime",
		"avi":  "video/x-msvideo",
		"wmv":  "video/x-ms-wmv",
		"mpg":  "video/mpeg",
		"mpeg": "video/mpeg",
		"flv":  "video/x-flv",
		"3gp":  "video/3gpp",
		// Audio
		"mid":  "audio/midi",
		"mp3":  "audio/mpeg",
		"m4a":  "audio/m4a",
		"ogg":  "audio/ogg",
		"flac": "audio/x-flac",
		"wav":  "audio/x-wav",
		"amr":  "audio/amr",
		"aac":  "audio/aac",
		// Archive
		"epub":   "application/epub+zip",
		"zip":    "application/zip",
		"tar":    "application/x-tar",
		"rar":    "application/vnd.rar",
		"gz":     "application/gzip",
		"bz2":    "application/x-bzip2",
		"7z":     "application/x-7z-compressed",
		"xz":     "application/x-xz",
		"pdf":    "application/pdf",
		"exe":    "application/vnd.microsoft.portable-executable",
		"swf":    "application/x-shockwave-flash",
		"rtf":    "application/rtf",
		"iso":    "application/x-iso9660-image",
		"eot":    "application/octet-stream",
		"ps":     "application/postscript",
		"sqlite": "application/vnd.sqlite3",
		"nes":    "application/x-nintendo-nes-rom",
		"crx":    "application/x-google-chrome-extension",
		"cab":    "application/vnd.ms-cab-compressed",
		"deb":    "application/vnd.debian.binary-package",
		"ar":     "application/x-unix-archive",
		"Z":      "application/x-compress",
		"lz":     "application/x-lzip",
		"rpm":    "application/x-rpm",
		"elf":    "application/x-executable",
		"dcm":    "application/dicom",
		// Documents
		"doc":  "application/msword",
		"docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"xls":  "application/vnd.ms-excel",
		"xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"ppt":  "application/vnd.ms-powerpoint",
		"pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		// Font
		"woff":  "application/font-woff",
		"woff2": "application/font-woff",
		"ttf":   "application/font-sfnt",
		"otf":   "application/font-sfnt",
		// Application
		"wasm": "application/wasm",
	}

	if mimeTypes[ext] != "" {
		return mimeTypes[ext], nil
	}
	return "", fmt.Errorf("No mimetype found for extension %s", ext)
}
