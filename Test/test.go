package Test

//func GetConfig() {
//
//	data, err := ioutil.ReadFile("/Applications/coding/go/go_algorithms/Configs/db.yml")
//	if err != nil {
//		log.Printf("yamlFile.Get err   #%v ", err)
//	}
//	var v interface{}
//	err = yaml.Unmarshal(data, &v)
//	if err != nil {
//		log.Printf("unmarshal err   #%v ", err)
//	}
//
//	var DbInfo = make(map[string]string)
//	for key, value := range v.(map[interface{}]interface{}) {
//
//		if key == TEXT_DEFAULT_CONNECTION {
//			for k, v := range value.(map[interface{}]interface{}) {
//
//			}
//		}
//	}
//}