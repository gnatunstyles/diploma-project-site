import styles from '../styles/modalWindow.module.sass'
import {Dispatch, SyntheticEvent, useEffect, useState} from "react";
import {IProj} from "./ProjectsLayout";

export default function ModalEditProject(props: { closeModal: Dispatch<boolean>, proj: IProj}){

    const [inputedValueInfo, setInputedValueInfo] = useState(props.proj.info);
    const [inputedValueName, setInputedValueName] = useState(props.proj.project_name);

    useEffect(() => {
        setInputedValueInfo(props.proj.info)
        setInputedValueName(props.proj.project_name)
    }, [])

    const handleChangeValueEmail = (e: any) => {
        setInputedValueInfo(e.target.value)
    }

    const handleChangeValueName = (e: any) => {
        setInputedValueName(e.target.value)
    }

    const handleSave = async () => {
        const response = await fetch(
            `https://localhost:8000/api/projects/update/${props.proj.project_name}`,
            {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                credentials: "include", //cookie getter
                body: JSON.stringify({
                    new_project_name: inputedValueName,
                    new_info: inputedValueInfo
                }),
            }
        );
        const content = await response.json();
        props.closeModal(false)
        window.location.reload()
        //setProjects(content.projects);
    };

    return (
        <div className={styles.modalWindow}>
            <div className={styles.modalWindowContainer}>
                <div className={styles.modalWindowContainerContentLayout}>
                    <div>Смена информации проекта</div>
                    <div className={styles.modalWindowContainerContentLayoutEditUserLayout}>
                        <div className={styles.modalWindowContainerContentLayoutInputLayoutText}>Название</div>
                        <input type={"text"} value={inputedValueName} onChange={(event) => handleChangeValueName(event)}/>
                    </div>
                    <div className={styles.modalWindowContainerContentLayoutEditUserLayout}>
                        <div className={styles.modalWindowContainerContentLayoutInputLayoutText}>Инфо</div>
                        <input type={"text"} value={inputedValueInfo} onChange={(event) => handleChangeValueEmail(event)}/>
                    </div>
                </div>
                <div className={styles.modalWindowContainerButtonsLayout}>
                    <button className={styles.modalWindowContainerButtonsLayoutButtonChancel}
                            onClick={() => props.closeModal(false)}
                    >Close</button>
                    <button className={styles.modalWindowContainerButtonsLayoutButtonProcess}
                            onClick={handleSave}
                    >Save</button>
                </div>
            </div>
        </div>
    )
}