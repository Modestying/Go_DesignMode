package prototype_test

import "testing"
import "prototype/prototype"
func TestPrototype(t *testing.T) {
	file1 := &prototype.file{
		name: "mp4",
	}
	file2 := &prototype.file{
		name: "music",
	}
	file3 := &prototype.file{
		name: "picture",
	}

	folder1 := &prototype.folder{
		name:  "Download",
		child: []inode{file1, file2, file3},
	}

	file1.print()

	folder2 := folder1.clone()
	folder2.print()

}
