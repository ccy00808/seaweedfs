package mount

import (
	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/seaweedfs/seaweedfs/weed/glog"
	"github.com/seaweedfs/seaweedfs/weed/pb/filer_pb"
	"github.com/seaweedfs/seaweedfs/weed/util"
	"strconv"
)

func (wfs *WFS) AcquireHandle(inode uint64, flags, uid, gid uint32) (fileHandle *FileHandle, status fuse.Status) {
	var entry *filer_pb.Entry
	var path util.FullPath
	path, _, entry, status = wfs.maybeReadEntry(inode)
	if status == fuse.OK {
		if wfs.wormEnabledForEntry(path, entry) && flags&fuse.O_ANYWRITE != 0 {
			return nil, fuse.EPERM
		}
		// need to AcquireFileHandle again to ensure correct handle counter
		fileHandle = wfs.fhMap.AcquireFileHandle(wfs, inode, entry)
	}
	glog.V(4).Infof("******** AcquireHandle:" + entry.GetName() + ":" + strconv.FormatUint(uint64(fileHandle.fh), 10))
	return
}

func (wfs *WFS) ReleaseHandle(handleId FileHandleId) {
	wfs.fhMap.ReleaseByHandle(handleId)
}

func (wfs *WFS) GetHandle(handleId FileHandleId) *FileHandle {
	glog.V(4).Infof("******** ReleaseHandle:" + strconv.FormatUint(uint64(handleId), 10))
	return wfs.fhMap.GetFileHandle(handleId)
}
