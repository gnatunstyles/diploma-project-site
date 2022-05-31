import {Dispatch, SyntheticEvent, useEffect, useState} from "react";
import styles from "../styles/projectCard.module.sass";
import {IProj} from "./ProjectsLayout";

const MILLIARD = 1000000000;

export default function ProjectCard(props: {
    proj: IProj,
    setMode: Dispatch<string>,
    openModal: Dispatch<boolean> ,
    setSelectedProj: Dispatch<IProj>
    onClickDelete: (itemId: any, proj_name: string) => void
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

  const onClickButtonsProcessing = (e : any) => {
      console.log(e.target.value);
      props.setMode(e.target.value);
      props.setSelectedProj(props.proj)
      props.openModal(true)
  }

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
            Размер: {(props.proj.size / MILLIARD).toFixed(3)} GB
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
            onClick={(event) => onClickButtonsProcessing(event)}
            value={'random'}

          >
            Processing (Random Sampling)
          </button>
          <button
            type="button"
            className="btn btn-outline-dark w-100 mx-0"
            data-bs-dismiss="modal"
            onClick={(event) => onClickButtonsProcessing(event)}
            value={'barycenter'}

          >
            Processing (Grid Barycenter)
          </button>
          <button
            type="button"
            className="btn btn-outline-dark w-100 mx-0"
            data-bs-dismiss="modal"
            onClick={(event) => onClickButtonsProcessing(event)}
            value={'candidate'}

          >
            Processing (Grid Candidate)
          </button>
          <button
            type="button"
            className="btn btn-outline-dark w-100 mx-0"
            data-bs-dismiss="modal"
            onClick={() => props.onClickDelete(props.proj.ID, props.proj.project_name)}
          >
            Delete
          </button>
        </div>
      </div>
    </div>
  );
}
