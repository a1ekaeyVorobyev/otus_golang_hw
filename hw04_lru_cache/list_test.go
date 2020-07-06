package hw04_lru_cache //nolint:golint,stylecheck
import (
	"fmt"
	"testing"
)

func Test_CheckCreate(t *testing.T) {
	fmt.Println("aa")
	//checkValue := []string{"dd", "4", "5", "true"}
	fmt.Println("aaa")
	s := ListNew()
	fmt.Println("aaa")
	s.PushFront("dd")
	s.PushFront("4")
	s.PushFront(5)
	s.PushFront(true)
	fmt.Println("Len = ", s.Len())
	if s.Len() != 4 {
		t.Error("Не верный размер ", s.Len())
	}
}
