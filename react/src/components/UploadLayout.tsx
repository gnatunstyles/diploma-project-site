import MenuCard from "./MenuCard";
import LeftNav from "./LeftNav";
import styles from "../styles/uploadLayout.module.sass";
import { useState } from "react";
import Upload from "../pages/Upload";

export default function UploadLayout(props: { user: any; id: number }) {
  const [hidePannel, setHidePannel] = useState(false);
  const onCLickMenu = () => {
    setHidePannel((prev) => !prev);
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
            <div className={styles.uploadLayout}>
              <Upload userid={props.id} />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
