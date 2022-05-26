import React, { useEffect, useState } from "react";

const Home = (props: { userData: any }) => {
  console.log(props.userData);
  const [show, setShow] = useState(false)
  const [hidePannel, setHidePannel] = useState(false)
  const [buttonText, setButtonText] = useState('Подробнее +')

  const onClickItem = () => {
      setShow((prev) => !prev)
      if(!show){
          setButtonText('Подробнее -')
      } else {
          setButtonText('Подробнее +')
      }
  }

  const onCLickMenu = () => {
      setHidePannel((prev) => !prev)
  }

  return (
    <div>
        <div className='layout'>
            <div className='leftPannel' style={hidePannel ? {width: 80, minWidth: 80} : {}}>
                <div className='menuButton' onClick={onCLickMenu}>Меню</div>
                {!hidePannel &&
                    <div className='profileCard'>
                        <div className='cardLayout'>
                            <div className='cardHeader'>
                                <b className='cardTitle'>
                                    Профиль
                                </b>
                            </div>
                            <div className='userInfo'>Логин: {props.userData.username}</div>
                            <div className='moreInfoButton' onClick={onClickItem}>{buttonText}</div>
                            {show &&
                            <div className='moreInfoBlock'>
                                <div className='itemRow'>
                                    <div className='textRow'><b>E-mail: </b> {props.userData.email}</div>
                                </div>
                                <div className='itemRow'>
                                    <div className='textRow'><b>Проекты:</b> 1</div>
                                </div>
                                <div className='itemRow'>
                                    <div className='textRow'> <b>Места использовано: </b>{(props.userData.used_space/1073741824).toFixed(4)} GB</div>
                                </div>
                                <div className='itemRow'>
                                    <div className='textRow'><b>Места осталось: </b>{(props.userData.avaliable/1073741824).toFixed(4)} GB</div>
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
            </div>
            <div className='mainPannel'>
                <div className='mainContent'>контент</div>
            </div>
        </div>
    </div>
  );

};

export default Home;
