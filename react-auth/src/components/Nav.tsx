import React, { useState } from "react";
import { Link } from "react-router-dom";

const Nav = (props: {
  username: string,
  setName: (username: string) => void
}) => {

  const logout = async () => {
    await fetch("http://localhost:8000/api/user/logout", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });

    props.setName("");
  };

  let menu;

  if (props.username === undefined) {
    menu = (
      <ul className="navbar-nav me-auto mb-2 mb-md-0">
        <li className="nav-item">
          <Link to={"/login"} className="nav-link active" aria-current="page">
            Вход
          </Link>
        </li>

        <li className="nav-item">
          <Link
            to={"/register"}
            className="nav-link active"
            aria-current="page"
          >
            Регистрация
          </Link>
        </li>
      </ul>
    );
  } else {
    menu = (
      <ul className="navbar-nav me-auto mb-2 mb-md-0">
        
        <li className="nav-item">
          <Link to={"/upload"} className="nav-link">
            Загрузить
          </Link>
        </li>

        <li className="nav-item">
          <Link to={"/projects"} className="nav-link">
            Проекты
          </Link>
        </li>

        <li className="nav-item">
          <Link to={"/login"} className="nav-link active" onClick={logout}>
            Выйти
          </Link>
        </li>
      </ul>
    );
  }

  return (
    <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
      <div className="container-fluid">
        <Link to={"/"} className="navbar-brand">
          Potree Visualizer
        </Link>
        <div>{menu}</div>
      </div>
    </nav>
  );
};

export default Nav;
