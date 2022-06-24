import styles from '../styles/modalWindow.module.sass'
import {Dispatch, SyntheticEvent, useState} from "react";
import {IProj} from "./ProjectsLayout";

export default function ModalVindow(props: {mode: string, closeModal: Dispatch<boolean>, proj: IProj}){

    const [inputedValue, setInputedValue] = useState(0);
    const [error, setError] = useState(false);
    const [errorText, setErrorText] = useState('')

    const getNameOfInput = () => {
        if(props.mode === 'random'){
            return 'Decimation: ';
        } else {
            return 'Voxel size: ';
        }
    }

    const random = async () => {
        alert(props.proj.project_name)
        const response = await fetch(
            `https://localhost:8000/api/processing/random`,
            {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                credentials: "include", //cookie getter
                body: JSON.stringify({
                    project_name: props.proj.project_name,
                    user_id: props.proj.user_id,
                    file_path: props.proj.file_path,
                    factor: inputedValue,
                }),
            }
        );
        const content = await response.json();
        window.location.reload()

    };

    const barycenter = async () => {
        const response = await fetch(
            `https://localhost:8000/api/processing/barycenter`,
            {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                credentials: "include", //cookie getter
                body: JSON.stringify({
                    project_name: props.proj.project_name,
                    user_id: props.proj.user_id,
                    file_path: props.proj.file_path,
                    voxel_size: inputedValue,
                }),
            }
        );
        const content = await response.json();
        window.location.reload()
    };

    const candidate = async () => {
        const response = await fetch(
            `https://localhost:8000/api/processing/candidate`,
            {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                credentials: "include", //cookie getter
                body: JSON.stringify({
                    project_name: props.proj.project_name,
                    user_id: props.proj.user_id,
                    file_path: props.proj.file_path,
                    voxel_size: inputedValue,
                }),
            }
        );
        const content = await response.json();
        window.location.reload()
    };

    const onClickProcess = () => {
        if(inputedValue!= 0) {
            if (!error) {
                switch (props.mode) {
                    case 'random': {
                        random()
                        props.closeModal(false)
                        break;
                    }
                    case 'candidate': {
                        candidate()
                        props.closeModal(false)
                        break;
                    }
                    case 'barycenter': {
                        barycenter()
                        props.closeModal(false)
                        break;
                    }
                }
            }
        } else {
            setError(true);
            setErrorText('Значение не может быть пустым');
        }
    }

    const handleChangeValue = (event: any) => {
        const partialValue = +event.target.value
        if(isNaN(partialValue)){
            setError(true);
            setErrorText(event.target.value + ' содержит запрещенные символы');
        } else {
            setError(false);
            setErrorText('')
            setInputedValue(partialValue)
        }
    }

    return (
        <div className={styles.modalWindow}>
            <div className={styles.modalWindowContainer}>
                <div className={styles.modalWindowContainerContentLayout}>
                    <div>Необходимо задать параметры:</div>
                    <div className={styles.modalWindowContainerContentLayoutInputLayout}>
                        <div className={styles.modalWindowContainerContentLayoutInputLayoutText}>{getNameOfInput()}</div>
                        <input type={"text"} onChange={(event) => handleChangeValue(event)}/>
                    </div>
                </div>
                <div className={styles.modalWindowContainerErrorText}>{errorText}</div>
                <div className={styles.modalWindowContainerButtonsLayout}>
                    <button className={styles.modalWindowContainerButtonsLayoutButtonChancel}
                            onClick={() => props.closeModal(false)}
                    >Close</button>
                    <button className={styles.modalWindowContainerButtonsLayoutButtonProcess}
                            onClick={onClickProcess}
                    >
                        Process
                    </button>
                </div>
            </div>
        </div>
    )
}