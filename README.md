## **Basic To-Do List (Go CLI Application)**  

### **Project Description**  
This is a simple **command-line-based To-Do List application** written in **Go**. It allows users to create and manage a list of tasks with estimated completion times. Users can **add tasks** interactively and **view all tasks** before exiting the program.  

### **Features:**  
✅ **Add Tasks** – Users can enter a task name and an estimated completion time.  
✅ **View Tasks** – Displays all added tasks in an organized format.  
✅ **Exit Option** – Users can type `"ESC"` to exit the application.  
✅ **Input Cleanup** – Uses `strings.TrimSpace()` to remove unwanted spaces in user input.  

### **How It Works:**  
1. The program starts by displaying a welcome message.  
2. Users are prompted to either continue or exit (`ESC`).  
3. If continuing, they enter:  
   - Task Name  
   - Estimated Time  
4. The task is added to a list and confirmed with a success message.  
5. When the user exits, all tasks are displayed in a structured format.  

### **Technologies Used:**  
🛠 **Golang** – For building the application.  
📜 **bufio.Scanner** – For handling user input efficiently.  
📝 **Slices** – For storing and managing tasks dynamically.  

### **How to Run the To-Do List Program in VS Code (Step-by-Step)**  

#### **Step 1: Install Go**  
If you haven't installed Go yet, download and install it from:  
🔗 [https://go.dev/dl/](https://go.dev/dl/)  

After installation, check if Go is installed correctly by running:  
```sh
go version
```

---

#### **Step 2: Install VS Code**  
Download and install **Visual Studio Code** from:  
🔗 [https://code.visualstudio.com/](https://code.visualstudio.com/)  

---

#### **Step 3: Install the Go Extension in VS Code**  
1. Open VS Code.  
2. Click on the **Extensions** icon (`Ctrl + Shift + X`).  
3. Search for **"Go"** (by the Go Team at Google).  
4. Click **Install** to add Go support.  

---

#### **Step 4: Create a Go Project**  
1. Open **VS Code** and create a new folder (e.g., `todo-list`).  
2. Open the folder in VS Code (`File` > `Open Folder`).  
3. Inside the folder, create a new file:  
   - Click **File > New File**  
   - Name it **`main.go`**  

---

#### **Step 5: Write the Code**  
Copy and paste the **To-Do List** Go code into `main.go`.  

---

#### **Step 6: Run the Program**  
1. Open the **Terminal** in VS Code (`Ctrl + ~` or `View > Terminal`).  
2. Run the program using:  
   ```sh
   go run main.go
   ```
3. Follow the on-screen prompts to add tasks.  
4. Type `"ESC"` to exit and view your task list.  

---
