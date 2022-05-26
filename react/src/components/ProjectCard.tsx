import styles from '../styles/projectCard.module.sass'

export default function ProjectCard(props: {proj: any}){

    //     ID: 10,
    //     CreatedAt: "2022-05-26T21:34:18.144414+03:00",
    //     UpdatedAt: "2022-05-26T22:38:07.510554+03:00",
    //     DeletedAt: null,
    //     user_id: 1,
    //     project_name: "asdqwe",
    //     info: "zxczxczxc",
    //     size: 185075623,
    //     link: "http://localhost:1234/projects/1/kekich.html"
    return (
        <div className={styles.item}>
            <div className={styles.itemName}>Проект {props.proj.project_name}</div>
            <div className={styles.itemDivider}/>
            <div className={styles.itemBody}>
                <b>Информация о проекте </b>
                <div className={styles.itemBodyListParams}>
                    <div className={styles.itemBodyListParamsListItem}>Описание: {props.proj.info}</div>
                    <div className={styles.itemBodyListParamsListItem}>Размер: {props.proj.size}</div>
                </div>
                ghjfhyjf
            </div>
        </div>
    )
}