package MHSD
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
	output[0] = intShiftByte(length,24);
	output[1] = intShiftByte(length,16);
	output[2] = intShiftByte(length,8);
	output[3] = intShiftByte(length,0);
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

//return the last 8 bit of a int after right shift
func intShiftByte(input int,shiftNum uint) byte{
	return byte((input>>shiftNum)&0xff);
}
//return the int value after left shift
func byteShiftInt(input byte,shiftNum uint) int{
	return int(input)<<shiftNum;
}
func DeserializeSingle(input []byte) []uint64{
	//the first 4 members hold the total length
	length := byteShiftInt(input[0],24)+byteShiftInt(input[1],16)+byteShiftInt(input[2],8)+byteShiftInt(input[3],0);
	output := make([]uint64,length);
	for i :=range output{
		index := (i<<3)+4;
		sum := uint64(0);
		for j:=7;j>=0;j-- {
			sum<<=8
			sum += uint64(input[index+j]);
		}
		output[i] = sum;
	}
	return output;
}

