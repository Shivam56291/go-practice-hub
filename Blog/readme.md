<div align="center">

<img src="https://raw.githubusercontent.com/MariaLetta/free-gophers-pack/master/characters/png/9.png" width="120" alt="Blog Gopher">

# 📰 Go MVC Blog Engine

<a href="https://git.io/typing-svg">
  <img src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&weight=600&size=22&pause=1000&color=00ADD8&center=true&vCenter=true&width=500&lines=Service-Repository+Pattern;CLI-Driven+Configuration;Custom+Routing+Engine;Secure+Authentication+Flow" alt="Typing SVG" />
</a>

<br/>

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-DB-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
![Bootstrap](https://img.shields.io/badge/Bootstrap-UI-7952B3?style=for-the-badge&logo=bootstrap&logoColor=white)

<br/>

<a href="#-architecture"><strong>Architecture</strong></a> · <a href="#-roadmap"><strong>Roadmap</strong></a> · <a href="#-setup"><strong>Setup</strong></a>

</div>

---

### 🏗️ Architecture

A strict implementation of the **Service-Repository Pattern** to decouple logic from data.

<div align="center">

```mermaid
graph LR
    Req(Http Request) --> Router[🚦 Router]
    Router --> Handler[🎮 Handler]
    Handler --> Service[💼 Service]
    Service --> Repo[🗄️ Repository]
    Repo --> DB[(MySQL)]
    
    style Router fill:#00ADD8,stroke:#333,color:#fff
    style Service fill:#7952B3,stroke:#333,color:#fff
    style Repo fill:#4479A1,stroke:#333,color:#fff
```
</div>

---

### 🛣️ Module Roadmap

| Phase | Key Concepts | Focus |
| :--- | :--- | :---: |
| **⚙️ Core** | CLI Config, Custom Routing, HTML Engine | `System` |
| **🏛️ Patterns** | Repository Layer, Service Layer, Models | `Arch` |
| **🔐 Auth** | Middleware Guards, Session Logic, Registration | `Security` |
| **🎨 UI/UX** | Form Validation, Error persistence, Flash Msg | `Frontend` |

---

### ✨ Capabilities

| Feature | Description | Status |
| :--- | :--- | :---: |
| **CLI Tools** | `go run . migrate` / `seed` / `serve` | ✅ |
| **Validation** | Form error handling with data persistence. | ✅ |
| **Security** | Route protection via Auth Middleware. | ✅ |
| **Dynamic UI** | Server-side rendering with HTML templates. | ✅ |

---

### 🛠️ Stack

<div align="center">
  <img src="https://skillicons.dev/icons?i=go,mysql,html,css,bootstrap,vscode" alt="Tech Stack" />
</div>

---

### ⚡ Quick Start

1. **Setup & Config**
   ```bash
   git clone <repo>
   # Update config.yml with MySQL creds
  ```

2. **Database Init**
  ```Bash
  go run . migrate
  go run . seed
```

3. **Launch**
   ```Bash
   go run . serve
   ```
<div align="center">

Author: Shivam

<img src="https://raw.githubusercontent.com/MariaLetta/free-gophers-pack/master/characters/png/11.png" width="50" alt="Gopher">

</div>