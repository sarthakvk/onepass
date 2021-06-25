package backend

import (
	// "encoding/json"
	// "os"
)

func AddPassword(name, url, pass string){
	entry := Password{
		name,
		url,
		pass,
	}
	loadData()
	_data[name] = entry
	saveData()
}

