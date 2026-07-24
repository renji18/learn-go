/*
You are building a Workflow Execution Engine.

A workflow consists of multiple tasks.

Each task:

* has an ID
* has a type
* has a status
* may depend on other tasks

⸻

⚙️ Functional Requirements

1. Task Interface

Each task must:

* Execute itself
* Report status
* Expose dependencies

⸻

2. Dependency System (this is the heart)

A task can depend on other tasks.

Example:
Task B depends on Task A
Task C depends on Task B

Rules:

* A task cannot execute until ALL dependencies are COMPLETED
* If any dependency FAILS → dependent task must NOT run
* Detect invalid dependency references

⸻

3. Workflow Manager

You need a central system that:

* Stores all tasks
* Knows dependencies
* Executes tasks in correct order

⸻

Required operations:

* AddTask(task, dependencies)
* Run(taskID)
* RunAll()
* ListTasks()

⸻

4. Execution Behavior

Scenario:
A → B → C

Running RunAll() should:

1. Execute A
2. Then B
3. Then C

⸻

If A fails:

* B should NOT run
* C should NOT run

⸻

5. Cycle Detection (this is where you’ll struggle)

This is mandatory.

If someone creates:
A → B → C → A

You must detect it and return an error.

No execution should happen.

⸻

6. State Model

At minimum:

* Pending
* Running
* Completed
* Failed
* Skipped (important — think why)

⸻

7. Pointer Discipline (non-negotiable)

If your system accidentally copies tasks:

👉 dependency state WILL break

⸻

8. String Parsing (Chapter 7 teaser — harder now)

Input format:
email|task1|user@test.com|Hello|
file|task2|/tmp/file.txt|data|task1
whatsapp|task3|999999|Hi|task2

Last field = comma-separated dependencies

⸻

You must:

* Parse input
* Build tasks
* Wire dependencies

⸻

9. Error Discipline

You must handle:

* Unknown task type
* Duplicate task ID
* Missing dependency
* Circular dependency
* Executing already completed task
* Executing task with failed dependency

⸻

10. JS Filter (this will break you if you cheat)

If you:

* store dependencies as raw strings everywhere
* skip modeling relationships
* use loose maps everywhere

👉 your system will collapse under cycles and execution logic

⸻

🧠 Hidden Complexity (this is intentional)

You will be forced to think about:

* graph traversal (you don’t know it yet — good)
* execution ordering
* failure propagation
* state consistency

⸻

✅ Definition of Done

You are NOT done until all of these work:

⸻

✔️ Case 1 — Linear execution
A → B → C

All run in order.

⸻

✔️ Case 2 — Branching
A → B
A → C

B and C both wait for A.

⸻

✔️ Case 3 — Failure propagation

If A fails:

* B = Skipped
* C = Skipped

⸻

✔️ Case 4 — Cycle detection
A → B → C → A

system rejects BEFORE execution.

⸻

✔️ Case 5 — Partial execution

Running only Run("B") should:

* run A first
* then B

⸻

✔️ Case 6 — Idempotency

Running already completed workflow should not re-execute tasks.

⸻

✔️ Case 7 — State correctness

At any time:

* No task jumps states illegally
* No task runs twice
* No task runs without dependencies satisfied
*/

package main

import (
	"fmt"
	"slices"
	"strings"
)

// This type is used giving numeric values to the below constants
type Status int

const (
	PENDING Status = iota + 1
	RUNNING
	COMPLETED
	FAILED
	SKIPPED
)

// This stringer fn returns string values for the status
func (s Status) String() string {
	switch s {
	case PENDING:
		return "Pending"
	case RUNNING:
		return "Processing"
	case COMPLETED:
		return "Completed"
	case SKIPPED:
		return "Skipped"
	default:
		return "Failed"
	}
}

type Task interface {
	Id() string                       // The id method returns the 'name' of the task
	Execute() error                   // This execute method works as an internal private helper method. This method is to called from inside the Run() or RunAll() method. Calling this method directly on a task should be avoided. This is something of like an override function.
	Status() Status                   // The status method returns the status of the task
	Dependencies() []string           // The dependencies method returns the list of dependency tasks and the number of dependency tasks. Here, I thought of having []string which will store the name of the dependency tasks, the only reason I did not proceed ahead with this manner, is because then the Dependencies method would require needing workflow, to fetch a task using name. Or what I can do is maintain a local var which stores the same workflow map or replace it.
	ChangeStatus(status Status) error // This is a direct helper method to manipulate the status of a task. This is to be used safely as it overwrites the status of the task without any validations.
	fmt.Stringer                      // The stringer method returns a string representation about a task.
}

