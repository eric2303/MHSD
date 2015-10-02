package MHSD
//init shift during start up
var shift []uint;
func init(){
	shift =make([]uint,8)
	for i:=range shift{
		shift[i] = uint(i)<<3;
	}
}
/*
General idea：get a 64 bit int, split it into 8 bytes , the first four character
hold the length of the minhash array
the max size if 2^32-1
 */
func SerializeSingle(minimums []uint64) []byte{
	//8 bytes in 64 bit unsign int
	length := len(minimums);
	output := make([]byte,(length<<3)+4);
	//the first four bytes of the output store the length of the signature can represent up to 2^32-1
	output[0] = byte((length & 0xff000000)>>24);
	output[1] = byte((length & 0xff0000)>>16);
	output[2] = byte((length & 0xff00)>>8);
	output[3] = byte((length & 0xff));
	for i := range minimums{
		inputInt := minimums[i];
		index := (i<<3)+4;
		for j:=0; j< 8;j++ {
			output[index+j] = byte(inputInt & 0xff);
			inputInt>>=8;
		}
	}
	return output;
}

func DeserializeSingle(input []byte) []uint64{
	//the first 4 members hold the total length
	length := (int(input[0])<<24)+(int(input[1])<<16)+(int(input[2])<<8)+(int(input[3]));
	output := make([]uint64,length);
	for i :=range output{
		index := (i<<3)+4;
		sum := uint64(0);
		for j:=0;j<8;j++ {
			sum += uint64(input[index+j])<<shift[j];
		}
		output[i] = sum;
	}
	return output;
}

