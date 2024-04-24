import React from 'react';
import './OnOff.css';
import "./../fonts/Inter-Regular.ttf"

function Namespaces({isName,setIsName}) {

  const handleClick = () => {
    setIsName(prevState => !prevState);
  };

  return (
    <div style={{marginTop:'10px',marginLeft:'-20px'}}>
        <div style={{display:'inline-block',verticalAlign:'middle',marginRight:'10px',fontFamily:"Inter Regular",color:"white"}}>Without Namespaces</div>
        <div className={`switch ${isName ? 'on' : 'off'}`} onClick={handleClick} style={{verticalAlign:'middle'}}>
            <div className="slider"></div>
        </div>
        <div style={{display:'inline-block',verticalAlign:'middle',marginLeft:'10px',fontFamily:"Inter Regular",color:"white"}}>With Namespaces</div>
    </div>
  );
}

export default Namespaces;