import { SyntheticEvent, useEffect, useState } from "react";
import { getProjectsOfUser } from "../requests/projectsRequest";
import ProjectCard from "./ProjectCard";
import styles from "../styles/projectsLayout.module.sass";
import MenuCard from "./MenuCard";
import LeftNav from "./LeftNav";
import ModalVindow from "./ModalWindow";

export interface IProj {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string;
  user_id: number;
  project_name: string;
  info: string;
  size: number;
  file_path: string;
  link: string;
}

const getInitialState = {
  ID: 0,
  CreatedAt: "",
  UpdatedAt: "",
  DeletedAt: "",
  user_id: 0,
  project_name: "",
  info: "",
  size: 0,
  file_path: "",
  link: "",
};

export default function ProjectsLayout(props: { user: any; id: number }) {
  const [projects, setProjects] = useState<IProj[]>([]);
  const [hidePannel, setHidePannel] = useState(false);

  const [searchValue, setSearchValue] = useState("");
  const [submittedValue, setSubmittedValue] = useState("");

  const [showModal, setShowModal] = useState<boolean>(false);
  const [mode, setMode] = useState("none");

  const [selectedProj, setSelectedProj] = useState<IProj>(getInitialState);

  useEffect(() => {
    (async () => {
      console.log(searchValue, props.id);
      if (submittedValue === "") {
        const response = await fetch(
          `http://localhost:8000/api/projects/${props.id}`,
          {
            headers: { "Content-Type": "application/json" },
            credentials: "include", //cookie getter
          }
        );
        const content = await response.json();
        setProjects(content.projects);
      } else {
        const response = await fetch(
          `http://localhost:8000/api/projects/find/${props.id}/${submittedValue}`,
          {
            headers: { "Content-Type": "application/json" },
            credentials: "include", //cookie getter
          }
        );
        const content = await response.json();
        setProjects(content.projects);
      }
    })();
  }, [props.id, submittedValue]);

  const handleSearchValueChange = (event: any) => {
    setSearchValue(event.target.value);
  };

  const onClickSubmitButton = () => {
    setSubmittedValue(searchValue);
  };

  const onCLickMenu = () => {
    setHidePannel((prev) => !prev);
  };

  const getCards = () => {
    let arr = [];
    if (projects) {
      for (let i = 0; i < projects.length; i++) {
        let currCard = (
          <ProjectCard
            proj={projects[i]}
            setSelectedProj={setSelectedProj}
            onClickDelete={onDeleteItem}
            setMode={setMode}
            openModal={setShowModal}
          />
        );
        arr.push(currCard);
      }
    }
    return arr;
  };

  const onDeleteItem = (itemId: any, proj_name: string) => {
    let newArr = [];
    console.log(itemId);
    newArr = projects.filter((item: any) => item.ID !== +itemId);
    console.log(newArr);
    setProjects(newArr);
    deleteProj(proj_name);
  };

  const deleteProj = async (proj_name: string) => {
    const response = await fetch(
      `http://localhost:8000/api/projects/delete/${proj_name}`,
      {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        credentials: "include", //cookie getter
      }
    );
    const content = await response.json();
    //setProjects(content.projects);
  };

  return (
    <div className="layout">
      <div
        className="leftPannel"
        style={hidePannel ? { width: 80, minWidth: 80 } : {}}
      >
        <div className="mainPadding">
          <div className="menuButton" onClick={onCLickMenu}>
            Меню
          </div>
          <div>
            <MenuCard mode={hidePannel} user={props.user} />
            <LeftNav mode={hidePannel} />
          </div>
        </div>
      </div>
      <div className="mainPannel">
        <div className="mainPadding">
          <div className="mainContent">
            {showModal && (
              <ModalVindow
                mode={mode}
                closeModal={setShowModal}
                proj={selectedProj}
              />
            )}
            <div className={styles.projectsLayout}>
              <div className={styles.projectsLayoutTitleLayout}>
                <div className={styles.projectsLayoutTitleLayoutText}>
                  Ваши проекты:{" "}
                </div>
                <input
                  placeholder={"Find project..."}
                  onChange={(event) => handleSearchValueChange(event)}
                  value={searchValue}
                />
                <button onClick={onClickSubmitButton}>Search</button>
              </div>
              <div className={styles.projectsLayoutList}>{getCards()}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
