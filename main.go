
package main

import (
    "log"
    "os"

    "github.com/golang/protobuf/proto"
     // Replace with the actual generated package path
)

func main() {
    // Create a new Person message
    john := &pb.Person{
        Id:    1234,
        Name:  "John Doe",
        Email: "jdoe@example.com",
    }

    // Serialize to bytes
    data, err := proto.Marshal(john)
    if err != nil {
        log.Fatalf("Failed to marshal person: %v", err)
    }

    // Write to file
    filename := "output.bin" // Replace with your desired filename
    if err := os.WriteFile(filename, data, 0644); err != nil {
        log.Fatalf("Failed to write to file: %v", err)
    }

    log.Printf("Person data written to %s", filename)
}
