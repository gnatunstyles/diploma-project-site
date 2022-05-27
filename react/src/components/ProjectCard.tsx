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
    
    let link = props.proj.link
    return (
        <div className={styles.item}>
            <div className={styles.itemName}>{props.proj.project_name}</div>
            <div className={styles.itemDivider}/>
            <div className={styles.itemBody}>
                <b>Информация о проекте </b>
                <div className={styles.itemBodyListParams}>
                    <div className={styles.itemBodyListParamsListItem}>Описание: {props.proj.info}</div>
                    <div className={styles.itemBodyListParamsListItem}>Размер: {((props.proj.size)/1000000000).toFixed(3)} GB</div>
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
            onClick={() =>
              navigator.clipboard.writeText(
                `${link}`
              )
            }
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
          >
            Processing
          </button>
          <button
            type="button"
            className="btn btn-outline-dark w-100 mx-0"
            data-bs-dismiss="modal"
          >
            Delete
          </button>
        </div>
            </div>
        </div>
    )
}