// The base struct. Here as well, I would like to explore other ways to store the task dependency slice as storing multiple tasks is not going to be scalable.
type BaseTask struct {
	name          string
	status        Status
	dependencyIds []string
}

// Extends the BaseTask
type EmailTask struct {
	BaseTask
	Receiver string
	Message  string
}

// Extends the BaseTask
type FileTask struct {
	BaseTask
	Path    string
	Content string
}

// The workflow manager to maintain a map of task name -> task
type Workflow struct {
	tasks     map[string]Task
	taskOrder []string
}

// A factory type which makes and returns a Task
type TaskFactory func(baseTask BaseTask, mainContent string, subContent string) (Task, error)

// The actual factory which returns a generic Task based on the type of the task. This helps in generating the Task for different types of task, instead of doing if else of switch case.
var taskFactories map[string]TaskFactory = map[string]TaskFactory{
	"email": func(baseTask BaseTask, mainContent, subContent string) (Task, error) {

		// custom validation for email type. If not satisfied, mark task as failed and return error
		if !strings.Contains(mainContent, "@") {
			return nil, fmt.Errorf("Invalid receiver. Task %s not created!!!", mainContent)
		}

		return &EmailTask{BaseTask: baseTask, Receiver: mainContent, Message: subContent}, nil
	},
	"file": func(baseTask BaseTask, mainContent, subContent string) (Task, error) {
		if mainContent == "" {
			return nil, fmt.Errorf("Invalid path. Task %s not created!!!", mainContent)
		}

		return &FileTask{BaseTask: baseTask, Path: mainContent, Content: subContent}, nil
	},
}

// EmailTask methods
func (e EmailTask) Id() string {
	return e.name
}

func (e *EmailTask) Execute() error {
	fmt.Printf("========\nStarting with main task %s...\n", e.name)

	if e.status == COMPLETED {
		// We don't return error here, as dependency cycle needs it. A depends on B, B is completed -> returns error, thus A gets Skipped. Thus we return nil
		fmt.Printf("Task %s already Completed!!!\n", e.name)
		return nil
	}

	// Return error on failed tasks
	if e.status == FAILED {
		return fmt.Errorf("Task %s already Failed!!!", e.name)
	}

	// Return error on skipped tasks
	if e.status == SKIPPED {
		return fmt.Errorf("A dependency of task %s failed, Skipping!!!", e.name)
	}

	fmt.Printf("Executing task %s...\n", e.name)
	e.status = RUNNING
	e.status = COMPLETED
	fmt.Println("Execution complete!!!\n========")

	return nil
}

func (e EmailTask) Status() Status {
	return e.status
}

func (e EmailTask) Dependencies() []string {
	if len(e.dependencyIds) <= 0 {
		return []string{}
	}

	return e.dependencyIds
}

func (e *EmailTask) ChangeStatus(status Status) error {
	e.status = status
	return nil
}

func (e EmailTask) String() string {
	dependencyIds := e.Dependencies()

	var builder strings.Builder

	if len(dependencyIds) > 0 {
		builder.WriteString("[ ")

		for i, id := range dependencyIds {
			builder.WriteString(id)
			if i != len(dependencyIds)-1 {
				builder.WriteString(", ")
			}
		}

		builder.WriteString(" ]")
	}

	dependencyString := builder.String()

	return fmt.Sprintf("Task Type: (email)\nTask Name: (%s)\nTask Status: (%s)\nEmail Receiver: (%s)\nEmail Message: (%s)\nDependencies: (%s)", e.name, e.Status().String(), e.Receiver, e.Message, string(dependencyString))
}

// EmailTask methods
func (f FileTask) Id() string {
	return f.name
}

func (f *FileTask) Execute() error {
	fmt.Printf("========\nStarting with main task %s...\n", f.name)

	if f.status == COMPLETED {
		// We don't return error here, as dependency cycle needs it. A depends on B, B is completed -> returns error, thus A gets Skipped. Thus we return nil
		fmt.Printf("Task %s already Completed!!!\n", f.name)
		return nil
	}

	if f.status == FAILED {
		return fmt.Errorf("Task %s already Failed!!!", f.name)
	}

	if f.status == SKIPPED {
		return fmt.Errorf("A dependency of task %s failed, Skipping!!!", f.name)
	}

	fmt.Printf("Executing task %s...\n", f.name)
	f.status = RUNNING
	f.status = COMPLETED
	fmt.Println("Execution complete!!!\n========")

	return nil
}

