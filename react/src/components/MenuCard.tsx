import React, {useState} from "react";

export default function MenuCard(props:{user: any, mode: any}){

    const [show, setShow] = useState(false)
    const [buttonText, setButtonText] = useState('Подробнее +')

    const onClickItem = () => {
        setShow((prev) => !prev)
        if(!show){
            setButtonText('Подробнее -')
        } else {
            setButtonText('Подробнее +')
        }
    }

    return (
        <>
            {props.mode &&
                <div>user</div>
            }
            {!props.mode &&
            <div className='profileCard'>
                <div className='cardLayout'>
                    <div className='cardHeader'>
                        <b className='cardTitle'>
                            Профиль
                        </b>
                    </div>
                    <div className='userInfo'>Логин: {props.user.username}</div>
                    <div className='moreInfoButton' onClick={onClickItem}>{buttonText}</div>
                    {show &&
                    <div className='moreInfoBlock'>
                        <div className='itemRow'>
                            <div className='textRow'>E-mail: {props.user.email}</div>
                        </div>
                        <div className='itemRow'>
                            <div className='textRow'>Проекты: {props.user.project_number}</div>
                        </div>
                        <div className='itemRow'>
                            <div className='textRow'>Места использовано: {(props.user.used_space/1073741824).toFixed(4)} GB</div>
                        </div>
                        <div className='itemRow'>
                            <div className='textRow'>Места доступно: {(props.user.available/1073741824).toFixed(4)} GB</div>
                        </div>
                        <div className='itemsButtonsGroup'>
                            <button className='buttonSettings'/>
                            <button className='buttonSettings'/>
                            <button className='buttonSettings'/>
                        </div>
                    </div>
                    }
                </div>
            </div>
            }
        </>
    )
}