import { SyntheticEvent, useEffect, useState } from "react";
import styles from "../styles/projectCard.module.sass";

export default function ProjectCard(props: { proj: any }) {
  const [projects, setProjects] = useState<Array<object>>([]);

  //     ID: 10,
  //     CreatedAt: "2022-05-26T21:34:18.144414+03:00",
  //     UpdatedAt: "2022-05-26T22:38:07.510554+03:00",
  //     DeletedAt: null,
  //     user_id: 1,
  //     project_name: "asdqwe",
  //     info: "zxczxczxc",
  //     size: 185075623,
  //     link: "http://localhost:1234/projects/1/kekich.html"

  const random = async (e: SyntheticEvent) => {
    const response = await fetch(
      `http://localhost:8000/api/processing/random`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include", //cookie getter
        body: JSON.stringify({
          project_name: props.proj.project_name,
          user_id: props.proj.user_id,
          file_path: props.proj.file_path,
          factor: 10,
        }),
      }
    );
    const content = await response.json();
    setProjects(content.projects);
  };

  const barycenter = async (e: SyntheticEvent) => {
    const response = await fetch(
      `http://localhost:8000/api/processing/barycenter`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include", //cookie getter
        body: JSON.stringify({
          project_name: props.proj.project_name,
          user_id: props.proj.user_id,
          file_path: props.proj.file_path,
          voxel_size: 3,
        }),
      }
    );
    const content = await response.json();
    setProjects(content.projects);
  };

  const candidate = async (e: SyntheticEvent) => {
    const response = await fetch(
      `http://localhost:8000/api/processing/candidate`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include", //cookie getter
        body: JSON.stringify({
          project_name: props.proj.project_name,
          user_id: props.proj.user_id,
          file_path: props.proj.file_path,
          voxel_size: 3,
        }),
      }
    );
    const content = await response.json();
    setProjects(content.projects);
  };

  const deleteProj = async (e: SyntheticEvent) => {
    const response = await fetch(
      `http://localhost:8000/api/projects/delete/${props.proj.project_name}`,
      {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        credentials: "include", //cookie getter
      }
    );
    const content = await response.json();
    setProjects(content.projects);
  };



  let link = props.proj.link;
  return (
    <div className={styles.item}>
      <div className={styles.itemName}>{props.proj.project_name}</div>
      <div className={styles.itemDivider} />
      <div className={styles.itemBody}>
        <b>Информация о проекте </b>
        <div className={styles.itemBodyListParams}>
          <div className={styles.itemBodyListParamsListItem}>
            Описание: {props.proj.info}
          </div>
          <div className={styles.itemBodyListParamsListItem}>
            Размер: {(props.proj.size / 1000000000).toFixed(3)} GB
          </div>
        </div>
        <div className="modal-footer flex-column border-top-0">
          <a
            type="submit"
            href={link}
            target="_blank"
            className="btn btn-outline-primary w-100 mx-0"
            data-bs-dismiss="modal"
          >
            Open
          </a>
          <button
            onClick={() => navigator.clipboard.writeText(`${link}`)}
            type="submit"
            className="btn btn-outline-dark w-100 mx-0"
            data-bs-dismiss="modal"
          >
            Share
          </button>

          <button
            type="button"
            className="btn btn-outline-dark w-100 mx-0"
            data-bs-dismiss="modal"
            onClick={random}

          >
            Processing (Random Sampling)
          </button>
          <button
            type="button"
            className="btn btn-outline-dark w-100 mx-0"
            data-bs-dismiss="modal"
            onClick={barycenter}

          >
            Processing (Grid Barycenter)
          </button>
          <button
            type="button"
            className="btn btn-outline-dark w-100 mx-0"
            data-bs-dismiss="modal"
            onClick={candidate}

          >
            Processing (Grid Candidate)
          </button>
          <button
            type="button"
            className="btn btn-outline-dark w-100 mx-0"
            data-bs-dismiss="modal"
            onClick={deleteProj}
          >
            Delete
          </button>
        </div>
      </div>
    </div>
  );
}
