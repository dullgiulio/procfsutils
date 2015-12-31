package procfsutils

import (
    "fmt"
    "os"
    "io/ioutil"
)

type RMem struct {
    Size int64
    Resident int64
    Share int64
    Text int64
    Lib int64
    Data int64
    Dt int64
}

func (rmem *RMem) String() string {
    return fmt.Sprintf("%d %d %d %d %d %d %d", rmem.Size,
            rmem.Resident, rmem.Share, rmem.Text, rmem.Lib, rmem.Data, rmem.Dt)
}

func Getrmem(pid int) (*RMem, error) {
    if pid == 0 {
        pid = os.Getpid()
    }
    data, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/statm", pid))
    if err != nil {
        return nil, err
    }
    rmem := new(RMem)
    _, err = fmt.Sscanf(string(data), "%d %d %d %d %d %d %d", &rmem.Size,
        &rmem.Resident, &rmem.Share, &rmem.Text, &rmem.Lib, &rmem.Data, &rmem.Dt)
    if err != nil {
        return nil, err
    }
    return rmem, nil
}
