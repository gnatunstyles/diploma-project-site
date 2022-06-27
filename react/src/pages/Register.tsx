import React, { SyntheticEvent, useState } from "react";
import { Navigate } from "react-router-dom";
import styles from '../styles/leftMenu.module.sass'

const Register = () => {
  const [username, setName] = useState(""); //handle states [{variable}, {function, that changes variable}]
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [nav, setNav] = useState(false);

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();
    const creds = btoa(`${email}:${password}:${username}`);
    await fetch("https://localhost:8000/api/sign-up", {
      method: "POST",
      headers: {
        "Authorization": `Basic ${creds}`,
      },
    });
    setNav(true);
  };
  
  if (nav) {
    return <Navigate to={"/login"} />;
  }

  return (
    <form  className={styles.wrapper} onSubmit={submit}>
      <h1 className="h1 mb-4 mt-4 fw-normal">Please sign up</h1>

      <div className="form-floating">
        <input
          type="text"
          className="form-control"
          placeholder="John Doe"
          onChange={(e) => setName(e.target.value)}
        />
        <label htmlFor="floatingInput">Username</label>
      </div>

      <div className="form-floating">
        <input
          type="email"
          className="form-control"
          placeholder="name@example.com"
          onChange={(e) => setEmail(e.target.value)}
        />
        <label htmlFor="floatingInput">Email address</label>
      </div>

      <div className="form-floating">
        <input
          type="password"
          className="form-control"
          placeholder="Password"
          onChange={(e) => setPassword(e.target.value)}
        />
        <label htmlFor="floatingPassword">Password</label>
      </div>
      <button className="w-100 btn btn-lg mt-3 btn-primary" type="submit">
        Sign up
      </button>
    </form>
  );
};

export default Register;
