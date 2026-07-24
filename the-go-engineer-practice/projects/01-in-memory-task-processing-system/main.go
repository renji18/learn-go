/*
Build a system that:
	•	accepts tasks
	•	processes them
	•	tracks state transitions
	•	supports multiple task types
*/

/*
My understanding:

First we create some helpers

We create a const of type int representation of all the status types
- Pending
- Processing
- Completed
- Failed

We also added a fmt.Stringer method on the Status

Task is an interface which defines 4 methods on it:
- Add Task
- Process Task
- List Tasks
- fmt.Stringer

We then have a TaskManager which holds a map of all tasks where the key is the task name and the value is the Task

We also create a map to hold supported task types, for validation

BaseTask is a common struct for all Tasks with 2 fields
- name
- status

We embed this BaseTask in our specific task types
- EmailTask
- FileTask

Once all necessary setup is done, we then build the 4 methods for both the EmailTask and FileTask
- Name() -> returns the .name of the task struct
- Process() -> checks if .status of the task is completed/failed then return error.
	- For EmailTask: If .Receiver is empty or does not contain an '@' character, then status is failed
	- For FileTask: If .Path is empty, then status is failed
	- The status is updated to Processing
	- The status is updated to Completed
- Status() -> returns the .status.String()
- String() -> implements the fmt.Stringer interface by providing context for a task

We then work on the methods for TaskManager
- AddTask(taskString string) -> this method first splits the empty string ito 4 parts, first being the type email/file, second is the name of task, third is the main content and forth is the sub content
	-> We check if the taskType is supported or not using the supportedTaskTypes
	-> We check if a task already exists with the same name using tm.tasks
	-> We then build the BaseTask
	-> We then build the EmailTask or FileTask based on the taskType provided
- ListTasks() -> We make a new slice of Tasks, map over all the tasks from the task manager and append each task to our new slice, and then return it

Now in the main function, when we get tasks from ListTasks fn, it returns us all the Tasks in a slice.
When we run the .Process() method it attaches the address of the slice to the Process method and then the operation is performed on the actual Task instead of on a copy.
*/

package main

import (
	"fmt"
	"strings"
)

type Status int

const (
	PENDING Status = iota
	PROCESSING
	COMPLETED
	FAILED
)

func (s Status) String() string {
	switch s {
	case PENDING:
		return "Pending"
	case PROCESSING:
		return "Processing"
	case COMPLETED:
		return "Completed"
	default:
		return "Failed"
	}
}

type Task interface {
	Name() string
	Process() error
	Status() string
	fmt.Stringer
}

type TaskManager struct {
	tasks map[string]Task
}

type TaskFactory func(baseTask BaseTask, mainContent string, subContent string) Task

var taskFactories map[string]TaskFactory = map[string]TaskFactory{
	"email": func(baseTask BaseTask, mainContent, subContent string) Task {
		return &EmailTask{BaseTask: baseTask, Receiver: mainContent, Message: subContent}
	},
	"file": func(baseTask BaseTask, mainContent, subContent string) Task {
		return &FileTask{BaseTask: baseTask, Path: mainContent, Content: subContent}
	},
	"whatsapp": func(baseTask BaseTask, mainContent, subContent string) Task {
		return &WhatsappTask{BaseTask: baseTask, Number: mainContent, Chat: subContent}
	},
}

type BaseTask struct {
	name   string
	status Status
}

type EmailTask struct {
	BaseTask
	Receiver string
	Message  string
}

type FileTask struct {
	BaseTask
	Path    string
	Content string
}

type WhatsappTask struct {
	BaseTask
	Number string
	Chat   string
}

// Email Methods
func (e EmailTask) Name() string {
	return e.name
}

func (e *EmailTask) Process() error {
	if e.status == COMPLETED {
		return fmt.Errorf("Cannot process task %s, as it is already completed", e.name)
	}

	if e.status == FAILED {
		return fmt.Errorf("Cannot process task %s, as it has failed", e.name)
	}

	if e.Receiver == "" || !strings.Contains(e.Receiver, "@") {
		e.status = FAILED
		return fmt.Errorf("Invalid email")
	}

	fmt.Println("========\nProcessing task...")
	e.status = PROCESSING
	e.status = COMPLETED
	fmt.Println("Processing complete!!!\n========")

	return nil
}

func (e EmailTask) Status() string {
	return e.status.String()
}

func (e EmailTask) String() string {
	return fmt.Sprintf("Task Type: (email)\nTask Name: (%s)\nTask Status: (%s)\nEmail Receiver: (%s)\nEmail Message: (%s)", e.name, e.Status(), e.Receiver, e.Message)
}

// File Methods
func (f FileTask) Name() string {
	return f.name
}

