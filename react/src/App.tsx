import React, { useEffect, useState } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import Nav from "./components/Nav";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Projects from "./pages/Projects";
import Register from "./pages/Register";
import Upload from "./pages/Upload";

function App() {

  const [username, setName] = useState(""); //handle states [{variable}, {function, that changes variable}]

  useEffect(() => {
    (async () => {
      const response = await fetch("http://localhost:8000/api/user", {
        headers: { "Content-Type": "application/json" },
        credentials: "include", //cookie getter
      });
      const content = await response.json();
      setName(content.username);
    })();
  });




  return (
    <div className="App">
      <BrowserRouter>
        <Nav username={username} setName={setName}/>
        <main className="form-signin">
          <Routes>
            <Route path="/" element={<Home username={username} />} />
            <Route path="/login" element={<Login setName={setName}/>} />
            <Route path="/register" element={<Register />} />
            <Route path="/projects" element={<Projects/>}/>
            <Route path="/upload" element={<Upload />}/>
          </Routes>
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
