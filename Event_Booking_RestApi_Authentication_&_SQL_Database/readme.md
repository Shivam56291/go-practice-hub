<div align="center">

<img src="https://raw.githubusercontent.com/MariaLetta/free-gophers-pack/master/characters/png/13.png" width="120" alt="Event Gopher">

# 🎟️ Event Booking API

<a href="https://git.io/typing-svg">
  <img src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&weight=600&size=22&pause=1000&color=00C7B7&center=true&vCenter=true&width=500&lines=Secure+Event+Management;JWT+Authentication+Flow;Production-Ready+REST+API" alt="Typing SVG" />
</a>

<br/>

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00C7B7?style=for-the-badge&logo=gin&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-Auth-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-DB-003B57?style=for-the-badge&logo=sqlite&logoColor=white)

<br/>

<a href="#-features"><strong>Features</strong></a> · <a href="#-architecture"><strong>Architecture</strong></a> · <a href="#-setup"><strong>Setup</strong></a>

</div>

---

### 📖 Overview

This project implements a **production-style RESTful API** for managing events and user registrations. It focuses on the intersection of **Security**, **Scalability**, and **Modular Design**.



---

### 🏗️ System Architecture

The API follows a clean, layered approach to separate concerns and ensure maintainability.

<div align="center">

```mermaid
graph LR
    User([📱 Client]) -->|JSON| API[🚀 Gin Server]
    
    subgraph "Logic Layers"
        API --> Middleware{🛡️ Auth Guard}
        Middleware --> Handlers[🎮 Handlers]
        Handlers --> DB[(🗄️ Database)]
    end
    
    subgraph "Security"
        Middleware -.-> JWT[🔐 JWT Verify]
    end

    style API fill:#00ADD8,stroke:#333,stroke-width:2px,color:#fff
    style DB fill:#336791,stroke:#333,stroke-width:2px,color:#fff
    style Middleware fill:#00C7B7,stroke:#333,stroke-width:2px,color:#fff
```

---

### ✨ Core Capabilities

| Feature | Description | Status |
| :--- | :--- | :---: |
| **🔐 Auth Flow** | User signup/login with Bcrypt hashing & JWT. | ✅ |
| **📅 Event CRUD** | Create, View, Update, and Delete events. | ✅ |
| **🛡️ Protected** | Only event creators can edit/delete their data. | ✅ |
| **🧾 Registration** | Multi-user event participation logic. | ✅ |

---

### 🛠️ Tech Stack Icons

<div align="center">
  <img src="https://skillicons.dev/icons?i=go,sqlite,postman,git,vscode" alt="Tech Stack" />
</div>

---

### ⚡ Installation & Setup

1. **Clone & Navigate**
   ```bash
   cd event-booking-api
   ```
   
2. **Install Dependencies**
   ```bash
   go mod download
   ```
   
3. **Run the Application**
   ```bash
   go run main.go
   ```

<div align="center">

Author: Shivam Building scalable backends, one line at a time.

<img src="https://www.google.com/search?q=https://raw.githubusercontent.com/MariaLetta/free-gophers-pack/master/characters/png/7.png" width="60" alt="Gopher">

</div>