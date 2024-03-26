package main

import (
  "fmt"
  "flag"
  "os"
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
      line_elements := strings.Split(line, ";")
      status, _ := strconv.ParseBool(line_elements[1]) 
      *task_list = append(*task_list, Task{line_elements[0], status})
    }
  } else {
    log.Fatal("Error while reading a file")
  }
}


func file_write (file_path string, task_list *[]Task) {
  if os.Truncate(file_path, 0) != nil {
    log.Fatal("Error while truncating a file")
  }
  file, err := os.OpenFile(file_path, os.O_WRONLY, 0644)
  if err != nil {
    log.Fatal("Error while opening a file")
  }
  for _, line := range *task_list {
    file.WriteString(line.name + ";" + strconv.FormatBool(line.status) + "\n")
  }
  defer file.Close()
}


func main () {
  var file_path string = "./memory"  
  var taskList []Task;

  file_check(file_path)
  file_read(file_path, &taskList)

  var name = flag.String("n", "" ,"New task" )
  var remove = flag.Int("r", -1, "Remove task with given index")
  var status = flag.Int("s", -1, "Change the status of a task")
  var out = flag.Bool("p", false, "Print tasks")

  flag.Parse()

  if *name != "" {
      var task = Task{name: *name, status: false}
      taskList = append(taskList, task)
  }
  if *out {
      fmt.Println("Tasks:")
      print_elements(taskList)
  } 
  if *remove != -1 {
      taskList = append(taskList[:*remove], taskList[*remove+1:]...)
  }
  if *status != -1 {
      taskList[*status].status = !(taskList[*status].status) 
  }
  
  file_write(file_path, &taskList)
}

