package procfsutils

import "testing"

func TestGetrmem(t *testing.T) {
    rmem, err := Getrmem(0)
    if err != nil {
        t.Fatal("unexpected error in Getrem: ", err)
    }
    if rmem.Size == 0 || rmem.Resident == 0 {
        t.Error("unexpected zero size for Size or Resident")
    }
}
