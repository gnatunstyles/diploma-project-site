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

        <div className="modal-footer flex-column border-top-0">
          <a
            type="submit"
            href="http://localhost:1234/examples/lion.html"
            target="_blank"
            className="btn btn-lg btn-primary w-100 mx-0 mb-2"
            data-bs-dismiss="modal"
          >
            Open
          </a>
          <button
            onClick={() =>
              navigator.clipboard.writeText(
                "http://localhost:1234/examples/lion.html"
              )
            }
            type="submit"
            className="btn btn-lg btn-warning w-100 mx-0"
            data-bs-dismiss="modal"
          >
            Share
          </button>
          <button
            type="button"
            className="btn btn-lg btn-danger w-100 mx-0"
            data-bs-dismiss="modal"
          >
            Delete
          </button>
        </div>
      </div>
    </form>
  );
};

export default Projects;
