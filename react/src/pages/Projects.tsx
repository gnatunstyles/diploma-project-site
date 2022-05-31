import React, { SyntheticEvent, useState } from "react";

const Projects = () => {
  const [username, setName] = useState(""); //handle states [{variable}, {function, that changes variable}]

  const openProject = async () => {
    const response = await fetch("http://localhost:1234/examples/lion.html");
    return response;
  };

  return (
    <form onSubmit={openProject}>
      <div className="modal-content rounded-4 shadow">
        <div className="modal-header border-bottom-0">
          <h5 className="modal-title">Project</h5>
        </div>
      </div>
    </form>
  );
};

export default Projects;