func (f *FileTask) Process() error {
	if f.status == COMPLETED {
		return fmt.Errorf("Cannot process task %s, as it is already completed", f.name)
	}

	if f.status == FAILED {
		return fmt.Errorf("Cannot process task %s, as it has failed", f.name)
	}

	if f.Path == "" {
		f.status = FAILED
		return fmt.Errorf("Missing file name")
	}

	fmt.Println("========\nProcessing task...")
	f.status = PROCESSING
	f.status = COMPLETED
	fmt.Println("Processing complete!!!\n========")

	return nil
}

func (f FileTask) Status() string {
	return f.status.String()
}

func (f FileTask) String() string {
	return fmt.Sprintf("Task Type: (file)\nTask Name: (%s)\nTask Status: (%s)\nFile Path: (%s)\nFile Content: (%s)", f.name, f.Status(), f.Path, f.Content)
}

// Whatsapp Methods
func (w WhatsappTask) Name() string {
	return w.name
}

func (w *WhatsappTask) Process() error {
	if w.status == COMPLETED {
		return fmt.Errorf("Cannot process task %s, as it is already completed", w.name)
	}

	if w.status == FAILED {
		return fmt.Errorf("Cannot process task %s, as it has failed", w.name)
	}

	if w.Number == "" {
		w.status = FAILED
		return fmt.Errorf("Phone number missing")
	}

	fmt.Println("========\nProcessing task...")
	w.status = PROCESSING
	w.status = COMPLETED
	fmt.Println("Processing complete!!!\n========")

	return nil
}

func (w WhatsappTask) Status() string {
	return w.status.String()
}

func (w WhatsappTask) String() string {
	return fmt.Sprintf("Task Type: (file)\nTask Name: (%s)\nTask Status: (%s)\nNumber: (%s)\nChat Message: (%s)", w.name, w.Status(), w.Number, w.Chat)
}

// Task Manager Methods
func (tm *TaskManager) AddTask(taskString string) error {
	// task string pattern task-type|task-name|c1|c2
	taskParts := strings.SplitN(taskString, "|", 4)

	if len(taskParts) != 4 {
		return fmt.Errorf("Invalid input pattern")
	}

	taskType := taskParts[0]
	taskName := taskParts[1]
	mainContent := taskParts[2]
	subContent := taskParts[3]

	factory, supported := taskFactories[taskType]

	if !supported {
		return fmt.Errorf("Unsupported task type: %s", taskType)
	}

	if _, exists := tm.tasks[taskName]; exists {
		return fmt.Errorf("Task with name %s already exists", taskName)
	}

	baseTask := BaseTask{name: taskName, status: PENDING}

	tm.tasks[taskName] = factory(baseTask, mainContent, subContent)

	return nil
}

func (tm TaskManager) ListTasks() []Task {
	tasks := make([]Task, 0)

	for _, b := range tm.tasks {
		tasks = append(tasks, b)
	}

	return tasks
}

func main() {
	tm := TaskManager{
		tasks: make(map[string]Task, 0),
	}

	err := tm.AddTask("file|task1|/user/dustbin|Or Operator is used like this: true | false, yes | no, and in some programming languages: a || b")

	if err != nil {
		fmt.Println(err)
	}

	err = tm.AddTask("file|task2||Hey!~")

	if err != nil {
		fmt.Println(err)
	}

	err = tm.AddTask("email|task3|random@email.com|Regards, Thanks")

	if err != nil {
		fmt.Println(err)
	}

	err = tm.AddTask("email|task3|random|Hey!~")

	if err != nil {
		fmt.Println(err)
	}

	err = tm.AddTask("email|task4|random|Hey!~ How are you?|?")

	if err != nil {
		fmt.Println(err)
	}

	err = tm.AddTask("whatsapp|task5|+917600081901|Hi, Aadarsh this side")

	if err != nil {
		fmt.Println(err)
	}

	err = tm.AddTask("whatsapp|task6||Hi, Aadarsh this side")

	if err != nil {
		fmt.Println(err)
	}

	err = tm.AddTask("music|task7||Hi, Aadarsh this side")

	if err != nil {
		fmt.Println(err)
	}

	tasks := tm.ListTasks()

	for _, b := range tasks {
		fmt.Println("=================================================================================================================")
		fmt.Println("Task Name: ", b.Name())
		fmt.Println("Task Status before Processing ", b.Status())
		err := b.Process()
		if err != nil {
			fmt.Printf("Error processing task: %v\n", err)
		} else {
			fmt.Println("Task Status after Processing ", b.Status())
			fmt.Println(b.String())
		}
		fmt.Println("=================================================================================================================")
	}

	for _, b := range tasks {
		fmt.Println("=================================================================================================================")
		fmt.Println("Task Name: ", b.Name())
		fmt.Println("Task Status before Processing ", b.Status())
		err := b.Process()
		if err != nil {
			fmt.Printf("Error processing task: %v\n", err)
		} else {
			fmt.Println("Task Status after Processing ", b.Status())
			fmt.Println(b.String())
		}
		fmt.Println("=================================================================================================================")
	}
}
