package main

func matrix() (string,error){
	f, err := os.Open("ids.txt")
	if err != nil {
		fmt.Println(" downloaded.txt not exist ")
		return "",errors.New("open error")
	}
	defer f.Close()

	var last string
	reader := bufio.NewReader(f)
	
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 1 {
			last = line
		}
		if err != nil || io.EOF == err {
			break
		}
		// fmt.Println(line)
	}
	fmt.Println("last downloaded key: ",last)
	return last,nil
}


func main(){

}