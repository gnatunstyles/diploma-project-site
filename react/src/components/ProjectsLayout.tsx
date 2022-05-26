import { useEffect, useState } from "react";
import { getProjectsOfUser } from "../requests/projectsRequest";
import ProjectCard from "./ProjectCard";
import styles from "../styles/projectsLayout.module.sass";
import MenuCard from "./MenuCard";
import LeftNav from "./LeftNav";
export default function ProjectsLayout(props: { user: any, id: number }) {
  const [projects, setProjects] = useState<Array<object>>([]);

  useEffect(() => {
    (async () => {
      const response = await fetch(
        `http://localhost:8000/api/projects/${props.id}`,
        {
          headers: { "Content-Type": "application/json" },
          credentials: "include", //cookie getter
        }
      );
      const content = await response.json();
      setProjects(content.projects);
    })();
  }, [props.id]);

  const [hidePannel, setHidePannel] = useState(false);

  const onCLickMenu = () => {
    setHidePannel((prev) => !prev);
  };

  const getCards = () => {
    let arr = [];
    if (projects) {
      for (let i = 0; i < projects.length; i++) {
        let currCard = <ProjectCard proj={projects[i]} />;
        arr.push(currCard);
      }
    }
    return arr;
  };

  return (
    <div className='layout'>
      <div
        className="leftPannel"
        style={hidePannel ? { width: 80, minWidth: 80 } : {}}
      >
        <div className="menuButton" onClick={onCLickMenu}>
          Меню
        </div>
        <div>
          <MenuCard mode={hidePannel} user={props.user} />
          <LeftNav mode={hidePannel} />
        </div>
      </div>
      <div className="mainPannel">
        <div className="mainContent">
            <div className={styles.projectsLayout}>
              <div className={styles.projectsLayoutTitle}>Ваши проекты: </div>
              <div className={styles.projectsLayoutList}>{getCards()}</div>
            </div>
        </div>
      </div>
    </div>
  );
}
