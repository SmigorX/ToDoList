package main

import (
  "fmt"
  "flag"
  "os"
  //"io/ioutil"
  "log"
  "strings"
  "strconv"
)


type Task struct {
  name string
  status bool
}


func print_elements (list []Task) {
  for i:=0; i<len(list); i++ {
    if list[i].status {
      fmt.Printf("%d.[x] %s\n", i, list[i].name)
    } else {
      fmt.Printf("%d.[ ] %s\n", i, list[i].name)
    }
  }
}


func file_check (file_path string) {
  file, err := os.OpenFile(file_path, os.O_CREATE, 0644)
 
  if err != nil {
    fmt.Printf("Error while creating a file %e", err)
  }
  defer file.Close() 
}


func file_read (file_path string, task_list *[]Task) {
  data, err := os.ReadFile(file_path)
  
  if err == nil {
    document_lines := strings.Split(string(data), "\n")
    document_lines = document_lines[:len(document_lines)-1]
    for _, line := range document_lines {
      s := strings.Split(line, ";")
      var b, _ = strconv.ParseBool(s[1])
      fmt.Println("Tiru riru")
      fmt.Println(s[1], b)
      var new_task = Task{s[0], b}
      *task_list = append(*task_list, new_task)
    }
  } else {
    log.Fatal("Error while reading a file")
  }
}


func file_write (file_path string) {

}


func main () {
  var file_path string = "./memory"  
  var taskList []Task;

  file_check(file_path)

  file_read(file_path, &taskList)

  print_elements(taskList)
//  var name = flag.String("n", "" ,"Name of the task" )
//  var remove = flag.Int("r", 0, "Remove task with given index")
//  var status = flag.Int("s", 0, "Change the status of a task")
//  var out = flag.Bool("p", false, "Print tasks")

  flag.Parse()

//  if *name != "" {
    //var task = Task{id:len(taskList), name: *name,}
    //taskList = append(taskList, task)
//  } else if *out {
//    print_elements(taskList)
//  } else if *remove != 0 {
//    taskList = append(taskList[:*remove], taskList[*remove+1:]...)
//  } else if *status != 0 {
//    taskList[*status].status = !(taskList[*status].status)  
//  }
}
