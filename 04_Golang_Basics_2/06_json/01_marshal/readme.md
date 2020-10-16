**What does 'Marshaling' mean?**

* Marshaling is the process of transforming the memory representation of an object to a data format suitable for storage or transmission, and it is typically used when data must be
  moved between different parts of a computer program or from one program to another.
* This Marshaling and UnMarshaling is Golang trying to convert struct into JSON objects and JSON objects into Golang structs 

**func Marshal**
```
func Marshal(v interface{}) ([]byte, error)
```
* Marshal returns the JSON encoding of v.
  
