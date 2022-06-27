import React, { useState } from "react";
import ModalEditUser from "../components/ModalEditUser";

export default function MenuCard(props: { user: any; mode: any }) {
  const [show, setShow] = useState(false);
  const [hidePannel, setHidePannel] = useState(false);
  const [buttonText, setButtonText] = useState("Подробнее +");
  const [modalEditOpen, setModalEditOpen] = useState(false);

  const onClickItem = () => {
    setShow((prev) => !prev);
    if (!show) {
      setButtonText("Подробнее -");
    } else {
      setButtonText("Подробнее +");
    }
  };

  const onCLickMenu = () => {
    setHidePannel((prev) => !prev);
  };

  const onEditClick = () => {
    setModalEditOpen((prev) => !prev);
  };

  const closeEdit = () => {
    setModalEditOpen(false);
  };

  return (
    <>
      {props.mode && <div>user</div>}
      {!props.mode && (
        <div className="profileCard">
          <div className="cardLayout">
            <div className="cardHeader">
              <b className="cardTitle">Профиль</b>
              <div className="editButton" onClick={onEditClick}>
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
            <div className="userInfo">
              Имя пользователя: {props.user.username}
            </div>
            <div className="moreInfoButton" onClick={onClickItem}>
              {buttonText}
            </div>
            {show && (
              <div className="moreInfoBlock">
                <div className="itemRow">
                  <div className="textRow">E-mail: {props.user.email}</div>
                </div>
                <div className="itemRow">
                  <div className="textRow">
                    Проекты: {props.user.project_number}
                  </div>
                </div>
                <div className="itemRow">
                  <div className="textRow">
                    Места использовано:{" "}
                    {(props.user.used_space / 1073741824).toFixed(4)} GB
                  </div>
                </div>
                <div className="itemRow">
                  <div className="textRow">
                    Места доступно:{" "}
                    {(props.user.available / 1073741824).toFixed(4)} GB
                  </div>
                </div>
              </div>
            )}
            {modalEditOpen && (
              <div>
                <ModalEditUser
                  closeModal={setModalEditOpen}
                  email={props.user.email}
                  name={props.user.username}
                />
              </div>
            )}
          </div>
        </div>
      )}
    </>
  );
}
