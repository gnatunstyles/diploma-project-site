import styles from '../styles/modalWindow.module.sass'
import {Dispatch, SyntheticEvent, useEffect, useState} from "react";
import {IProj} from "./ProjectsLayout";

export default function ModalEditUser(props: { closeModal: Dispatch<boolean>, email: string, name: string}){

    const [inputedValueEmail, setInputedValueEmail] = useState(props.email);
    const [inputedValueName, setInputedValueName] = useState(props.name);

    useEffect(() => {
        setInputedValueEmail(props.email)
        setInputedValueName(props.name)
    }, [])

    const handleChangeValueEmail = (e: any) => {
        setInputedValueEmail(e.target.value)
    }

    const handleChangeValueName = (e: any) => {
        setInputedValueName(e.target.value)
    }

    const handleSave = async () => {
        const response = await fetch(
            `https://localhost:8000/api/user/update`,
            {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                credentials: "include", //cookie getter
                body: JSON.stringify({
                    new_username: inputedValueName,
                    new_email: inputedValueEmail
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
                    <div>Смена информации пользователя</div>
                    <div className={styles.modalWindowContainerContentLayoutEditUserLayout}>
                        <div className={styles.modalWindowContainerContentLayoutInputLayoutText}>email</div>
                        <input type={"text"} value={inputedValueEmail} onChange={(event) => handleChangeValueEmail(event)}/>
                    </div>
                    <div className={styles.modalWindowContainerContentLayoutEditUserLayout}>
                        <div className={styles.modalWindowContainerContentLayoutInputLayoutText}>UserName</div>
                        <input type={"text"} value={inputedValueName} onChange={(event) => handleChangeValueName(event)}/>
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