func (f FileTask) Status() Status {
	return f.status
}

func (f FileTask) Dependencies() []string {
	if len(f.dependencyIds) <= 0 {
		return []string{}
	}

	return f.dependencyIds
}

func (f *FileTask) ChangeStatus(status Status) error {
	f.status = status
	return nil
}

func (f FileTask) String() string {
	dependencyIds := f.Dependencies()

	var builder strings.Builder

	if len(dependencyIds) > 0 {
		builder.WriteString("[ ")

		for i, id := range dependencyIds {
			builder.WriteString(id)
			if i != len(dependencyIds)-1 {
				builder.WriteString(", ")
			}
		}

		builder.WriteString(" ]")
	}

	dependencyString := builder.String()

	return fmt.Sprintf("Task Type: (file)\nTask Name: (%s)\nTask Status: (%s)\nFile Path: (%s)\nFile Content: (%s)\nDependencies: (%s)", f.name, f.Status(), f.Path, f.Content, dependencyString)
}

// Helder function to test cycle. NOTE: My architecture does not allow cycles to exist at all. This is just for learning purpose.
func (wf Workflow) CheckCycle(dependencies []string, checkFor string) bool {
	exists := slices.Contains(dependencies, checkFor)

	if exists {
		return true
	}

	for _, dependencyName := range dependencies {
		task := wf.tasks[dependencyName]

		internalTaskDependencies := task.Dependencies()

		cycleDetected := wf.CheckCycle(internalTaskDependencies, checkFor)

		if cycleDetected {
			return true
		}
	}

	return false
}

// Workflow methods
func (wf *Workflow) AddTask(inputString string) error {
	// This divides the input string in 5 parts on "|"
	inputArr := strings.SplitN(inputString, "|", 5)

	// The input string is valid if the number of parts is either 4 or 5. Any other count makes the input string invalid.
	if len(inputArr) < 4 || len(inputArr) > 5 {
		return fmt.Errorf("Invalid input string (%s)", inputString)
	}

	taskType := inputArr[0]
	taskName := inputArr[1]
	taskMainContent := inputArr[2]
	taskSubContent := inputArr[3]
	taskDependencies := []string{}

	if taskName == "" || taskMainContent == "" {
		return fmt.Errorf("Invalid input string (%s)", inputString)
	}

	// Validation to not allow creating tasks with same name.
	_, has := wf.tasks[taskName]

	if has {
		return fmt.Errorf("Task with name %s already exists: input string (%s)", taskName, inputString)
	}

	if len(inputArr) == 5 {
		// check for circular dependency
		// How to -> I don't know graph traversal, thus I don't know how to use or implement it. I have no knowledge on how to manage cycles. I do have some ideas, but each involves the Task to touch the workfactory which I don't think would be correct architecture.

		// Get the dependency tasks
		for dependencyName := range strings.SplitSeq(inputArr[4], ",") {
			_, has := wf.tasks[dependencyName]
			if !has {
				return fmt.Errorf("Dependency task (%s) not found", dependencyName)
			}

			taskDependencies = append(taskDependencies, dependencyName)
		}
	}

	// validate if the task type is allowed, and also get it's factory fn which returns a task
	factory, has := taskFactories[taskType]

	if !has {
		return fmt.Errorf("Unknown task type provided: %s", taskType)
	}

	// validate if no cycles exist. NOTE: Our current architecture does not allow creating cycles at all. This block is just for learning purpose.
	cycleDetected := wf.CheckCycle(taskDependencies, taskName)

	if cycleDetected {
		return fmt.Errorf("Cycle detected in one of the dependencies. Task not created!!!")
	}

	// The common base Task
	baseTask := BaseTask{name: taskName, status: PENDING, dependencyIds: taskDependencies}

	task, err := factory(baseTask, taskMainContent, taskSubContent)

	if err != nil {
		return fmt.Errorf("Task %s not created: %w", taskName, err)
	}

	wf.tasks[taskName] = task

	wf.taskOrder = append(wf.taskOrder, taskName)

	fmt.Printf("Task (%s) added successfully\n", taskName)

	return nil
}

