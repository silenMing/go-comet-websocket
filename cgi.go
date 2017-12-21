package main

import(
	"net/http/cgi"
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		handler := new(cgi.Handler);
		handler.Path = "/usr/local/go/bin/go";
		script := "$HOME/cgi-script" + r.URL.Path;
		fmt.Println(handler.Path);
		args := []string{"run",script}
		handler.Args = append(handler.Args, args...);

		handler.ServeHTTP(w, r);
	});
	http.ListenAndServe(":8899",nil)
	select {}//阻塞进程
	
}

