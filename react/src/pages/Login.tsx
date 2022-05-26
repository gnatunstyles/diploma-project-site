import React, { SyntheticEvent, useState } from "react";
import { Navigate } from "react-router-dom";
import styles from '../styles/leftMenu.module.sass'

const Login = (props: { setName: (username: string) => void }) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [nav, setNav] = useState(false);

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();
    const response = await fetch("http://localhost:8000/api/sign-in", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include", //cookie getter
      body: JSON.stringify({ email, password }),
    });

    const content = await response.json();
    props.setName(content.username);

    setNav(true);
  };
  if (nav) {
    return <Navigate to={"/"} />;
  }

  return (
    <form  className={styles.wrapper}onSubmit={submit}>
      <h1 className="h3 mb-3 fw-normal" >Please sign in</h1>

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

      <button className="w-100 btn btn-lg btn-primary" type="submit">
        Sign in
      </button>
    </form>
  );
};

export default Login;
