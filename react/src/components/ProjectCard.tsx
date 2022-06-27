import { Dispatch, SyntheticEvent, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import styles from "../styles/projectCard.module.sass";
import { IProj } from "./ProjectsLayout";
import Warning from "./Warning";

const BYTES = 1073741824;

export default function ProjectCard(props: {
  proj: IProj;
  setMode: Dispatch<string>;
  openModal: Dispatch<boolean>;
  openEdit: Dispatch<boolean>;

  setSelectedProj: Dispatch<IProj>;
  onClickDelete: (itemId: any, proj_name: string) => void;
  // onClickDownload: (itemId: any, proj_name: string) => void;
}) {
  //     ID: 10,
  //     CreatedAt: "2022-05-26T21:34:18.144414+03:00",
  //     UpdatedAt: "2022-05-26T22:38:07.510554+03:00",
  //     DeletedAt: null,
  //     user_id: 1,
  //     project_name: "asdqwe",
  //     info: "zxczxczxc",
  //     size: 185075623,
  //     link: "http://localhost:1234/projects/1/kekich.html"

  const onClickButtonsProcessing = (e: any) => {
    console.log(e.target.value);
    props.setMode(e.target.value);
    props.setSelectedProj(props.proj);
    props.openModal(true);
  };

  const onClickButtonsEdit = () => {
    props.setSelectedProj(props.proj);
    props.openEdit(true);
  };

  let link = props.proj.link;
  return (
    <div className={styles.item}>
      <div className={styles.itemheaderLayoutName}>
        <div className={styles.itemheaderLayoutName}>
          {props.proj.project_name}
        </div>
        <div
          className={styles.itemheaderLayoutEditButton}
          onClick={onClickButtonsEdit}
        >
          <svg
            version="1.1"
            id="Capa_1"
            xmlns="http://www.w3.org/2000/svg"
            x="0px"
            y="0px"
            viewBox="0 0 469.331 469.331"
          >
            <g>
              <path
                d="M438.931,30.403c-40.4-40.5-106.1-40.5-146.5,0l-268.6,268.5c-2.1,2.1-3.4,4.8-3.8,7.7l-19.9,147.4
		c-0.6,4.2,0.9,8.4,3.8,11.3c2.5,2.5,6,4,9.5,4c0.6,0,1.2,0,1.8-0.1l88.8-12c7.4-1,12.6-7.8,11.6-15.2c-1-7.4-7.8-12.6-15.2-11.6
		l-71.2,9.6l13.9-102.8l108.2,108.2c2.5,2.5,6,4,9.5,4s7-1.4,9.5-4l268.6-268.5c19.6-19.6,30.4-45.6,30.4-73.3
		S458.531,49.903,438.931,30.403z M297.631,63.403l45.1,45.1l-245.1,245.1l-45.1-45.1L297.631,63.403z M160.931,416.803l-44.1-44.1
		l245.1-245.1l44.1,44.1L160.931,416.803z M424.831,152.403l-107.9-107.9c13.7-11.3,30.8-17.5,48.8-17.5c20.5,0,39.7,8,54.2,22.4
		s22.4,33.7,22.4,54.2C442.331,121.703,436.131,138.703,424.831,152.403z"
              />
            </g>
          </svg>
        </div>
      </div>
      <div className={styles.itemDivider} />
      <div className={styles.itemBody}>
        <b>Информация о проекте </b>
        <div className={styles.itemBodyListParams}>
          <div className={styles.itemBodyListParamsListItem}>
            Описание: {props.proj.info}
          </div>
          <div className={styles.itemBodyListParamsListItem}>
            Размер: {(props.proj.size / BYTES).toFixed(3)} GB
          </div>
          <div className={styles.itemBodyListParamsListItem}>
            Количество точек: {props.proj.points}
            {props.proj.points > 10000000 && <Warning />}
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
          <a
            href={`https://localhost:8000/api/projects/download/${props.proj.project_name}`}
            target="_blank"
            download
            type="button"
            className="btn btn-outline-primary w-100 mx-0"
            data-bs-dismiss="modal"
          >
            Download
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
            onClick={(event) => onClickButtonsProcessing(event)}
            value={"random"}
          >
            Processing (Thinning)
          </button>
          <button
            type="button"
            className="btn btn-outline-dark w-100 mx-0"
            data-bs-dismiss="modal"
            onClick={(event) => onClickButtonsProcessing(event)}
            value={"barycenter"}
          >
            Processing (Center gravity)
          </button>
          <button
            type="button"
            className="btn btn-outline-dark w-100 mx-0"
            data-bs-dismiss="modal"
            onClick={(event) => onClickButtonsProcessing(event)}
            value={"candidate"}
          >
            Processing (Candidate center)
          </button>
          <button
            type="button"
            className="btn btn-outline-danger w-100 mx-0"
            data-bs-dismiss="modal"
            onClick={() =>
              props.onClickDelete(props.proj.ID, props.proj.project_name)
            }
          >
            Delete
          </button>
        </div>
      </div>
    </div>
  );
}
