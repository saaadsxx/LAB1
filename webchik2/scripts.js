document.addEventListener("DOMContentLoaded", () => {
    const addTaskButton = document.getElementById("addTaskButton");
    const taskDialog = document.getElementById("taskDialog");
    const taskInput = document.getElementById("taskInput");
    const confirmAddTask = document.getElementById("confirmAddTask");

    const toDoColumn = document.getElementById("toDo").querySelector("ul");
    const inProgressColumn = document.getElementById("inProgress").querySelector("ul");
    const completedColumn = document.getElementById("completed").querySelector("ul");

    addTaskButton.addEventListener("click", () => {
        taskDialog.showModal();
    });

    taskDialog.addEventListener("click", (e) => {
        if (e.target === taskDialog) taskDialog.close();
    });

    confirmAddTask.addEventListener("click", (e) => {
        e.preventDefault();
        const taskText = taskInput.value.trim();
        if (taskText) {
            addTask(taskText);
            taskInput.value = "";
            taskDialog.close();
        }
    });

    function addTask(taskText) {
        const taskItem = createTaskItem(taskText, moveToNextColumn, deleteTask);
        toDoColumn.appendChild(taskItem);
    }

    function createTaskItem(text, onStatusChange, onDelete) {
        const li = document.createElement("li");
        li.textContent = text;

        const marker = document.createElement("button");
        marker.textContent = "✔️";
        marker.addEventListener("click", () => onStatusChange(li));
        li.appendChild(marker);

        const deleteButton = document.createElement("button");
        deleteButton.textContent = "🗑️";
        deleteButton.addEventListener("click", () => onDelete(li));
        li.appendChild(deleteButton);

        return li;
    }

    function moveToNextColumn(task) {
        if (task.parentElement === toDoColumn) {
            inProgressColumn.appendChild(task);
        } else if (task.parentElement === inProgressColumn) {
            completedColumn.appendChild(task);
            task.querySelector("button").remove(); // Убираем кнопку "✔️"
        }
    }

    function deleteTask(task) {
        task.remove();
    }
});











