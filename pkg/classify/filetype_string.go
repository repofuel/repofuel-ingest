// Code generated by "stringer -type=FileType -linecomment -trimprefix=File"; DO NOT EDIT.

package classify

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FileBinary-1]
	_ = x[FileConfiguration-2]
	_ = x[FileDocumentation-3]
	_ = x[FileCode-4]
	_ = x[FileGenerated-5]
	_ = x[FileDependency-6]
	_ = x[FileTests-7]
	_ = x[FileSymlink-8]
}

const _FileType_name = "BinaryConfigurationDocumentationCodeGeneratedDependencyTestsSymlink"

var _FileType_index = [...]uint8{0, 6, 19, 32, 36, 45, 55, 60, 67}

func (i FileType) String() string {
	i -= 1
	if i >= FileType(len(_FileType_index)-1) {
		return "FileType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _FileType_name[_FileType_index[i]:_FileType_index[i+1]]
}