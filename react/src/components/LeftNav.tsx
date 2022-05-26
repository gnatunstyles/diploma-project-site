import styles from '../styles/leftMenu.module.sass'
import {Link} from "react-router-dom";
import React from "react";
export default  function LeftNav(props: {mode: boolean}){
    return (
        <div className={styles.Menu}>
            <Link to={"/projects"}  className={styles.MenuItem}>
                {!props.mode &&
                    <div className={styles.MenuItemButtonText}>Проекты</div>
                }
                {props.mode &&
                <div className={styles.MenuItemButtonText}>|||</div>
                }
            </Link>
            <Link to={"/upload"} className={styles.MenuItem}>
                {!props.mode &&
                    <div className={styles.MenuItemButtonText}>Загрузить</div>
                }
                {props.mode &&
                <div className={styles.MenuItemButtonText}>==</div>
                }
            </Link>
        </div>
    )
}