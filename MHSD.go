package MHSD
/*
General idea：get a 64 bit int, split it into 8 bytes , the first two character
hold the length of the minhash array
the max size if 2^16-1 = 65535
 */
func SerializeSingle(minimums []uint64) []byte{
	//8 bytes in 64 bit unsign int
	length := len(minimums);
	output := make([]byte,(length<<3)+2);
	//the first two bytes of the output store the length of the signature can represent up to 2^16
	output[0] = byte((length & 0xff00)>>8);
	output[1] = byte(length & 0xff);
	for i := range minimums{
		inputInt := minimums[i];
		index := (i<<3)+2;
		for j:=0; j< 8;j++ {
			output[index+j] = byte(inputInt & 0xff);
			inputInt>>=8;
		}
	}
	return output;
}

func DeserializeSingle(input []byte) []uint64{
	//the first two member hold the total length
	length := (int(input[0])<<8)+int(input[1]);
	output := make([]uint64,length);
	shift :=make([]uint,8)
	for i:=range shift{
		shift[i] = uint(i)<<3;
	}
	for i :=range output{
		index := (i<<3)+2;
		sum := uint64(0);
		for j:=0;j<8;j++ {
			sum += uint64(input[index+j])<<shift[j];
		}
		output[i] = sum;
	}
	return output;
}

