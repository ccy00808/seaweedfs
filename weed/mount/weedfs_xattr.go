package mount

import (
	"github.com/seaweedfs/seaweedfs/weed/stats"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/hanwen/go-fuse/v2/fuse"
	sys "golang.org/x/sys/unix"
)

const (
	// https://man7.org/linux/man-pages/man7/xattr.7.html#:~:text=The%20VFS%20imposes%20limitations%20that,in%20listxattr(2)).
	MAX_XATTR_NAME_SIZE  = 255
	MAX_XATTR_VALUE_SIZE = 65536
	XATTR_PREFIX         = "xattr-" // same as filer
)

// GetXAttr reads an extended attribute, and should return the
// number of bytes. If the buffer is too small, return ERANGE,
// with the required buffer size.
func (wfs *WFS) GetXAttr(cancel <-chan struct{}, header *fuse.InHeader, attr string, dest []byte) (size uint32, code fuse.Status) {
	stats.FuseRequestCounter.WithLabelValues("getxattr").Inc()
	start := time.Now()
	if wfs.option.DisableXAttr {
		return 0, fuse.Status(syscall.ENOTSUP)
	}

	//validate attr name
	if len(attr) > MAX_XATTR_NAME_SIZE {
		if runtime.GOOS == "darwin" {
			return 0, fuse.EPERM
		} else {
			return 0, fuse.ERANGE
		}
	}
	if len(attr) == 0 {
		return 0, fuse.EINVAL
	}

	_, _, entry, status := wfs.maybeReadEntry(header.NodeId)
	if status != fuse.OK {
		return 0, status
	}
	if entry == nil {
		return 0, fuse.ENOENT
	}
	if entry.Extended == nil {
		return 0, fuse.ENOATTR
	}
	data, found := entry.Extended[XATTR_PREFIX+attr]
	if !found {
		return 0, fuse.ENOATTR
	}
	if len(dest) < len(data) {
		return uint32(len(data)), fuse.ERANGE
	}
	copy(dest, data)

	stats.FuseRequestCost.WithLabelValues("getxattr").Observe(float64(time.Since(start).Microseconds()))
	return uint32(len(data)), fuse.OK
}

// SetXAttr writes an extended attribute.
// https://man7.org/linux/man-pages/man2/setxattr.2.html
//
//	 By default (i.e., flags is zero), the extended attribute will be
//	created if it does not exist, or the value will be replaced if
//	the attribute already exists.  To modify these semantics, one of
//	the following values can be specified in flags:
//
//	XATTR_CREATE
//	       Perform a pure create, which fails if the named attribute
//	       exists already.
//
//	XATTR_REPLACE
//	       Perform a pure replace operation, which fails if the named
//	       attribute does not already exist.
func (wfs *WFS) SetXAttr(cancel <-chan struct{}, input *fuse.SetXAttrIn, attr string, data []byte) fuse.Status {
	stats.FuseRequestCounter.WithLabelValues("setxattr").Inc()
	start := time.Now()
	if wfs.option.DisableXAttr {
		return fuse.Status(syscall.ENOTSUP)
	}

	if wfs.IsOverQuota {
		return fuse.Status(syscall.ENOSPC)
	}

	//validate attr name
	if len(attr) > MAX_XATTR_NAME_SIZE {
		if runtime.GOOS == "darwin" {
			return fuse.EPERM
		} else {
			return fuse.ERANGE
		}
	}
	if len(attr) == 0 {
		return fuse.EINVAL
	}
	//validate attr value
	if len(data) > MAX_XATTR_VALUE_SIZE {
		if runtime.GOOS == "darwin" {
			return fuse.Status(syscall.E2BIG)
		} else {
			return fuse.ERANGE
		}
	}

	path, fh, entry, status := wfs.maybeReadEntry(input.NodeId)
	if status != fuse.OK {
		return status
	}
	if entry == nil {
		return fuse.ENOENT
	}
	if fh != nil {
		fh.entryLock.Lock()
		defer fh.entryLock.Unlock()
	}

	if entry.Extended == nil {
		entry.Extended = make(map[string][]byte)
	}
	oldData, _ := entry.Extended[XATTR_PREFIX+attr]
	switch input.Flags {
	case sys.XATTR_CREATE:
		if len(oldData) > 0 {
			break
		}
		fallthrough
	case sys.XATTR_REPLACE:
		fallthrough
	default:
		entry.Extended[XATTR_PREFIX+attr] = data
	}

	if fh != nil {
		fh.dirtyMetadata = true
		return fuse.OK
	}

	result := wfs.saveEntry(path, entry)
	stats.FuseRequestCost.WithLabelValues("getxattr").Observe(float64(time.Since(start).Microseconds()))
	return result

}

// ListXAttr lists extended attributes as '\0' delimited byte
// slice, and return the number of bytes. If the buffer is too
// small, return ERANGE, with the required buffer size.
func (wfs *WFS) ListXAttr(cancel <-chan struct{}, header *fuse.InHeader, dest []byte) (n uint32, code fuse.Status) {
	stats.FuseRequestCounter.WithLabelValues("listxattr").Inc()
	start := time.Now()
	if wfs.option.DisableXAttr {
		return 0, fuse.Status(syscall.ENOTSUP)
	}

	_, _, entry, status := wfs.maybeReadEntry(header.NodeId)
	if status != fuse.OK {
		return 0, status
	}
	if entry == nil {
		return 0, fuse.ENOENT
	}
	if entry.Extended == nil {
		return 0, fuse.OK
	}

	var data []byte
	for k := range entry.Extended {
		if strings.HasPrefix(k, XATTR_PREFIX) {
			data = append(data, k[len(XATTR_PREFIX):]...)
			data = append(data, 0)
		}
	}
	if len(dest) < len(data) {
		return uint32(len(data)), fuse.ERANGE
	}

	copy(dest, data)

	stats.FuseRequestCost.WithLabelValues("listxattr").Observe(float64(time.Since(start).Microseconds()))
	return uint32(len(data)), fuse.OK
}

// RemoveXAttr removes an extended attribute.
func (wfs *WFS) RemoveXAttr(cancel <-chan struct{}, header *fuse.InHeader, attr string) fuse.Status {
	stats.FuseRequestCounter.WithLabelValues("removexattr").Inc()
	start := time.Now()
	if wfs.option.DisableXAttr {
		return fuse.Status(syscall.ENOTSUP)
	}

	if len(attr) == 0 {
		return fuse.EINVAL
	}
	path, fh, entry, status := wfs.maybeReadEntry(header.NodeId)
	if status != fuse.OK {
		return status
	}
	if entry == nil {
		return fuse.OK
	}
	if fh != nil {
		fh.entryLock.Lock()
		defer fh.entryLock.Unlock()
	}

	if entry.Extended == nil {
		return fuse.ENOATTR
	}
	_, found := entry.Extended[XATTR_PREFIX+attr]

	if !found {
		return fuse.ENOATTR
	}

	delete(entry.Extended, XATTR_PREFIX+attr)

	result := wfs.saveEntry(path, entry)
	stats.FuseRequestCost.WithLabelValues("removexattr").Observe(float64(time.Since(start).Microseconds()))
	return result
}
