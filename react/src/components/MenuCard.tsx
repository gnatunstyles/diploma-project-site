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
                    <div className='userInfo'>Логин: alexx9911</div>
                    <div className='moreInfoButton' onClick={onClickItem}>{buttonText}</div>
                    {show &&
                    <div className='moreInfoBlock'>
                        <div className='itemRow'>
                            <div className='textRow'>E-mail: alexx@alexx.ru</div>
                        </div>
                        <div className='itemRow'>
                            <div className='textRow'>Телефон: 8800-555-35-35</div>
                        </div>
                        <div className='itemRow'>
                            <div className='textRow'>Еще инфа: текст</div>
                        </div>
                        <div className='itemRow'>
                            <div className='textRow'>E-mail: alexx@alexx.ru</div>
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