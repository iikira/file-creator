package trunctest

import (
	"fmt"
	"github.com/iikira/BaiduPCS-Go/pcsutil/converter"
	"golang.org/x/sys/windows"
	"os"
	"syscall"
	"time"
)

var (
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	procSetFileValidData = kernel32.NewProc("SetFileValidData")
)

func TestTrunc2() {
	filename := "C:/Users/syy/file.txt"
	func() {
		current, err := windows.GetCurrentProcess()
		if err != nil {
			fmt.Printf("GetCurrentProcess error: %s\n", err)
			return
		}

		var hToken windows.Token
		err = windows.OpenProcessToken(current, windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &hToken)
		if err != nil {
			fmt.Printf("OpenProcessToken error: %s\n", err)
			return
		}

		var (
			SE_MANAGE_VOLUME_NAME, _ = windows.UTF16PtrFromString("SeManageVolumePrivilege")
			tp                       = windows.Tokenprivileges{
				PrivilegeCount: 1,
				Privileges: [1]windows.LUIDAndAttributes{
					windows.LUIDAndAttributes{
						Luid:       windows.LUID{},
						Attributes: windows.SE_PRIVILEGE_ENABLED,
					},
				},
			}
		)
		err = windows.LookupPrivilegeValue(nil, SE_MANAGE_VOLUME_NAME, &tp.Privileges[0].Luid)
		if err != nil {
			fmt.Printf("LookupPrivilegeValue error: %s\n", err)
			return
		}

		err = windows.AdjustTokenPrivileges(hToken, false, &tp, 0, nil, nil)
		if err != nil {
			fmt.Printf("AdjustTokenPrivileges error: %s\n", err)
			return
		}

		fmt.Printf("%#v\n", &tp)

		f, err := os.Create(filename)
		if err != nil {
			fmt.Printf("create error: %s\n", err)
			return
		}
		defer f.Close()

		err = f.Truncate(128 * converter.MB)
		if err != nil {
			fmt.Printf("trunc error: %s\n", err)
			return
		}

		r1, _, err := procSetFileValidData.Call(f.Fd(), uintptr(128*converter.MB))
		if r1 == 0 {
			fmt.Printf("SetFileValidData error: %s\n", err)
			return
		}

		start := time.Now()
		_, err = f.WriteAt([]byte("111"), 27*converter.MB)
		if err != nil {
			fmt.Printf("write error: %s\n", err)
			return
		}

		now := time.Now()
		fmt.Printf("time elapse: %s\n", now.Sub(start))
	}()

	err := os.Remove(filename)
	if err != nil {
		fmt.Printf("remove error: %s\n", err)
	}
}
