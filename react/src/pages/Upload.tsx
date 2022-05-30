import React, { SyntheticEvent, useState } from "react";

const Upload = () => {
  const [username, setName] = useState(""); //handle states [{variable}, {function, that changes variable}]

  const upload = async () => {
    const response = await fetch("http://localhost:8000/api/users/upload", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({ username }),
    });
    return response;
  };

  return (
    <form onSubmit={upload}>
      <input type="file"/>
      <button
        type="submit" name="cloud"
        className="btn btn-lg btn-info w-100 mx-0 mt-4 mb-2"
        onClick={upload}
      >
        Send
      </button>
    </form>
  );
};
export default Upload;
