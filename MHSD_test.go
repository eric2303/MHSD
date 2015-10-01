package MHSD
import (
	"hash/fnv"
	"testing"
	"fmt"
	"github.com/dgryski/go-minhash"
	"time"
)
var h1, h2 minhash.Hash64
func init() {
	fnv1 := fnv.New64a()
	fnv2 := fnv.New64a()
	h1 = func(b []byte) uint64 {
		fnv1.Reset()
		fnv1.Write([]byte{0xd5, 0x9d, 0x34, 0x24, 0xf4, 0x15, 0xa4, 0xfe})
		fnv1.Write(b)
		return fnv1.Sum64()
	}
	h2 = func(b []byte) uint64 {
		fnv2.Reset()
		fnv2.Write([]byte{0x06, 0x43, 0x0c, 0xd9, 0xde, 0x74, 0xbf, 0xf1})
		fnv2.Write(b)
		return fnv2.Sum64()
	}
}
func data(size int) [][]byte {
	d := make([][]byte, size)
	for i := range d {
		d[i] = []byte(fmt.Sprintf("salt%d %d", i, size))
	}
	return d
}
func hashing(mh *minhash.MinWise, size int, data [][]byte) {
	for i := 0; i < size; i++ {
		mh.Push(data[i])
	}
}
func implementMinHash(dSize int,minSize int)*minhash.MinWise{
	d:=data(dSize);
	m1 := minhash.NewMinWise(h1,h2,minSize);
	hashing(m1, dSize,d)
	return m1;
}
func runTest(dSize int, minSize int,iteration int){
	m1:=implementMinHash(dSize,minSize).Signature();
	t0 :=time.Now();
	var ms []byte;
	var m2 []uint64;
	for i:=0;i<iteration;i++{
		ms = SerializeSingle(m1);
	}
	t1 :=time.Now();
	for i:=0;i<iteration;i++{
		m2 =DeserializeSingle(ms);
	}
	t2 :=time.Now()
	for i:=0;i<minSize;i++{
		if(m1[i]!=m2[i]){
			fmt.Printf("index %d not equal soruce:%d deserialize:%d\n",i,m1[i],m2[i]);
		}
	}
	fmt.Printf("time to take to run Serializing of minhash size %d %d times:%v\n",minSize,iteration,t1.Sub(t0));
	fmt.Printf("time to take to run Deserializing of minhash size %d %d times:%v\n",minSize,iteration,t2.Sub(t1));
}
func Test(t *testing.T){
	runTest(50000,128,1000);
	runTest(50000,256,1000);
	runTest(50000,512,1000);
	runTest(50000,1024,1000);
	runTest(50000,2048,1000);
	runTest(50000,4096,1000);
	runTest(50000,8192,1000);
	runTest(50000,16384,1000);
	runTest(50000,32768,1000);
	runTest(50000,65535,1000);

}

