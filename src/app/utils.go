package main

func checkPath(url, reqUrl string) (bool){
	if reqUrl != url{		
		return false
	}
	return true
}