// Takes the name of the task to run, validate if task exists or not. And then call the execute fn.
// The execute fn internally executes any/all dependencies, if possible.
// This is the actual function which follows the project requirement to execute dependencies of a task and then execute the actual task.
func (wf *Workflow) Run(taskName string) error {
	task, has := wf.tasks[taskName]

	if !has {
		return fmt.Errorf("Task %s not found", taskName)
	}

	dependencies := task.Dependencies()

	// has dependencies
	if len(dependencies) > 0 {
		for _, dependency := range dependencies {
			err := wf.Run(dependency)

			if err != nil {
				// If the execution comes to this block, it means that a dependency of a task has failed.
				// The execute method on a task takes care of the following statues -> Failed (if a task has failed independently) & [Running, Completed] if a task has completed successfully.
				// The only status which won't be handled inside the Execute fn now is the Skipped status.
				// This status is in the scenario that a task has failed, and that task was a dependency, so now after error in that dependency, the flow has reached in this block, where we return error that the dependency x of task y has failed.
				// Now, before returning we need to skip this task. Because if not, then this task will continue to remain in Pending.
				// But we can not manipulate the status of a task from the workFlow. Thus, I'll be creating a new method on the Task interface, specifically for updating the status of the task to Skipped. If needed in future, then we'll update the method to accept parameter and update the status as provided.
				skipErr := task.ChangeStatus(SKIPPED)

				if skipErr != nil {
					return fmt.Errorf("Error skipping task %s: %w", taskName, skipErr)
				}

				return fmt.Errorf("Error during execution of dependency %s of task %s: %w", dependency, taskName, err)
			}
		}
	}

	err := task.Execute()

	if err != nil {
		return fmt.Errorf("Error during execution of task %s: %w", taskName, err)
	}

	return nil
}

// Loops over all the tasks, and calls the execute method on each task.
func (wf *Workflow) RunAll() error {
	for _, taskName := range wf.taskOrder {
		task, has := wf.tasks[taskName]

		if !has {
			return fmt.Errorf("Task %s not found", taskName)
		}

		err := wf.Run(task.Id())

		if err != nil {
			// Here, we don't return because if we encounter error in task n-1 then the nth task won't execute as the flow of execution has gone from this function.
			fmt.Printf("Error during execution of task %s: %v\n", task.Id(), err)
		}
	}

	return nil
}

// Returns a slice of Tasks for UI representation in the terminal.
func (wf Workflow) ListTasks() []Task {
	tasks := make([]Task, 0)

	for _, taskName := range wf.taskOrder {
		task := wf.tasks[taskName]

		tasks = append(tasks, task)
	}

	return tasks
}

func main() {
	// Initializing the workflow
	wf := Workflow{
		tasks: make(map[string]Task),
	}

	err := wf.AddTask("email|task1|user@email.com|Hello")

	if err != nil {
		fmt.Println(err)
	}

	err = wf.AddTask("email|task2|user@email.com|Hello")

	if err != nil {
		fmt.Println(err)
	}

	err = wf.AddTask("file|task3|/user/dustbin/bashing|Hello from file|task1,task2")

	if err != nil {
		fmt.Println(err)
	}

	err = wf.AddTask("file||/user/dustbin/bashing|Hello from file|task1,task2")

	if err != nil {
		fmt.Println(err)
	}

	err = wf.AddTask("email|task4|/user/dustbin/bashing|Hello from file|task1,task2")

	if err != nil {
		fmt.Println(err)
	}

	err = wf.AddTask("whatsapp|task5|/user/dustbin/bashing|Hello from file|task1,task2")

	if err != nil {
		fmt.Println(err)
	}

	err = wf.AddTask("email|task5|/user/dustbin/bashing|Hello from file|task1,task2")

	if err != nil {
		fmt.Println(err)
	}

	err = wf.AddTask("file|task5|/user/dustbin/bashing|Hello from file|task1,task2")

	if err != nil {
		fmt.Println(err)
	}

	tasks := wf.ListTasks()

	fmt.Println()
	fmt.Println("==========ALL TASKS==========")
	for _, task := range tasks {
		fmt.Println(task.String())
		fmt.Println()
	}
	fmt.Println("==========DONE PRINTING==========")

	if len(tasks) > 0 {
		task1 := tasks[0]
		err = wf.Run(task1.Id())

		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println()
	fmt.Println("==========ALL TASKS AFTER EXECUTING A SINGLE TASK==========")
	for _, task := range tasks {
		fmt.Println(task.String())
		fmt.Println()
	}
	fmt.Println("==========DONE PRINTING==========")

	fmt.Println()
	fmt.Println("==========RUNNING ALL TASKS==========")
	err = wf.RunAll()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println("==========ALL TASKS AFTER EXECUTING ALL TASKS==========")
	for _, task := range tasks {
		fmt.Println(task.String())
		fmt.Println()
	}
	fmt.Println("==========DONE PRINTING==========")
}
