import React from 'react';
import './OnOff.css';
import "./../fonts/Inter-Regular.ttf"

function OnOff({isOn,setIsOn}) {

  const handleClick = () => {
    setIsOn(prevState => !prevState);
  };

  return (
    <div>
        <div style={{display:'inline-block',verticalAlign:'middle',marginRight:'10px',fontFamily:"Inter Regular",color:"white"}}>BFS</div>
        <div className={`switch ${isOn ? 'on' : 'off'}`} onClick={handleClick} style={{verticalAlign:'middle'}}>
            <div className="slider"></div>
        </div>
        <div style={{display:'inline-block',verticalAlign:'middle',marginLeft:'10px',fontFamily:"Inter Regular",color:"white"}}>IDS</div>
    </div>
  );
}

export default OnOff;