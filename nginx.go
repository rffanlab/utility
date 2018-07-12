package utility

import (

)

//type Nginx struct {
//	Upstream struct{
//		Name string
//		Values struct{
//			Ip_hash bool;
//			Server []struct{
//				Server string
//			};
//		}
//	}
//	Server struct{
//		Listen int
//		ServerName []string
//		Allow string
//		Index []string
//		Root string
//		Location []struct{
//			Root string
//			Proxy_next_upstream []string
//			Proxy_pass string
//			Proxy_set_header []struct{
//				Host string
//				XRealIP string  //X-Real-IP
//				XForwardedFor string // X-Forwarded-For
//			}
//			If []struct{
//				Condition struct{
//
//				}
//				Action []struct{
//
//				}
//			}
//		}
//	}
//}


//
//func ReadLines(path string) (lines []string,err error) {
//	f,err := os.Open(path)
//	if err != nil{
//		fmt.Println(err)
//	}
//	defer f.Close()
//
//	scanner := bufio.NewScanner(f)
//	for  scanner.Scan()  {
//		lines = append(lines,scanner.Text())
//	}
//	if err := scanner.Err();err != nil{
//		fmt.Fprintln(os.Stderr,err)
//	}
//
//	return
//}
//



